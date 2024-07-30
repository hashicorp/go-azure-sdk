package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRuleX509CertificateRequiredAffinityLevel string

const (
	X509CertificateRuleX509CertificateRequiredAffinityLevel_High X509CertificateRuleX509CertificateRequiredAffinityLevel = "high"
	X509CertificateRuleX509CertificateRequiredAffinityLevel_Low  X509CertificateRuleX509CertificateRequiredAffinityLevel = "low"
)

func PossibleValuesForX509CertificateRuleX509CertificateRequiredAffinityLevel() []string {
	return []string{
		string(X509CertificateRuleX509CertificateRequiredAffinityLevel_High),
		string(X509CertificateRuleX509CertificateRequiredAffinityLevel_Low),
	}
}

func (s *X509CertificateRuleX509CertificateRequiredAffinityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateRuleX509CertificateRequiredAffinityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateRuleX509CertificateRequiredAffinityLevel(input string) (*X509CertificateRuleX509CertificateRequiredAffinityLevel, error) {
	vals := map[string]X509CertificateRuleX509CertificateRequiredAffinityLevel{
		"high": X509CertificateRuleX509CertificateRequiredAffinityLevel_High,
		"low":  X509CertificateRuleX509CertificateRequiredAffinityLevel_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateRuleX509CertificateRequiredAffinityLevel(input)
	return &out, nil
}
