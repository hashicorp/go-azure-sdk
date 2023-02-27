// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositories

type RepoOperationPredicate struct {
	FullName *string
	Url      *string
}

func (p RepoOperationPredicate) Matches(input Repo) bool {

	if p.FullName != nil && (input.FullName == nil && *p.FullName != *input.FullName) {
		return false
	}

	if p.Url != nil && (input.Url == nil && *p.Url != *input.Url) {
		return false
	}

	return true
}
