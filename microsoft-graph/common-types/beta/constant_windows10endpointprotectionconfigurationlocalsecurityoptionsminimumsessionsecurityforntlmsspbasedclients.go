package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_None                         Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = "none"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_NtlmV2And128BitEncryption    Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = "ntlmV2And128BitEncryption"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_Require128BitEncryption      Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = "require128BitEncryption"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_RequireNtmlV2SessionSecurity Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = "requireNtmlV2SessionSecurity"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_None),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_NtlmV2And128BitEncryption),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_Require128BitEncryption),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_RequireNtmlV2SessionSecurity),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients{
		"none":                         Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_None,
		"ntlmv2and128bitencryption":    Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_NtlmV2And128BitEncryption,
		"require128bitencryption":      Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_Require128BitEncryption,
		"requirentmlv2sessionsecurity": Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients_RequireNtmlV2SessionSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(input)
	return &out, nil
}
