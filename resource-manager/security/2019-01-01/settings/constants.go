package settings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingKind string

const (
	SettingKindAlertSuppressionSetting SettingKind = "AlertSuppressionSetting"
	SettingKindDataExportSettings      SettingKind = "DataExportSettings"
)

func PossibleValuesForSettingKind() []string {
	return []string{
		string(SettingKindAlertSuppressionSetting),
		string(SettingKindDataExportSettings),
	}
}

func (s *SettingKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSettingKind(input string) (*SettingKind, error) {
	vals := map[string]SettingKind{
		"alertsuppressionsetting": SettingKindAlertSuppressionSetting,
		"dataexportsettings":      SettingKindDataExportSettings,
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

func (s *SettingName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
