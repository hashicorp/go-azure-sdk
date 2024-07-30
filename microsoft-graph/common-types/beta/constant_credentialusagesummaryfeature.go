package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummaryFeature string

const (
	CredentialUsageSummaryFeature_Registration CredentialUsageSummaryFeature = "registration"
	CredentialUsageSummaryFeature_Reset        CredentialUsageSummaryFeature = "reset"
)

func PossibleValuesForCredentialUsageSummaryFeature() []string {
	return []string{
		string(CredentialUsageSummaryFeature_Registration),
		string(CredentialUsageSummaryFeature_Reset),
	}
}

func (s *CredentialUsageSummaryFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialUsageSummaryFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialUsageSummaryFeature(input string) (*CredentialUsageSummaryFeature, error) {
	vals := map[string]CredentialUsageSummaryFeature{
		"registration": CredentialUsageSummaryFeature_Registration,
		"reset":        CredentialUsageSummaryFeature_Reset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUsageSummaryFeature(input)
	return &out, nil
}
