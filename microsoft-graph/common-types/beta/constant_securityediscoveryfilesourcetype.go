package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryFileSourceType string

const (
	SecurityEdiscoveryFileSourceType_Mailbox SecurityEdiscoveryFileSourceType = "mailbox"
	SecurityEdiscoveryFileSourceType_Site    SecurityEdiscoveryFileSourceType = "site"
)

func PossibleValuesForSecurityEdiscoveryFileSourceType() []string {
	return []string{
		string(SecurityEdiscoveryFileSourceType_Mailbox),
		string(SecurityEdiscoveryFileSourceType_Site),
	}
}

func (s *SecurityEdiscoveryFileSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryFileSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryFileSourceType(input string) (*SecurityEdiscoveryFileSourceType, error) {
	vals := map[string]SecurityEdiscoveryFileSourceType{
		"mailbox": SecurityEdiscoveryFileSourceType_Mailbox,
		"site":    SecurityEdiscoveryFileSourceType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryFileSourceType(input)
	return &out, nil
}
