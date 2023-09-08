package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnSettings struct {
	KerberosServicePrincipalName       *string                                                   `json:"kerberosServicePrincipalName,omitempty"`
	KerberosSignOnMappingAttributeType *KerberosSignOnSettingsKerberosSignOnMappingAttributeType `json:"kerberosSignOnMappingAttributeType,omitempty"`
	ODataType                          *string                                                   `json:"@odata.type,omitempty"`
}
