package monitorsresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorProperties struct {
	ElasticProperties           *ElasticProperties       `json:"elasticProperties,omitempty"`
	GenerateApiKey              *bool                    `json:"generateApiKey,omitempty"`
	LiftrResourceCategory       *LiftrResourceCategories `json:"liftrResourceCategory,omitempty"`
	LiftrResourcePreference     *int64                   `json:"liftrResourcePreference,omitempty"`
	MonitoringStatus            *MonitoringStatus        `json:"monitoringStatus,omitempty"`
	PlanDetails                 *PlanDetails             `json:"planDetails,omitempty"`
	ProvisioningState           *ProvisioningState       `json:"provisioningState,omitempty"`
	SaaSAzureSubscriptionStatus *string                  `json:"saaSAzureSubscriptionStatus,omitempty"`
	SourceCampaignId            *string                  `json:"sourceCampaignId,omitempty"`
	SourceCampaignName          *string                  `json:"sourceCampaignName,omitempty"`
	SubscriptionState           *string                  `json:"subscriptionState,omitempty"`
	UserInfo                    *UserInfo                `json:"userInfo,omitempty"`
	Version                     *string                  `json:"version,omitempty"`
}
