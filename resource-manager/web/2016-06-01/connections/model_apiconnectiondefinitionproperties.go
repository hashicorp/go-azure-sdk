package connections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiConnectionDefinitionProperties struct {
	Api                      *ApiReference                 `json:"api,omitempty"`
	ChangedTime              *string                       `json:"changedTime,omitempty"`
	CreatedTime              *string                       `json:"createdTime,omitempty"`
	CustomParameterValues    *map[string]string            `json:"customParameterValues,omitempty"`
	DisplayName              *string                       `json:"displayName,omitempty"`
	NonSecretParameterValues *map[string]string            `json:"nonSecretParameterValues,omitempty"`
	ParameterValues          *map[string]string            `json:"parameterValues,omitempty"`
	Statuses                 *[]ConnectionStatusDefinition `json:"statuses,omitempty"`
	TestLinks                *[]ApiConnectionTestLink      `json:"testLinks,omitempty"`
}
