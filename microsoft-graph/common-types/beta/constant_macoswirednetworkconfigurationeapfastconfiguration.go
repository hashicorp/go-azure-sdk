package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationEapFastConfiguration string

const (
	MacOSWiredNetworkConfigurationEapFastConfiguration_NoProtectedAccessCredential                         MacOSWiredNetworkConfigurationEapFastConfiguration = "noProtectedAccessCredential"
	MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredential                        MacOSWiredNetworkConfigurationEapFastConfiguration = "useProtectedAccessCredential"
	MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision            MacOSWiredNetworkConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvision"
	MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously MacOSWiredNetworkConfigurationEapFastConfiguration = "useProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForMacOSWiredNetworkConfigurationEapFastConfiguration() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationEapFastConfiguration_NoProtectedAccessCredential),
		string(MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredential),
		string(MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision),
		string(MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func (s *MacOSWiredNetworkConfigurationEapFastConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiredNetworkConfigurationEapFastConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiredNetworkConfigurationEapFastConfiguration(input string) (*MacOSWiredNetworkConfigurationEapFastConfiguration, error) {
	vals := map[string]MacOSWiredNetworkConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         MacOSWiredNetworkConfigurationEapFastConfiguration_NoProtectedAccessCredential,
		"useprotectedaccesscredential":                        MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": MacOSWiredNetworkConfigurationEapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationEapFastConfiguration(input)
	return &out, nil
}
