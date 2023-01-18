package integrationserviceenvironmentskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkuDefinitionSku struct {
	Name *IntegrationServiceEnvironmentSkuName `json:"name,omitempty"`
	Tier *string                               `json:"tier,omitempty"`
}
