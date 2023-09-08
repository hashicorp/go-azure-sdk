package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterInfo struct {
	Identity         *IdentitySet                   `json:"identity,omitempty"`
	ODataType        *string                        `json:"@odata.type,omitempty"`
	PresenterDetails *VirtualEventPresenterDetails  `json:"presenterDetails,omitempty"`
	Role             *VirtualEventPresenterInfoRole `json:"role,omitempty"`
	Upn              *string                        `json:"upn,omitempty"`
}
