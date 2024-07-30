package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel string

const (
	X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_High X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel = "high"
	X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_Low  X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel = "low"
)

func PossibleValuesForX509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel() []string {
	return []string{
		string(X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_High),
		string(X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_Low),
	}
}

func (s *X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel(input string) (*X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel, error) {
	vals := map[string]X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel{
		"high": X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_High,
		"low":  X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAuthenticationModeConfigurationX509CertificateDefaultRequiredAffinityLevel(input)
	return &out, nil
}
