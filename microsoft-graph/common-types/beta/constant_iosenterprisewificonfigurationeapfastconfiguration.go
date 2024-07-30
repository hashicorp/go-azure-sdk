package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationEapFastConfiguration string

const (
	IosEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential                         IosEnterpriseWiFiConfigurationEapFastConfiguration = "noProtectedAccessCredential"
	IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential                        IosEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredential"
	IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision            IosEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvision"
	IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously IosEnterpriseWiFiConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationEapFastConfiguration() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential),
		string(IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential),
		string(IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision),
		string(IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func (s *IosEnterpriseWiFiConfigurationEapFastConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationEapFastConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationEapFastConfiguration(input string) (*IosEnterpriseWiFiConfigurationEapFastConfiguration, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         IosEnterpriseWiFiConfigurationEapFastConfiguration_NoProtectedAccessCredential,
		"useprotectedaccesscredential":                        IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": IosEnterpriseWiFiConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationEapFastConfiguration(input)
	return &out, nil
}
