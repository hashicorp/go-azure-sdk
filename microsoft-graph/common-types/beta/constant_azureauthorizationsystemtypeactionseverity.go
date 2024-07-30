package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAuthorizationSystemTypeActionSeverity string

const (
	AzureAuthorizationSystemTypeActionSeverity_High   AzureAuthorizationSystemTypeActionSeverity = "high"
	AzureAuthorizationSystemTypeActionSeverity_Normal AzureAuthorizationSystemTypeActionSeverity = "normal"
)

func PossibleValuesForAzureAuthorizationSystemTypeActionSeverity() []string {
	return []string{
		string(AzureAuthorizationSystemTypeActionSeverity_High),
		string(AzureAuthorizationSystemTypeActionSeverity_Normal),
	}
}

func (s *AzureAuthorizationSystemTypeActionSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureAuthorizationSystemTypeActionSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureAuthorizationSystemTypeActionSeverity(input string) (*AzureAuthorizationSystemTypeActionSeverity, error) {
	vals := map[string]AzureAuthorizationSystemTypeActionSeverity{
		"high":   AzureAuthorizationSystemTypeActionSeverity_High,
		"normal": AzureAuthorizationSystemTypeActionSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureAuthorizationSystemTypeActionSeverity(input)
	return &out, nil
}
