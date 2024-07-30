package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsPolicyPermissionsDefinitionAction struct {
	AssignToRoleId *string                           `json:"assignToRoleId,omitempty"`
	ODataType      *string                           `json:"@odata.type,omitempty"`
	Policies       *[]PermissionsDefinitionAwsPolicy `json:"policies,omitempty"`
}
