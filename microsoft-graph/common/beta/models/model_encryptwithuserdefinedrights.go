package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithUserDefinedRights struct {
	AllowAdHocPermissions                *bool                                    `json:"allowAdHocPermissions,omitempty"`
	AllowMailForwarding                  *bool                                    `json:"allowMailForwarding,omitempty"`
	DecryptionRightsManagementTemplateId *string                                  `json:"decryptionRightsManagementTemplateId,omitempty"`
	EncryptWith                          *EncryptWithUserDefinedRightsEncryptWith `json:"encryptWith,omitempty"`
	Name                                 *string                                  `json:"name,omitempty"`
	ODataType                            *string                                  `json:"@odata.type,omitempty"`
}
