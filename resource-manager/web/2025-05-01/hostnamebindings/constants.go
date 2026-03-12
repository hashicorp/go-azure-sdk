package hostnamebindings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceType string

const (
	AzureResourceTypeTrafficManager AzureResourceType = "TrafficManager"
	AzureResourceTypeWebsite        AzureResourceType = "Website"
)

func PossibleValuesForAzureResourceType() []string {
	return []string{
		string(AzureResourceTypeTrafficManager),
		string(AzureResourceTypeWebsite),
	}
}

func (s *AzureResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureResourceType(input string) (*AzureResourceType, error) {
	vals := map[string]AzureResourceType{
		"trafficmanager": AzureResourceTypeTrafficManager,
		"website":        AzureResourceTypeWebsite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureResourceType(input)
	return &out, nil
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeA     CustomHostNameDnsRecordType = "A"
	CustomHostNameDnsRecordTypeCName CustomHostNameDnsRecordType = "CName"
)

func PossibleValuesForCustomHostNameDnsRecordType() []string {
	return []string{
		string(CustomHostNameDnsRecordTypeA),
		string(CustomHostNameDnsRecordTypeCName),
	}
}

func (s *CustomHostNameDnsRecordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomHostNameDnsRecordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomHostNameDnsRecordType(input string) (*CustomHostNameDnsRecordType, error) {
	vals := map[string]CustomHostNameDnsRecordType{
		"a":     CustomHostNameDnsRecordTypeA,
		"cname": CustomHostNameDnsRecordTypeCName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomHostNameDnsRecordType(input)
	return &out, nil
}

type HostNameType string

const (
	HostNameTypeManaged  HostNameType = "Managed"
	HostNameTypeVerified HostNameType = "Verified"
)

func PossibleValuesForHostNameType() []string {
	return []string{
		string(HostNameTypeManaged),
		string(HostNameTypeVerified),
	}
}

func (s *HostNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostNameType(input string) (*HostNameType, error) {
	vals := map[string]HostNameType{
		"managed":  HostNameTypeManaged,
		"verified": HostNameTypeVerified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostNameType(input)
	return &out, nil
}

type SslState string

const (
	SslStateDisabled       SslState = "Disabled"
	SslStateIPBasedEnabled SslState = "IpBasedEnabled"
	SslStateSniEnabled     SslState = "SniEnabled"
)

func PossibleValuesForSslState() []string {
	return []string{
		string(SslStateDisabled),
		string(SslStateIPBasedEnabled),
		string(SslStateSniEnabled),
	}
}

func (s *SslState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSslState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSslState(input string) (*SslState, error) {
	vals := map[string]SslState{
		"disabled":       SslStateDisabled,
		"ipbasedenabled": SslStateIPBasedEnabled,
		"snienabled":     SslStateSniEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SslState(input)
	return &out, nil
}
