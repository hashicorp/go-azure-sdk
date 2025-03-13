package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentProperties struct {
	AssignmentType             *AssignmentType                  `json:"assignmentType,omitempty"`
	DefinitionVersion          *string                          `json:"definitionVersion,omitempty"`
	Description                *string                          `json:"description,omitempty"`
	DisplayName                *string                          `json:"displayName,omitempty"`
	EffectiveDefinitionVersion *string                          `json:"effectiveDefinitionVersion,omitempty"`
	EnforcementMode            *EnforcementMode                 `json:"enforcementMode,omitempty"`
	InstanceId                 *string                          `json:"instanceId,omitempty"`
	LatestDefinitionVersion    *string                          `json:"latestDefinitionVersion,omitempty"`
	Metadata                   *interface{}                     `json:"metadata,omitempty"`
	NonComplianceMessages      *[]NonComplianceMessage          `json:"nonComplianceMessages,omitempty"`
	NotScopes                  *[]string                        `json:"notScopes,omitempty"`
	Overrides                  *[]Override                      `json:"overrides,omitempty"`
	Parameters                 *map[string]ParameterValuesValue `json:"parameters,omitempty"`
	PolicyDefinitionId         *string                          `json:"policyDefinitionId,omitempty"`
	ResourceSelectors          *[]ResourceSelector              `json:"resourceSelectors,omitempty"`
	Scope                      *string                          `json:"scope,omitempty"`
}
