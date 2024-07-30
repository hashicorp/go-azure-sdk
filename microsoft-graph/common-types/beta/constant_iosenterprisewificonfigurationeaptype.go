package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationEapType string

const (
	IosEnterpriseWiFiConfigurationEapType_EapFast IosEnterpriseWiFiConfigurationEapType = "eapFast"
	IosEnterpriseWiFiConfigurationEapType_EapSim  IosEnterpriseWiFiConfigurationEapType = "eapSim"
	IosEnterpriseWiFiConfigurationEapType_EapTls  IosEnterpriseWiFiConfigurationEapType = "eapTls"
	IosEnterpriseWiFiConfigurationEapType_EapTtls IosEnterpriseWiFiConfigurationEapType = "eapTtls"
	IosEnterpriseWiFiConfigurationEapType_Leap    IosEnterpriseWiFiConfigurationEapType = "leap"
	IosEnterpriseWiFiConfigurationEapType_Peap    IosEnterpriseWiFiConfigurationEapType = "peap"
	IosEnterpriseWiFiConfigurationEapType_Teap    IosEnterpriseWiFiConfigurationEapType = "teap"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationEapType_EapFast),
		string(IosEnterpriseWiFiConfigurationEapType_EapSim),
		string(IosEnterpriseWiFiConfigurationEapType_EapTls),
		string(IosEnterpriseWiFiConfigurationEapType_EapTtls),
		string(IosEnterpriseWiFiConfigurationEapType_Leap),
		string(IosEnterpriseWiFiConfigurationEapType_Peap),
		string(IosEnterpriseWiFiConfigurationEapType_Teap),
	}
}

func (s *IosEnterpriseWiFiConfigurationEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationEapType(input string) (*IosEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationEapType{
		"eapfast": IosEnterpriseWiFiConfigurationEapType_EapFast,
		"eapsim":  IosEnterpriseWiFiConfigurationEapType_EapSim,
		"eaptls":  IosEnterpriseWiFiConfigurationEapType_EapTls,
		"eapttls": IosEnterpriseWiFiConfigurationEapType_EapTtls,
		"leap":    IosEnterpriseWiFiConfigurationEapType_Leap,
		"peap":    IosEnterpriseWiFiConfigurationEapType_Peap,
		"teap":    IosEnterpriseWiFiConfigurationEapType_Teap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
