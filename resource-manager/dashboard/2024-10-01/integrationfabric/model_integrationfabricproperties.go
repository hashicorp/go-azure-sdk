package integrationfabric

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationFabricProperties struct {
	DataSourceResourceId *string            `json:"dataSourceResourceId,omitempty"`
	ProvisioningState    *ProvisioningState `json:"provisioningState,omitempty"`
	Scenarios            *[]string          `json:"scenarios,omitempty"`
	TargetResourceId     *string            `json:"targetResourceId,omitempty"`
}
