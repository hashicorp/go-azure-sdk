package integrationserviceenvironmentskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkuDefinition struct {
	Capacity     *IntegrationServiceEnvironmentSkuCapacity      `json:"capacity,omitempty"`
	ResourceType *string                                        `json:"resourceType,omitempty"`
	Sku          *IntegrationServiceEnvironmentSkuDefinitionSku `json:"sku,omitempty"`
}
