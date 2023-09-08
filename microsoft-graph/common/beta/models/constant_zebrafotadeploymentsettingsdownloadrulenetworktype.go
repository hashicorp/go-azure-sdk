package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsDownloadRuleNetworkType string

const (
	ZebraFotaDeploymentSettingsDownloadRuleNetworkTypeany             ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "Any"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkTypecellular        ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "Cellular"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifi            ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "Wifi"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifiAndCellular ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "WifiAndCellular"
)

func PossibleValuesForZebraFotaDeploymentSettingsDownloadRuleNetworkType() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkTypeany),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkTypecellular),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifi),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifiAndCellular),
	}
}

func parseZebraFotaDeploymentSettingsDownloadRuleNetworkType(input string) (*ZebraFotaDeploymentSettingsDownloadRuleNetworkType, error) {
	vals := map[string]ZebraFotaDeploymentSettingsDownloadRuleNetworkType{
		"any":             ZebraFotaDeploymentSettingsDownloadRuleNetworkTypeany,
		"cellular":        ZebraFotaDeploymentSettingsDownloadRuleNetworkTypecellular,
		"wifi":            ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifi,
		"wifiandcellular": ZebraFotaDeploymentSettingsDownloadRuleNetworkTypewifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsDownloadRuleNetworkType(input)
	return &out, nil
}
