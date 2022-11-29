package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationHealthPolicy struct {
	ConsiderWarningAsError                  bool                                `json:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          *ServiceTypeHealthPolicy            `json:"defaultServiceTypeHealthPolicy,omitempty"`
	MaxPercentUnhealthyDeployedApplications int64                               `json:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              *map[string]ServiceTypeHealthPolicy `json:"serviceTypeHealthPolicyMap,omitempty"`
}
