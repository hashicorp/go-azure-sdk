package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityToolAwsRoleAdministratorFinding struct {
	CreatedDateTime       *string                                               `json:"createdDateTime,omitempty"`
	Id                    *string                                               `json:"id,omitempty"`
	Identity              *AuthorizationSystemIdentity                          `json:"identity,omitempty"`
	IdentityDetails       *IdentityDetails                                      `json:"identityDetails,omitempty"`
	ODataType             *string                                               `json:"@odata.type,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex                                `json:"permissionsCreepIndex,omitempty"`
	SecurityTools         *SecurityToolAwsRoleAdministratorFindingSecurityTools `json:"securityTools,omitempty"`
}
