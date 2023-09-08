package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdatePolicy struct {
	Audience              *WindowsUpdatesDeploymentAudience     `json:"audience,omitempty"`
	ComplianceChangeRules *[]WindowsUpdatesComplianceChangeRule `json:"complianceChangeRules,omitempty"`
	ComplianceChanges     *[]WindowsUpdatesComplianceChange     `json:"complianceChanges,omitempty"`
	CreatedDateTime       *string                               `json:"createdDateTime,omitempty"`
	DeploymentSettings    *WindowsUpdatesDeploymentSettings     `json:"deploymentSettings,omitempty"`
	Id                    *string                               `json:"id,omitempty"`
	ODataType             *string                               `json:"@odata.type,omitempty"`
}
