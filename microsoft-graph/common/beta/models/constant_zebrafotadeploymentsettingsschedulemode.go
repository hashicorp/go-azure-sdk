package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsScheduleMode string

const (
	ZebraFotaDeploymentSettingsScheduleModeinstallNow ZebraFotaDeploymentSettingsScheduleMode = "InstallNow"
	ZebraFotaDeploymentSettingsScheduleModescheduled  ZebraFotaDeploymentSettingsScheduleMode = "Scheduled"
)

func PossibleValuesForZebraFotaDeploymentSettingsScheduleMode() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsScheduleModeinstallNow),
		string(ZebraFotaDeploymentSettingsScheduleModescheduled),
	}
}

func parseZebraFotaDeploymentSettingsScheduleMode(input string) (*ZebraFotaDeploymentSettingsScheduleMode, error) {
	vals := map[string]ZebraFotaDeploymentSettingsScheduleMode{
		"installnow": ZebraFotaDeploymentSettingsScheduleModeinstallNow,
		"scheduled":  ZebraFotaDeploymentSettingsScheduleModescheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsScheduleMode(input)
	return &out, nil
}
