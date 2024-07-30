package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InactiveGroupFinding struct {
	ActionSummary         *ActionSummary               `json:"actionSummary,omitempty"`
	CreatedDateTime       *string                      `json:"createdDateTime,omitempty"`
	Group                 *AuthorizationSystemIdentity `json:"group,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex       `json:"permissionsCreepIndex,omitempty"`
}
