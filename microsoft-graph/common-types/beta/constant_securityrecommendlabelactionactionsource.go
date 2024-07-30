package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRecommendLabelActionActionSource string

const (
	SecurityRecommendLabelActionActionSource_Automatic   SecurityRecommendLabelActionActionSource = "automatic"
	SecurityRecommendLabelActionActionSource_Default     SecurityRecommendLabelActionActionSource = "default"
	SecurityRecommendLabelActionActionSource_Manual      SecurityRecommendLabelActionActionSource = "manual"
	SecurityRecommendLabelActionActionSource_Recommended SecurityRecommendLabelActionActionSource = "recommended"
)

func PossibleValuesForSecurityRecommendLabelActionActionSource() []string {
	return []string{
		string(SecurityRecommendLabelActionActionSource_Automatic),
		string(SecurityRecommendLabelActionActionSource_Default),
		string(SecurityRecommendLabelActionActionSource_Manual),
		string(SecurityRecommendLabelActionActionSource_Recommended),
	}
}

func (s *SecurityRecommendLabelActionActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRecommendLabelActionActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRecommendLabelActionActionSource(input string) (*SecurityRecommendLabelActionActionSource, error) {
	vals := map[string]SecurityRecommendLabelActionActionSource{
		"automatic":   SecurityRecommendLabelActionActionSource_Automatic,
		"default":     SecurityRecommendLabelActionActionSource_Default,
		"manual":      SecurityRecommendLabelActionActionSource_Manual,
		"recommended": SecurityRecommendLabelActionActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRecommendLabelActionActionSource(input)
	return &out, nil
}
