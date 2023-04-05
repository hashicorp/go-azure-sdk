package settings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
