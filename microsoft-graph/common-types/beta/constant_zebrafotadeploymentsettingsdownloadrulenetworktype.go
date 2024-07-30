package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsDownloadRuleNetworkType string

const (
	ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Any             ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "any"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Cellular        ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "cellular"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Wifi            ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "wifi"
	ZebraFotaDeploymentSettingsDownloadRuleNetworkType_WifiAndCellular ZebraFotaDeploymentSettingsDownloadRuleNetworkType = "wifiAndCellular"
)

func PossibleValuesForZebraFotaDeploymentSettingsDownloadRuleNetworkType() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Any),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Cellular),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Wifi),
		string(ZebraFotaDeploymentSettingsDownloadRuleNetworkType_WifiAndCellular),
	}
}

func (s *ZebraFotaDeploymentSettingsDownloadRuleNetworkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentSettingsDownloadRuleNetworkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentSettingsDownloadRuleNetworkType(input string) (*ZebraFotaDeploymentSettingsDownloadRuleNetworkType, error) {
	vals := map[string]ZebraFotaDeploymentSettingsDownloadRuleNetworkType{
		"any":             ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Any,
		"cellular":        ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Cellular,
		"wifi":            ZebraFotaDeploymentSettingsDownloadRuleNetworkType_Wifi,
		"wifiandcellular": ZebraFotaDeploymentSettingsDownloadRuleNetworkType_WifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsDownloadRuleNetworkType(input)
	return &out, nil
}
