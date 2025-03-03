package policydefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalEvaluationEnforcementSettings struct {
	EndpointSettings   *ExternalEvaluationEndpointSettings `json:"endpointSettings,omitempty"`
	MissingTokenAction *string                             `json:"missingTokenAction,omitempty"`
	ResultLifespan     *string                             `json:"resultLifespan,omitempty"`
	RoleDefinitionIds  *[]string                           `json:"roleDefinitionIds,omitempty"`
}
