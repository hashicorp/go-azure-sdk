package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateUserBindingTrustAffinityLevel string

const (
	X509CertificateUserBindingTrustAffinityLevel_High X509CertificateUserBindingTrustAffinityLevel = "high"
	X509CertificateUserBindingTrustAffinityLevel_Low  X509CertificateUserBindingTrustAffinityLevel = "low"
)

func PossibleValuesForX509CertificateUserBindingTrustAffinityLevel() []string {
	return []string{
		string(X509CertificateUserBindingTrustAffinityLevel_High),
		string(X509CertificateUserBindingTrustAffinityLevel_Low),
	}
}

func (s *X509CertificateUserBindingTrustAffinityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateUserBindingTrustAffinityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateUserBindingTrustAffinityLevel(input string) (*X509CertificateUserBindingTrustAffinityLevel, error) {
	vals := map[string]X509CertificateUserBindingTrustAffinityLevel{
		"high": X509CertificateUserBindingTrustAffinityLevel_High,
		"low":  X509CertificateUserBindingTrustAffinityLevel_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateUserBindingTrustAffinityLevel(input)
	return &out, nil
}
