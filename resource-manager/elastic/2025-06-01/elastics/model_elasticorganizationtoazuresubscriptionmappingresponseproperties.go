package elastics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticOrganizationToAzureSubscriptionMappingResponseProperties struct {
	BilledAzureSubscriptionId *string              `json:"billedAzureSubscriptionId,omitempty"`
	ElasticOrganizationId     *string              `json:"elasticOrganizationId,omitempty"`
	ElasticOrganizationName   *string              `json:"elasticOrganizationName,omitempty"`
	MarketplaceSaasInfo       *MarketplaceSaaSInfo `json:"marketplaceSaasInfo,omitempty"`
}
