package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamMemberSettings struct {
	AllowAddRemoveApps                *bool   `json:"allowAddRemoveApps,omitempty"`
	AllowCreatePrivateChannels        *bool   `json:"allowCreatePrivateChannels,omitempty"`
	AllowCreateUpdateChannels         *bool   `json:"allowCreateUpdateChannels,omitempty"`
	AllowCreateUpdateRemoveConnectors *bool   `json:"allowCreateUpdateRemoveConnectors,omitempty"`
	AllowCreateUpdateRemoveTabs       *bool   `json:"allowCreateUpdateRemoveTabs,omitempty"`
	AllowDeleteChannels               *bool   `json:"allowDeleteChannels,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
}
