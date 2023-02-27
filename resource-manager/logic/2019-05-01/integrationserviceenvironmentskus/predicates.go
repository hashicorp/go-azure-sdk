package integrationserviceenvironmentskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkuDefinitionOperationPredicate struct {
	ResourceType *string
}

func (p IntegrationServiceEnvironmentSkuDefinitionOperationPredicate) Matches(input IntegrationServiceEnvironmentSkuDefinition) bool {

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
