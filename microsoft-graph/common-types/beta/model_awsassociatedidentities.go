package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAssociatedIdentities struct {
	All       *[]AwsIdentity `json:"all,omitempty"`
	ODataType *string        `json:"@odata.type,omitempty"`
	Roles     *[]AwsRole     `json:"roles,omitempty"`
	Users     *[]AwsUser     `json:"users,omitempty"`
}
