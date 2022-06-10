package autoscalevcores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScaleVCoreSku struct {
	Capacity *int64        `json:"capacity,omitempty"`
	Name     string        `json:"name"`
	Tier     *VCoreSkuTier `json:"tier,omitempty"`
}
