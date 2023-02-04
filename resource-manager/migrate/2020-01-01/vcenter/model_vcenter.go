package vcenter

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenter struct {
	Id         *string            `json:"id,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Properties *VCenterProperties `json:"properties,omitempty"`
	Type       *string            `json:"type,omitempty"`
}
