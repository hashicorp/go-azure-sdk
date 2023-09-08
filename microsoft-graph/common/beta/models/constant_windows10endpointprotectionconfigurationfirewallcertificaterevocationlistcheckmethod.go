package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodattempt       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "Attempt"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethoddeviceDefault Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "DeviceDefault"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodnone          Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "None"
	Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodrequire       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod = "Require"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodattempt),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethoddeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodnone),
		string(Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodrequire),
	}
}

func parseWindows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod{
		"attempt":       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodattempt,
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethoddeviceDefault,
		"none":          Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodnone,
		"require":       Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethodrequire,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod(input)
	return &out, nil
}
