package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoredEmailAddressSelectionLikelihood string

const (
	ScoredEmailAddressSelectionLikelihood_High         ScoredEmailAddressSelectionLikelihood = "high"
	ScoredEmailAddressSelectionLikelihood_NotSpecified ScoredEmailAddressSelectionLikelihood = "notSpecified"
)

func PossibleValuesForScoredEmailAddressSelectionLikelihood() []string {
	return []string{
		string(ScoredEmailAddressSelectionLikelihood_High),
		string(ScoredEmailAddressSelectionLikelihood_NotSpecified),
	}
}

func (s *ScoredEmailAddressSelectionLikelihood) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScoredEmailAddressSelectionLikelihood(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScoredEmailAddressSelectionLikelihood(input string) (*ScoredEmailAddressSelectionLikelihood, error) {
	vals := map[string]ScoredEmailAddressSelectionLikelihood{
		"high":         ScoredEmailAddressSelectionLikelihood_High,
		"notspecified": ScoredEmailAddressSelectionLikelihood_NotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScoredEmailAddressSelectionLikelihood(input)
	return &out, nil
}
