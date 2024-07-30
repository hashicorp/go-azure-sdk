package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationLocalIdentifier string

const (
	IosikEv2VpnConfigurationLocalIdentifier_ClientCertificateSubjectName IosikEv2VpnConfigurationLocalIdentifier = "clientCertificateSubjectName"
	IosikEv2VpnConfigurationLocalIdentifier_DeviceFQDN                   IosikEv2VpnConfigurationLocalIdentifier = "deviceFQDN"
	IosikEv2VpnConfigurationLocalIdentifier_Empty                        IosikEv2VpnConfigurationLocalIdentifier = "empty"
)

func PossibleValuesForIosikEv2VpnConfigurationLocalIdentifier() []string {
	return []string{
		string(IosikEv2VpnConfigurationLocalIdentifier_ClientCertificateSubjectName),
		string(IosikEv2VpnConfigurationLocalIdentifier_DeviceFQDN),
		string(IosikEv2VpnConfigurationLocalIdentifier_Empty),
	}
}

func (s *IosikEv2VpnConfigurationLocalIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationLocalIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationLocalIdentifier(input string) (*IosikEv2VpnConfigurationLocalIdentifier, error) {
	vals := map[string]IosikEv2VpnConfigurationLocalIdentifier{
		"clientcertificatesubjectname": IosikEv2VpnConfigurationLocalIdentifier_ClientCertificateSubjectName,
		"devicefqdn":                   IosikEv2VpnConfigurationLocalIdentifier_DeviceFQDN,
		"empty":                        IosikEv2VpnConfigurationLocalIdentifier_Empty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationLocalIdentifier(input)
	return &out, nil
}
