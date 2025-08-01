package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetIPTag struct {
	IPTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}
