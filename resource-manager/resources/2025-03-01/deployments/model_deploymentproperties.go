package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentProperties struct {
	DebugSetting                *DebugSetting                                        `json:"debugSetting,omitempty"`
	ExpressionEvaluationOptions *ExpressionEvaluationOptions                         `json:"expressionEvaluationOptions,omitempty"`
	ExtensionConfigs            *map[string]map[string]DeploymentExtensionConfigItem `json:"extensionConfigs,omitempty"`
	Mode                        DeploymentMode                                       `json:"mode"`
	OnErrorDeployment           *OnErrorDeployment                                   `json:"onErrorDeployment,omitempty"`
	Parameters                  *map[string]DeploymentParameter                      `json:"parameters,omitempty"`
	ParametersLink              *ParametersLink                                      `json:"parametersLink,omitempty"`
	Template                    *interface{}                                         `json:"template,omitempty"`
	TemplateLink                *TemplateLink                                        `json:"templateLink,omitempty"`
	ValidationLevel             *ValidationLevel                                     `json:"validationLevel,omitempty"`
}
