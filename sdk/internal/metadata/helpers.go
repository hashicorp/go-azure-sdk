// Copyright IBM Corp. 2022, 2026
// SPDX-License-Identifier: MPL-2.0

package metadata

import "strings"

func normalizeResourceId(resourceId string) string {
	return strings.TrimRight(resourceId, "/")
}

func normalizeEndpoint(endpoint string) string {
	return strings.TrimRight(endpoint, "/")
}
