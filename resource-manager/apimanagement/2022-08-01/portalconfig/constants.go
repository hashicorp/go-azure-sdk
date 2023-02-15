package portalconfig

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalSettingsCspMode string

const (
	PortalSettingsCspModeDisabled   PortalSettingsCspMode = "disabled"
	PortalSettingsCspModeEnabled    PortalSettingsCspMode = "enabled"
	PortalSettingsCspModeReportOnly PortalSettingsCspMode = "reportOnly"
)

func PossibleValuesForPortalSettingsCspMode() []string {
	return []string{
		string(PortalSettingsCspModeDisabled),
		string(PortalSettingsCspModeEnabled),
		string(PortalSettingsCspModeReportOnly),
	}
}

func parsePortalSettingsCspMode(input string) (*PortalSettingsCspMode, error) {
	vals := map[string]PortalSettingsCspMode{
		"disabled":   PortalSettingsCspModeDisabled,
		"enabled":    PortalSettingsCspModeEnabled,
		"reportonly": PortalSettingsCspModeReportOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PortalSettingsCspMode(input)
	return &out, nil
}
