package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOrganizationalScope struct {
	ODataType  *string                               `json:"@odata.type,omitempty"`
	ScopeNames *[]string                             `json:"scopeNames,omitempty"`
	ScopeType  *SecurityOrganizationalScopeScopeType `json:"scopeType,omitempty"`
}
