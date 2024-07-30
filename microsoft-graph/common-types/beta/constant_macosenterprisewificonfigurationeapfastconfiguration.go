package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationEapFastConfiguration string

const (
	MacOSEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential                         MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "noProtectedAccessCredential"
	MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential                        MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredential"
	MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision            MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvision"
	MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationEapFastConfiguration() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func (s *MacOSEnterpriseWiFiConfigurationEapFastConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEnterpriseWiFiConfigurationEapFastConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEnterpriseWiFiConfigurationEapFastConfiguration(input string) (*MacOSEnterpriseWiFiConfigurationEapFastConfiguration, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         MacOSEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential,
		"useprotectedaccesscredential":                        MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": MacOSEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationEapFastConfiguration(input)
	return &out, nil
}
