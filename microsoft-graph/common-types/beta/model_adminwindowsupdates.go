package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminWindowsUpdates struct {
	Catalog             *WindowsUpdatesCatalog              `json:"catalog,omitempty"`
	DeploymentAudiences *[]WindowsUpdatesDeploymentAudience `json:"deploymentAudiences,omitempty"`
	Deployments         *[]WindowsUpdatesDeployment         `json:"deployments,omitempty"`
	Id                  *string                             `json:"id,omitempty"`
	ODataType           *string                             `json:"@odata.type,omitempty"`
	ResourceConnections *[]WindowsUpdatesResourceConnection `json:"resourceConnections,omitempty"`
	UpdatableAssets     *[]WindowsUpdatesUpdatableAsset     `json:"updatableAssets,omitempty"`
	UpdatePolicies      *[]WindowsUpdatesUpdatePolicy       `json:"updatePolicies,omitempty"`
}
