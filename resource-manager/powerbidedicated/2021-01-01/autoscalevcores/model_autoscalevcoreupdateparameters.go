package autoscalevcores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScaleVCoreUpdateParameters struct {
	Properties *AutoScaleVCoreMutableProperties `json:"properties"`
	Sku        *AutoScaleVCoreSku               `json:"sku"`
	Tags       *map[string]string               `json:"tags,omitempty"`
}
