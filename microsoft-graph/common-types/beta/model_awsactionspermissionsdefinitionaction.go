package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsActionsPermissionsDefinitionAction struct {
	AssignToRoleId *string         `json:"assignToRoleId,omitempty"`
	ODataType      *string         `json:"@odata.type,omitempty"`
	Statements     *[]AwsStatement `json:"statements,omitempty"`
}
