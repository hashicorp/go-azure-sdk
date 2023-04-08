package regions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionalQuotaCapability struct {
	CoresAvailable *int64  `json:"cores_available,omitempty"`
	CoresUsed      *int64  `json:"cores_used,omitempty"`
	RegionName     *string `json:"region_name,omitempty"`
}
