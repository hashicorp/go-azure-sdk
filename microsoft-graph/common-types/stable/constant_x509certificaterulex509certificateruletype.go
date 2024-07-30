package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRuleX509CertificateRuleType string

const (
	X509CertificateRuleX509CertificateRuleType_IssuerSubject X509CertificateRuleX509CertificateRuleType = "issuerSubject"
	X509CertificateRuleX509CertificateRuleType_PolicyOID     X509CertificateRuleX509CertificateRuleType = "policyOID"
)

func PossibleValuesForX509CertificateRuleX509CertificateRuleType() []string {
	return []string{
		string(X509CertificateRuleX509CertificateRuleType_IssuerSubject),
		string(X509CertificateRuleX509CertificateRuleType_PolicyOID),
	}
}

func (s *X509CertificateRuleX509CertificateRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateRuleX509CertificateRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateRuleX509CertificateRuleType(input string) (*X509CertificateRuleX509CertificateRuleType, error) {
	vals := map[string]X509CertificateRuleX509CertificateRuleType{
		"issuersubject": X509CertificateRuleX509CertificateRuleType_IssuerSubject,
		"policyoid":     X509CertificateRuleX509CertificateRuleType_PolicyOID,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateRuleX509CertificateRuleType(input)
	return &out, nil
}
