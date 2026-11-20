#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PASS_COUNT=0

pass() {
  echo "PASS: $1"
  PASS_COUNT=$((PASS_COUNT + 1))
}

fail() {
  echo "FAIL: $1" >&2
  exit 1
}

echo "==> 1. Static workflow security & policy checks (go-azure-sdk)"

# Every uses: must be pinned to a 40-char commit SHA
unpinned=$(grep -rnE '^\s*-?\s*uses:\s*' "${REPO_ROOT}/.github/workflows/" | grep -Ev '@[0-9a-f]{40}(\s|#|$)' || true)
if [[ -n "$unpinned" ]]; then
  fail "Unpinned GitHub Actions found:\n$unpinned"
fi
pass "All GitHub Actions uses: references are pinned to 40-char SHAs"

# Disallowed patterns
disallowed_patterns=$(grep -rnE 'StrictHostKeyChecking=accept-new|persist-credentials:\s*true|::set-output|SERVICE_ACCOUNT_GO_SDK_TOKEN|AZURERM_DEPLOYMENT_WRITE_KEY|AZURERM_COMMENT_KEY|RELEASE_DEPLOYMENT_WRITE_KEY' "${REPO_ROOT}/.github/workflows/" "${REPO_ROOT}/scripts/" | grep -v 'scripts/test-workflows.sh' || true)
if [[ -n "$disallowed_patterns" ]]; then
  fail "Disallowed legacy/insecure pattern found:\n$disallowed_patterns"
fi
pass "No disallowed patterns or legacy secret references"

# Every actions/checkout must specify persist-credentials: false
checkout_missing_persist=$(awk '
  /uses: actions\/checkout@/ { in_co=1; found=0; file=FILENAME; line=FNR; next }
  in_co && /persist-credentials: false/ { found=1 }
  in_co && /^[[:space:]]*-[[:space:]]/ { if (!found) print file ":" line; in_co=0 }
  in_co && /^[[:space:]]*[a-zA-Z0-9_-]+:[[:space:]]*$/ && !/with:/ { if (!found) print file ":" line; in_co=0 }
  END { if (in_co && !found) print file ":" line }
' "${REPO_ROOT}"/.github/workflows/*.y*ml)
if [[ -n "$checkout_missing_persist" ]]; then
  fail "actions/checkout missing persist-credentials: false:\n$checkout_missing_persist"
fi
pass "Every actions/checkout sets persist-credentials: false"

# No single job may hold both contents: write and pull-requests: write
dual_write=$(awk '
  /^  [a-zA-Z0-9_-]+:$/ { job=$1; cw=0; prw=0; next }
  /contents:[[:space:]]*write/ { cw=1 }
  /pull-requests:[[:space:]]*write/ { prw=1 }
  cw && prw { print FILENAME " job " job; cw=0; prw=0 }
' "${REPO_ROOT}"/.github/workflows/*.y*ml)
if [[ -n "$dual_write" ]]; then
  fail "Job holds both contents: write and pull-requests: write:\n$dual_write"
fi
pass "No job holds both contents: write and pull-requests: write"

echo "==> 2. Hermetic tag-only release & version contract tests"

tag=$("${REPO_ROOT}/scripts/determine-git-tag.sh")
if [[ ! "$tag" =~ ^v0\.[0-9]{8}\.1[0-9]{6}$ ]]; then
  fail "determine-git-tag.sh produced unexpected version format: $tag"
fi
pass "determine-git-tag.sh output ($tag) matches publish-git-tag.sh regex"

tmp_dir=$(mktemp -d /tmp/gosdk-workflow-test.XXXXXX)
trap 'rm -rf "$tmp_dir"' EXIT
export GIT_CONFIG_NOSYSTEM=1
export HOME="$tmp_dir"

git init --bare -b main "$tmp_dir/remote.git" >/dev/null
git clone "$tmp_dir/remote.git" "$tmp_dir/work" >/dev/null 2>&1
cd "$tmp_dir/work"
git config user.name "test"
git config user.email "test@example.com"
echo "init" > README.md
git add README.md && git commit -m "initial" >/dev/null
git push origin HEAD:refs/heads/main >/dev/null 2>&1
head_before=$(git rev-parse HEAD)

# Invalid version format must fail
if "${REPO_ROOT}/scripts/publish-git-tag.sh" "v1.2.3" 2>/dev/null; then
  fail "publish-git-tag.sh should have rejected invalid version format"
fi
pass "publish-git-tag.sh rejected invalid version format"

# Mismatched RELEASE_SHA must fail
if RELEASE_SHA="0000000000000000000000000000000000000000" "${REPO_ROOT}/scripts/publish-git-tag.sh" "$tag" 2>/dev/null; then
  fail "publish-git-tag.sh should have rejected mismatched RELEASE_SHA"
fi
pass "publish-git-tag.sh rejected mismatched RELEASE_SHA"

# Valid tag-only release must push all 4 module tags without advancing main
"${REPO_ROOT}/scripts/publish-git-tag.sh" "$tag" >/dev/null
head_after=$(git --git-dir="$tmp_dir/remote.git" rev-parse refs/heads/main)
if [[ "$head_before" != "$head_after" ]]; then
  fail "publish-git-tag.sh modified main ($head_before -> $head_after)"
fi
for mod in sdk microsoft-graph resource-manager data-plane; do
  tag_sha=$(git --git-dir="$tmp_dir/remote.git" rev-parse "refs/tags/${mod}/${tag}" 2>/dev/null || true)
  if [[ "$tag_sha" != "$head_before" ]]; then
    fail "Missing or mismatched tag ${mod}/${tag} (got '$tag_sha', expected '$head_before')"
  fi
done
pass "publish-git-tag.sh pushed all 4 module tags to exact HEAD without modifying main"

echo "==> 3. Release jq label-filter tests"

JQ_EXPR='[.[] | select(.merged_at != null and .base.ref == "main") | .labels[].name] | index("release-once-merged") != null'

check_jq() {
  local desc=$1
  local payload=$2
  local expected=$3
  local actual
  actual=$(printf '%s' "$payload" | jq -r "$JQ_EXPR")
  if [[ "$actual" != "$expected" ]]; then
    fail "jq check '$desc': expected '$expected', got '$actual'"
  fi
}

check_jq "merged PR on main with label" \
  '[{"merged_at":"2026-11-20T09:00:00Z","base":{"ref":"main"},"labels":[{"name":"release-once-merged"}]}]' \
  "true"
check_jq "unmerged PR with label" \
  '[{"merged_at":null,"base":{"ref":"main"},"labels":[{"name":"release-once-merged"}]}]' \
  "false"
check_jq "merged PR on main without label" \
  '[{"merged_at":"2026-11-20T09:00:00Z","base":{"ref":"main"},"labels":[{"name":"documentation"}]}]' \
  "false"
check_jq "merged PR on main with empty labels" \
  '[{"merged_at":"2026-11-20T09:00:00Z","base":{"ref":"main"},"labels":[]}]' \
  "false"
check_jq "merged PR targeting non-main" \
  '[{"merged_at":"2026-11-20T09:00:00Z","base":{"ref":"dev"},"labels":[{"name":"release-once-merged"}]}]' \
  "false"
check_jq "direct commit with no PRs" \
  '[]' \
  "false"
pass "automation-release.yaml jq label-filter passed all 6 test cases"

echo "==> All ${PASS_COUNT} go-azure-sdk workflow checks passed!"
