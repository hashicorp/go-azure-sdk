package environments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentDeploymentProperties struct {
	ArmTemplateId *string                           `json:"armTemplateId,omitempty"`
	Parameters    *[]ArmTemplateParameterProperties `json:"parameters,omitempty"`
}
