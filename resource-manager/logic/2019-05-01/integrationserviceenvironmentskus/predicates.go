package integrationserviceenvironmentskus

type IntegrationServiceEnvironmentSkuDefinitionOperationPredicate struct {
	ResourceType *string
}

func (p IntegrationServiceEnvironmentSkuDefinitionOperationPredicate) Matches(input IntegrationServiceEnvironmentSkuDefinition) bool {

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
