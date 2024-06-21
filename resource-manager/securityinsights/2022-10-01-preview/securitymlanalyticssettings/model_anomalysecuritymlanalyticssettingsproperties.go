package securitymlanalyticssettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnomalySecurityMLAnalyticsSettingsProperties struct {
	AnomalySettingsVersion   *int64                                   `json:"anomalySettingsVersion,omitempty"`
	AnomalyVersion           string                                   `json:"anomalyVersion"`
	CustomizableObservations *interface{}                             `json:"customizableObservations,omitempty"`
	Description              *string                                  `json:"description,omitempty"`
	DisplayName              string                                   `json:"displayName"`
	Enabled                  bool                                     `json:"enabled"`
	Frequency                string                                   `json:"frequency"`
	IsDefaultSettings        bool                                     `json:"isDefaultSettings"`
	LastModifiedUtc          *string                                  `json:"lastModifiedUtc,omitempty"`
	RequiredDataConnectors   *[]SecurityMLAnalyticsSettingsDataSource `json:"requiredDataConnectors,omitempty"`
	SettingsDefinitionId     *string                                  `json:"settingsDefinitionId,omitempty"`
	SettingsStatus           SettingsStatus                           `json:"settingsStatus"`
	Tactics                  *[]AttackTactic                          `json:"tactics,omitempty"`
	Techniques               *[]string                                `json:"techniques,omitempty"`
}
