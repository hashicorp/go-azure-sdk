package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Attempt       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "attempt"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_DeviceDefault Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "deviceDefault"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_None          Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "none"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Require       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "require"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Attempt),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_DeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_None),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Require),
	}
}

func (s *Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod{
		"attempt":       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Attempt,
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_DeviceDefault,
		"none":          Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_None,
		"require":       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod_Require,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod(input)
	return &out, nil
}
