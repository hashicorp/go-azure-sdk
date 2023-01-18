package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentManagedApiProperties struct {
	ApiDefinitionUrl              *string                                                      `json:"apiDefinitionUrl,omitempty"`
	ApiDefinitions                *ApiResourceDefinitions                                      `json:"apiDefinitions,omitempty"`
	BackendService                *ApiResourceBackendService                                   `json:"backendService,omitempty"`
	Capabilities                  *[]string                                                    `json:"capabilities,omitempty"`
	Category                      *ApiTier                                                     `json:"category,omitempty"`
	ConnectionParameters          *map[string]interface{}                                      `json:"connectionParameters,omitempty"`
	DeploymentParameters          *IntegrationServiceEnvironmentManagedApiDeploymentParameters `json:"deploymentParameters,omitempty"`
	GeneralInformation            *ApiResourceGeneralInformation                               `json:"generalInformation,omitempty"`
	IntegrationServiceEnvironment *ResourceReference                                           `json:"integrationServiceEnvironment,omitempty"`
	Metadata                      *ApiResourceMetadata                                         `json:"metadata,omitempty"`
	Name                          *string                                                      `json:"name,omitempty"`
	Policies                      *ApiResourcePolicies                                         `json:"policies,omitempty"`
	ProvisioningState             *WorkflowProvisioningState                                   `json:"provisioningState,omitempty"`
	RuntimeUrls                   *[]string                                                    `json:"runtimeUrls,omitempty"`
}
