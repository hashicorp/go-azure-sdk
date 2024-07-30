package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlockFileResponseActionIdentifier string

const (
	SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA1   SecurityBlockFileResponseActionIdentifier = "initiatingProcessSHA1"
	SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA256 SecurityBlockFileResponseActionIdentifier = "initiatingProcessSHA256"
	SecurityBlockFileResponseActionIdentifier_Sha1                    SecurityBlockFileResponseActionIdentifier = "sha1"
	SecurityBlockFileResponseActionIdentifier_Sha256                  SecurityBlockFileResponseActionIdentifier = "sha256"
)

func PossibleValuesForSecurityBlockFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA1),
		string(SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA256),
		string(SecurityBlockFileResponseActionIdentifier_Sha1),
		string(SecurityBlockFileResponseActionIdentifier_Sha256),
	}
}

func (s *SecurityBlockFileResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlockFileResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlockFileResponseActionIdentifier(input string) (*SecurityBlockFileResponseActionIdentifier, error) {
	vals := map[string]SecurityBlockFileResponseActionIdentifier{
		"initiatingprocesssha1":   SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA1,
		"initiatingprocesssha256": SecurityBlockFileResponseActionIdentifier_InitiatingProcessSHA256,
		"sha1":                    SecurityBlockFileResponseActionIdentifier_Sha1,
		"sha256":                  SecurityBlockFileResponseActionIdentifier_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlockFileResponseActionIdentifier(input)
	return &out, nil
}
