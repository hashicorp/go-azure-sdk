package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoredEmailAddressSelectionLikelihood string

const (
	ScoredEmailAddressSelectionLikelihoodhigh         ScoredEmailAddressSelectionLikelihood = "High"
	ScoredEmailAddressSelectionLikelihoodnotSpecified ScoredEmailAddressSelectionLikelihood = "NotSpecified"
)

func PossibleValuesForScoredEmailAddressSelectionLikelihood() []string {
	return []string{
		string(ScoredEmailAddressSelectionLikelihoodhigh),
		string(ScoredEmailAddressSelectionLikelihoodnotSpecified),
	}
}

func parseScoredEmailAddressSelectionLikelihood(input string) (*ScoredEmailAddressSelectionLikelihood, error) {
	vals := map[string]ScoredEmailAddressSelectionLikelihood{
		"high":         ScoredEmailAddressSelectionLikelihoodhigh,
		"notspecified": ScoredEmailAddressSelectionLikelihoodnotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScoredEmailAddressSelectionLikelihood(input)
	return &out, nil
}
