package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageTierCapability struct {
	Iops       *int64  `json:"iops,omitempty"`
	IsBaseline *bool   `json:"isBaseline,omitempty"`
	Name       *string `json:"name,omitempty"`
	Status     *string `json:"status,omitempty"`
	TierName   *string `json:"tierName,omitempty"`
}
