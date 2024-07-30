package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode string

const (
	X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateMultiFactor  X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode = "x509CertificateMultiFactor"
	X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateSingleFactor X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode = "x509CertificateSingleFactor"
)

func PossibleValuesForX509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode() []string {
	return []string{
		string(X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateMultiFactor),
		string(X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateSingleFactor),
	}
}

func (s *X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode(input string) (*X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode, error) {
	vals := map[string]X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode{
		"x509certificatemultifactor":  X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateMultiFactor,
		"x509certificatesinglefactor": X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode(input)
	return &out, nil
}
