package locations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Quota struct {
	HostsRemaining *map[string]int64 `json:"hostsRemaining,omitempty"`
	QuotaEnabled   *QuotaEnabled     `json:"quotaEnabled,omitempty"`
}
