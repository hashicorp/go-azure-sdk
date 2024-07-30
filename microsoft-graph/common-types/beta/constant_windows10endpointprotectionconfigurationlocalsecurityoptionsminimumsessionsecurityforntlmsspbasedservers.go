package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_None                         Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = "none"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_NtlmV2And128BitEncryption    Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = "ntlmV2And128BitEncryption"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_Require128BitEncryption      Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = "require128BitEncryption"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_RequireNtmlV2SessionSecurity Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = "requireNtmlV2SessionSecurity"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_None),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_NtlmV2And128BitEncryption),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_Require128BitEncryption),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_RequireNtmlV2SessionSecurity),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers{
		"none":                         Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_None,
		"ntlmv2and128bitencryption":    Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_NtlmV2And128BitEncryption,
		"require128bitencryption":      Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_Require128BitEncryption,
		"requirentmlv2sessionsecurity": Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers_RequireNtmlV2SessionSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(input)
	return &out, nil
}
