package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosSingleSignOnSettings struct {
	AllowedAppsList       *[]AppListItem `json:"allowedAppsList,omitempty"`
	AllowedUrls           *[]string      `json:"allowedUrls,omitempty"`
	DisplayName           *string        `json:"displayName,omitempty"`
	KerberosPrincipalName *string        `json:"kerberosPrincipalName,omitempty"`
	KerberosRealm         *string        `json:"kerberosRealm,omitempty"`
	ODataType             *string        `json:"@odata.type,omitempty"`
}
