package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmApplicationHealthPolicy struct {
	ConsiderWarningAsError                  *bool                                  `json:"considerWarningAsError,omitempty"`
	DefaultServiceTypeHealthPolicy          *ArmServiceTypeHealthPolicy            `json:"defaultServiceTypeHealthPolicy,omitempty"`
	MaxPercentUnhealthyDeployedApplications *int64                                 `json:"maxPercentUnhealthyDeployedApplications,omitempty"`
	ServiceTypeHealthPolicyMap              *map[string]ArmServiceTypeHealthPolicy `json:"serviceTypeHealthPolicyMap,omitempty"`
}
