package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAssociatedIdentities struct {
	All             *[]GcpIdentity       `json:"all,omitempty"`
	ODataType       *string              `json:"@odata.type,omitempty"`
	ServiceAccounts *[]GcpServiceAccount `json:"serviceAccounts,omitempty"`
	Users           *[]GcpUser           `json:"users,omitempty"`
}
