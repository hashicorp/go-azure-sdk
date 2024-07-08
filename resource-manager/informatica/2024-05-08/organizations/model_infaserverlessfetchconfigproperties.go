package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InfaServerlessFetchConfigProperties struct {
	AdvancedCustomProperties  *string `json:"advancedCustomProperties,omitempty"`
	ApplicationType           *string `json:"applicationType,omitempty"`
	ComputeUnits              *string `json:"computeUnits,omitempty"`
	ExecutionTimeout          *string `json:"executionTimeout,omitempty"`
	Platform                  *string `json:"platform,omitempty"`
	Region                    *string `json:"region,omitempty"`
	ResourceGroupName         *string `json:"resourceGroupName,omitempty"`
	ServerlessArmResourceId   *string `json:"serverlessArmResourceId,omitempty"`
	Subnet                    *string `json:"subnet,omitempty"`
	SubscriptionId            *string `json:"subscriptionId,omitempty"`
	SupplementaryFileLocation *string `json:"supplementaryFileLocation,omitempty"`
	Tags                      *string `json:"tags,omitempty"`
	TenantId                  *string `json:"tenantId,omitempty"`
	Vnet                      *string `json:"vnet,omitempty"`
}
