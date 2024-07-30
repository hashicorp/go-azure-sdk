package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRuleX509CertificateAuthenticationMode string

const (
	X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateMultiFactor  X509CertificateRuleX509CertificateAuthenticationMode = "x509CertificateMultiFactor"
	X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateSingleFactor X509CertificateRuleX509CertificateAuthenticationMode = "x509CertificateSingleFactor"
)

func PossibleValuesForX509CertificateRuleX509CertificateAuthenticationMode() []string {
	return []string{
		string(X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateMultiFactor),
		string(X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateSingleFactor),
	}
}

func (s *X509CertificateRuleX509CertificateAuthenticationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateRuleX509CertificateAuthenticationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateRuleX509CertificateAuthenticationMode(input string) (*X509CertificateRuleX509CertificateAuthenticationMode, error) {
	vals := map[string]X509CertificateRuleX509CertificateAuthenticationMode{
		"x509certificatemultifactor":  X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateMultiFactor,
		"x509certificatesinglefactor": X509CertificateRuleX509CertificateAuthenticationMode_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateRuleX509CertificateAuthenticationMode(input)
	return &out, nil
}
