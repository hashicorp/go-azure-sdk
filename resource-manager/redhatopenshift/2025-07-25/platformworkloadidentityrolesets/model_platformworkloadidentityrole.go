package platformworkloadidentityrolesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformWorkloadIdentityRole struct {
	OperatorName       *string `json:"operatorName,omitempty"`
	RoleDefinitionId   *string `json:"roleDefinitionId,omitempty"`
	RoleDefinitionName *string `json:"roleDefinitionName,omitempty"`
}
