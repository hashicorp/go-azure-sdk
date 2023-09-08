package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnection struct {
	ActivitySettings          *ExternalConnectorsActivitySettings                            `json:"activitySettings,omitempty"`
	ComplianceSettings        *ExternalConnectorsComplianceSettings                          `json:"complianceSettings,omitempty"`
	Configuration             *ExternalConnectorsConfiguration                               `json:"configuration,omitempty"`
	ConnectorId               *string                                                        `json:"connectorId,omitempty"`
	Description               *string                                                        `json:"description,omitempty"`
	EnabledContentExperiences *ExternalConnectorsExternalConnectionEnabledContentExperiences `json:"enabledContentExperiences,omitempty"`
	Groups                    *[]ExternalConnectorsExternalGroup                             `json:"groups,omitempty"`
	Id                        *string                                                        `json:"id,omitempty"`
	IngestedItemsCount        *int64                                                         `json:"ingestedItemsCount,omitempty"`
	Items                     *[]ExternalConnectorsExternalItem                              `json:"items,omitempty"`
	Name                      *string                                                        `json:"name,omitempty"`
	ODataType                 *string                                                        `json:"@odata.type,omitempty"`
	Operations                *[]ExternalConnectorsConnectionOperation                       `json:"operations,omitempty"`
	Quota                     *ExternalConnectorsConnectionQuota                             `json:"quota,omitempty"`
	Schema                    *ExternalConnectorsSchema                                      `json:"schema,omitempty"`
	SearchSettings            *ExternalConnectorsSearchSettings                              `json:"searchSettings,omitempty"`
	State                     *ExternalConnectorsExternalConnectionState                     `json:"state,omitempty"`
}
