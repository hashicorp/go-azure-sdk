package settings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingKind string

const (
	SettingKindAlertSuppressionSetting SettingKind = "AlertSuppressionSetting"
	SettingKindDataExportSetting       SettingKind = "DataExportSetting"
)

func PossibleValuesForSettingKind() []string {
	return []string{
		string(SettingKindAlertSuppressionSetting),
		string(SettingKindDataExportSetting),
	}
}

func parseSettingKind(input string) (*SettingKind, error) {
	vals := map[string]SettingKind{
		"alertsuppressionsetting": SettingKindAlertSuppressionSetting,
		"dataexportsetting":       SettingKindDataExportSetting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingKind(input)
	return &out, nil
}

type SettingName string

const (
	SettingNameMCAS  SettingName = "MCAS"
	SettingNameWDATP SettingName = "WDATP"
)

func PossibleValuesForSettingName() []string {
	return []string{
		string(SettingNameMCAS),
		string(SettingNameWDATP),
	}
}

func parseSettingName(input string) (*SettingName, error) {
	vals := map[string]SettingName{
		"mcas":  SettingNameMCAS,
		"wdatp": SettingNameWDATP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingName(input)
	return &out, nil
}
