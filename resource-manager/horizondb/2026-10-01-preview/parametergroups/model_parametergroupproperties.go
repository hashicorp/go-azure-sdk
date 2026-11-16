package parametergroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParameterGroupProperties struct {
	Description       *string                `json:"description,omitempty"`
	Parameters        *[]ParameterProperties `json:"parameters,omitempty"`
	PostgreSqlVersion PostgreSqlVersion      `json:"postgreSqlVersion"`
	ProvisioningState *ProvisioningState     `json:"provisioningState,omitempty"`
}
