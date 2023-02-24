package settings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityProviders string

const (
	EntityProvidersActiveDirectory      EntityProviders = "ActiveDirectory"
	EntityProvidersAzureActiveDirectory EntityProviders = "AzureActiveDirectory"
)

func PossibleValuesForEntityProviders() []string {
	return []string{
		string(EntityProvidersActiveDirectory),
		string(EntityProvidersAzureActiveDirectory),
	}
}

func parseEntityProviders(input string) (*EntityProviders, error) {
	vals := map[string]EntityProviders{
		"activedirectory":      EntityProvidersActiveDirectory,
		"azureactivedirectory": EntityProvidersAzureActiveDirectory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityProviders(input)
	return &out, nil
}

type SettingKind string

const (
	SettingKindAnomalies       SettingKind = "Anomalies"
	SettingKindEntityAnalytics SettingKind = "EntityAnalytics"
	SettingKindEyesOn          SettingKind = "EyesOn"
	SettingKindUeba            SettingKind = "Ueba"
)

func PossibleValuesForSettingKind() []string {
	return []string{
		string(SettingKindAnomalies),
		string(SettingKindEntityAnalytics),
		string(SettingKindEyesOn),
		string(SettingKindUeba),
	}
}

func parseSettingKind(input string) (*SettingKind, error) {
	vals := map[string]SettingKind{
		"anomalies":       SettingKindAnomalies,
		"entityanalytics": SettingKindEntityAnalytics,
		"eyeson":          SettingKindEyesOn,
		"ueba":            SettingKindUeba,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingKind(input)
	return &out, nil
}

type UebaDataSources string

const (
	UebaDataSourcesAuditLogs     UebaDataSources = "AuditLogs"
	UebaDataSourcesAzureActivity UebaDataSources = "AzureActivity"
	UebaDataSourcesSecurityEvent UebaDataSources = "SecurityEvent"
	UebaDataSourcesSigninLogs    UebaDataSources = "SigninLogs"
)

func PossibleValuesForUebaDataSources() []string {
	return []string{
		string(UebaDataSourcesAuditLogs),
		string(UebaDataSourcesAzureActivity),
		string(UebaDataSourcesSecurityEvent),
		string(UebaDataSourcesSigninLogs),
	}
}

func parseUebaDataSources(input string) (*UebaDataSources, error) {
	vals := map[string]UebaDataSources{
		"auditlogs":     UebaDataSourcesAuditLogs,
		"azureactivity": UebaDataSourcesAzureActivity,
		"securityevent": UebaDataSourcesSecurityEvent,
		"signinlogs":    UebaDataSourcesSigninLogs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UebaDataSources(input)
	return &out, nil
}
