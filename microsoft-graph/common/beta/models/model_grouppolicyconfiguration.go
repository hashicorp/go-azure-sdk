package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfiguration struct {
	Assignments                      *[]GroupPolicyConfigurationAssignment                     `json:"assignments,omitempty"`
	CreatedDateTime                  *string                                                   `json:"createdDateTime,omitempty"`
	DefinitionValues                 *[]GroupPolicyDefinitionValue                             `json:"definitionValues,omitempty"`
	Description                      *string                                                   `json:"description,omitempty"`
	DisplayName                      *string                                                   `json:"displayName,omitempty"`
	Id                               *string                                                   `json:"id,omitempty"`
	LastModifiedDateTime             *string                                                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                        *string                                                   `json:"@odata.type,omitempty"`
	PolicyConfigurationIngestionType *GroupPolicyConfigurationPolicyConfigurationIngestionType `json:"policyConfigurationIngestionType,omitempty"`
	RoleScopeTagIds                  *[]string                                                 `json:"roleScopeTagIds,omitempty"`
}
