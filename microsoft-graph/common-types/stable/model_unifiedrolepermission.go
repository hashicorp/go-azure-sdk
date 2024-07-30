package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRolePermission struct {
	AllowedResourceActions  *[]string `json:"allowedResourceActions,omitempty"`
	Condition               *string   `json:"condition,omitempty"`
	ExcludedResourceActions *[]string `json:"excludedResourceActions,omitempty"`
	ODataType               *string   `json:"@odata.type,omitempty"`
}
