package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentProperties struct {
	DisplayName        *string `json:"displayName,omitempty"`
	PolicyDefinitionId *string `json:"policyDefinitionId,omitempty"`
	Scope              *string `json:"scope,omitempty"`
}
