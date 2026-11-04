package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaxLimitRangeCapability struct {
	MaxValue  *int64            `json:"maxValue,omitempty"`
	MinValue  *int64            `json:"minValue,omitempty"`
	Reason    *string           `json:"reason,omitempty"`
	ScaleSize *int64            `json:"scaleSize,omitempty"`
	Status    *CapabilityStatus `json:"status,omitempty"`
}
