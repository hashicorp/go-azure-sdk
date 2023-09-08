package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRecommendLabelActionActionSource string

const (
	SecurityRecommendLabelActionActionSourceautomatic   SecurityRecommendLabelActionActionSource = "Automatic"
	SecurityRecommendLabelActionActionSourcedefault     SecurityRecommendLabelActionActionSource = "Default"
	SecurityRecommendLabelActionActionSourcemanual      SecurityRecommendLabelActionActionSource = "Manual"
	SecurityRecommendLabelActionActionSourcerecommended SecurityRecommendLabelActionActionSource = "Recommended"
)

func PossibleValuesForSecurityRecommendLabelActionActionSource() []string {
	return []string{
		string(SecurityRecommendLabelActionActionSourceautomatic),
		string(SecurityRecommendLabelActionActionSourcedefault),
		string(SecurityRecommendLabelActionActionSourcemanual),
		string(SecurityRecommendLabelActionActionSourcerecommended),
	}
}

func parseSecurityRecommendLabelActionActionSource(input string) (*SecurityRecommendLabelActionActionSource, error) {
	vals := map[string]SecurityRecommendLabelActionActionSource{
		"automatic":   SecurityRecommendLabelActionActionSourceautomatic,
		"default":     SecurityRecommendLabelActionActionSourcedefault,
		"manual":      SecurityRecommendLabelActionActionSourcemanual,
		"recommended": SecurityRecommendLabelActionActionSourcerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRecommendLabelActionActionSource(input)
	return &out, nil
}
