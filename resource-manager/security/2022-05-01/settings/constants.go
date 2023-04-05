package settings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingKind string

const (
	SettingKindAlertSuppressionSetting SettingKind = "AlertSuppressionSetting"
	SettingKindAlertSyncSettings       SettingKind = "AlertSyncSettings"
	SettingKindDataExportSettings      SettingKind = "DataExportSettings"
)

func PossibleValuesForSettingKind() []string {
	return []string{
		string(SettingKindAlertSuppressionSetting),
		string(SettingKindAlertSyncSettings),
		string(SettingKindDataExportSettings),
	}
}

type SettingName string

const (
	SettingNameMCAS                           SettingName = "MCAS"
	SettingNameSentinel                       SettingName = "Sentinel"
	SettingNameWDATP                          SettingName = "WDATP"
	SettingNameWDATPEXCLUDELINUXPUBLICPREVIEW SettingName = "WDATP_EXCLUDE_LINUX_PUBLIC_PREVIEW"
	SettingNameWDATPUNIFIEDSOLUTION           SettingName = "WDATP_UNIFIED_SOLUTION"
)

func PossibleValuesForSettingName() []string {
	return []string{
		string(SettingNameMCAS),
		string(SettingNameSentinel),
		string(SettingNameWDATP),
		string(SettingNameWDATPEXCLUDELINUXPUBLICPREVIEW),
		string(SettingNameWDATPUNIFIEDSOLUTION),
	}
}
