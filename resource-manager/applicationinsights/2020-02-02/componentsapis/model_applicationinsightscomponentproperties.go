package componentsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationInsightsComponentProperties struct {
	AppId                           *string                      `json:"AppId,omitempty"`
	ApplicationId                   *string                      `json:"ApplicationId,omitempty"`
	ApplicationType                 ApplicationType              `json:"Application_Type"`
	ConnectionString                *string                      `json:"ConnectionString,omitempty"`
	CreationDate                    *string                      `json:"CreationDate,omitempty"`
	DisableIPMasking                *bool                        `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                        `json:"DisableLocalAuth,omitempty"`
	FlowType                        *FlowType                    `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                        `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                      `json:"HockeyAppId,omitempty"`
	HockeyAppToken                  *string                      `json:"HockeyAppToken,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                        `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *IngestionMode               `json:"IngestionMode,omitempty"`
	InstrumentationKey              *string                      `json:"InstrumentationKey,omitempty"`
	LaMigrationDate                 *string                      `json:"LaMigrationDate,omitempty"`
	Name                            *string                      `json:"Name,omitempty"`
	PrivateLinkScopedResources      *[]PrivateLinkScopedResource `json:"PrivateLinkScopedResources,omitempty"`
	ProvisioningState               *string                      `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *PublicNetworkAccessType     `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *PublicNetworkAccessType     `json:"publicNetworkAccessForQuery,omitempty"`
	RequestSource                   *RequestSource               `json:"Request_Source,omitempty"`
	RetentionInDays                 *int64                       `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                     `json:"SamplingPercentage,omitempty"`
	TenantId                        *string                      `json:"TenantId,omitempty"`
	WorkspaceResourceId             *string                      `json:"WorkspaceResourceId,omitempty"`
}
