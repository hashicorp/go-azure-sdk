package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentWhatIfProperties struct {
	DebugSetting                *DebugSetting                `json:"debugSetting"`
	ExpressionEvaluationOptions *ExpressionEvaluationOptions `json:"expressionEvaluationOptions"`
	Mode                        DeploymentMode               `json:"mode"`
	OnErrorDeployment           *OnErrorDeployment           `json:"onErrorDeployment"`
	Parameters                  *interface{}                 `json:"parameters,omitempty"`
	ParametersLink              *ParametersLink              `json:"parametersLink"`
	Template                    *interface{}                 `json:"template,omitempty"`
	TemplateLink                *TemplateLink                `json:"templateLink"`
	WhatIfSettings              *DeploymentWhatIfSettings    `json:"whatIfSettings"`
}
