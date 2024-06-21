package dataconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CosmosDbDataConnectionProperties struct {
	CosmosDbAccountResourceId string             `json:"cosmosDbAccountResourceId"`
	CosmosDbContainer         string             `json:"cosmosDbContainer"`
	CosmosDbDatabase          string             `json:"cosmosDbDatabase"`
	ManagedIdentityObjectId   *string            `json:"managedIdentityObjectId,omitempty"`
	ManagedIdentityResourceId string             `json:"managedIdentityResourceId"`
	MappingRuleName           *string            `json:"mappingRuleName,omitempty"`
	ProvisioningState         *ProvisioningState `json:"provisioningState,omitempty"`
	RetrievalStartDate        *string            `json:"retrievalStartDate,omitempty"`
	TableName                 string             `json:"tableName"`
}
