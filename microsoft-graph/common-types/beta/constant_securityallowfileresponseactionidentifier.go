package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAllowFileResponseActionIdentifier string

const (
	SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA1   SecurityAllowFileResponseActionIdentifier = "initiatingProcessSHA1"
	SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA256 SecurityAllowFileResponseActionIdentifier = "initiatingProcessSHA256"
	SecurityAllowFileResponseActionIdentifier_Sha1                    SecurityAllowFileResponseActionIdentifier = "sha1"
	SecurityAllowFileResponseActionIdentifier_Sha256                  SecurityAllowFileResponseActionIdentifier = "sha256"
)

func PossibleValuesForSecurityAllowFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA1),
		string(SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA256),
		string(SecurityAllowFileResponseActionIdentifier_Sha1),
		string(SecurityAllowFileResponseActionIdentifier_Sha256),
	}
}

func (s *SecurityAllowFileResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAllowFileResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAllowFileResponseActionIdentifier(input string) (*SecurityAllowFileResponseActionIdentifier, error) {
	vals := map[string]SecurityAllowFileResponseActionIdentifier{
		"initiatingprocesssha1":   SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA1,
		"initiatingprocesssha256": SecurityAllowFileResponseActionIdentifier_InitiatingProcessSHA256,
		"sha1":                    SecurityAllowFileResponseActionIdentifier_Sha1,
		"sha256":                  SecurityAllowFileResponseActionIdentifier_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAllowFileResponseActionIdentifier(input)
	return &out, nil
}
