package bot

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BotProperties struct {
	AllSettings                       *map[string]string           `json:"allSettings,omitempty"`
	AppPasswordHint                   *string                      `json:"appPasswordHint,omitempty"`
	CmekEncryptionStatus              *string                      `json:"cmekEncryptionStatus,omitempty"`
	CmekKeyVaultUrl                   *string                      `json:"cmekKeyVaultUrl,omitempty"`
	ConfiguredChannels                *[]string                    `json:"configuredChannels,omitempty"`
	Description                       *string                      `json:"description,omitempty"`
	DeveloperAppInsightKey            *string                      `json:"developerAppInsightKey,omitempty"`
	DeveloperAppInsightsApiKey        *string                      `json:"developerAppInsightsApiKey,omitempty"`
	DeveloperAppInsightsApplicationId *string                      `json:"developerAppInsightsApplicationId,omitempty"`
	DisableLocalAuth                  *bool                        `json:"disableLocalAuth,omitempty"`
	DisplayName                       string                       `json:"displayName"`
	EnabledChannels                   *[]string                    `json:"enabledChannels,omitempty"`
	Endpoint                          string                       `json:"endpoint"`
	EndpointVersion                   *string                      `json:"endpointVersion,omitempty"`
	IconUrl                           *string                      `json:"iconUrl,omitempty"`
	IsCmekEnabled                     *bool                        `json:"isCmekEnabled,omitempty"`
	IsDeveloperAppInsightsApiKeySet   *bool                        `json:"isDeveloperAppInsightsApiKeySet,omitempty"`
	IsStreamingSupported              *bool                        `json:"isStreamingSupported,omitempty"`
	LuisAppIds                        *[]string                    `json:"luisAppIds,omitempty"`
	LuisKey                           *string                      `json:"luisKey,omitempty"`
	ManifestUrl                       *string                      `json:"manifestUrl,omitempty"`
	MigrationToken                    *string                      `json:"migrationToken,omitempty"`
	MsaAppId                          string                       `json:"msaAppId"`
	MsaAppMSIResourceId               *string                      `json:"msaAppMSIResourceId,omitempty"`
	MsaAppTenantId                    *string                      `json:"msaAppTenantId,omitempty"`
	MsaAppType                        *MsaAppType                  `json:"msaAppType,omitempty"`
	OpenWithHint                      *string                      `json:"openWithHint,omitempty"`
	Parameters                        *map[string]string           `json:"parameters,omitempty"`
	PrivateEndpointConnections        *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                 *string                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess               *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
	PublishingCredentials             *string                      `json:"publishingCredentials,omitempty"`
	SchemaTransformationVersion       *string                      `json:"schemaTransformationVersion,omitempty"`
	StorageResourceId                 *string                      `json:"storageResourceId,omitempty"`
	TenantId                          *string                      `json:"tenantId,omitempty"`
}
