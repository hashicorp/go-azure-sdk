package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendLabelActionActionSource string

const (
	RecommendLabelActionActionSourceautomatic   RecommendLabelActionActionSource = "Automatic"
	RecommendLabelActionActionSourcedefault     RecommendLabelActionActionSource = "Default"
	RecommendLabelActionActionSourcemanual      RecommendLabelActionActionSource = "Manual"
	RecommendLabelActionActionSourcerecommended RecommendLabelActionActionSource = "Recommended"
)

func PossibleValuesForRecommendLabelActionActionSource() []string {
	return []string{
		string(RecommendLabelActionActionSourceautomatic),
		string(RecommendLabelActionActionSourcedefault),
		string(RecommendLabelActionActionSourcemanual),
		string(RecommendLabelActionActionSourcerecommended),
	}
}

func parseRecommendLabelActionActionSource(input string) (*RecommendLabelActionActionSource, error) {
	vals := map[string]RecommendLabelActionActionSource{
		"automatic":   RecommendLabelActionActionSourceautomatic,
		"default":     RecommendLabelActionActionSourcedefault,
		"manual":      RecommendLabelActionActionSourcemanual,
		"recommended": RecommendLabelActionActionSourcerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendLabelActionActionSource(input)
	return &out, nil
}
