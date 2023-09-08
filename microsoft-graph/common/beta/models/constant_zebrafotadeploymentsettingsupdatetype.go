package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsUpdateType string

const (
	ZebraFotaDeploymentSettingsUpdateTypeauto   ZebraFotaDeploymentSettingsUpdateType = "Auto"
	ZebraFotaDeploymentSettingsUpdateTypecustom ZebraFotaDeploymentSettingsUpdateType = "Custom"
	ZebraFotaDeploymentSettingsUpdateTypelatest ZebraFotaDeploymentSettingsUpdateType = "Latest"
)

func PossibleValuesForZebraFotaDeploymentSettingsUpdateType() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsUpdateTypeauto),
		string(ZebraFotaDeploymentSettingsUpdateTypecustom),
		string(ZebraFotaDeploymentSettingsUpdateTypelatest),
	}
}

func parseZebraFotaDeploymentSettingsUpdateType(input string) (*ZebraFotaDeploymentSettingsUpdateType, error) {
	vals := map[string]ZebraFotaDeploymentSettingsUpdateType{
		"auto":   ZebraFotaDeploymentSettingsUpdateTypeauto,
		"custom": ZebraFotaDeploymentSettingsUpdateTypecustom,
		"latest": ZebraFotaDeploymentSettingsUpdateTypelatest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsUpdateType(input)
	return &out, nil
}
