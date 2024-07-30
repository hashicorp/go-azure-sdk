package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesContentApproval struct {
	Content            *DeployableContent                `json:"content,omitempty"`
	CreatedDateTime    *string                           `json:"createdDateTime,omitempty"`
	DeploymentSettings *WindowsUpdatesDeploymentSettings `json:"deploymentSettings,omitempty"`
	Deployments        *[]WindowsUpdatesDeployment       `json:"deployments,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	IsRevoked          *bool                             `json:"isRevoked,omitempty"`
	ODataType          *string                           `json:"@odata.type,omitempty"`
	RevokedDateTime    *string                           `json:"revokedDateTime,omitempty"`
	UpdatePolicy       *WindowsUpdatesUpdatePolicy       `json:"updatePolicy,omitempty"`
}
