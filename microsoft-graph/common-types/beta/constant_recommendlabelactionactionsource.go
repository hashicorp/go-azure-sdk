package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendLabelActionActionSource string

const (
	RecommendLabelActionActionSource_Automatic   RecommendLabelActionActionSource = "automatic"
	RecommendLabelActionActionSource_Default     RecommendLabelActionActionSource = "default"
	RecommendLabelActionActionSource_Manual      RecommendLabelActionActionSource = "manual"
	RecommendLabelActionActionSource_Recommended RecommendLabelActionActionSource = "recommended"
)

func PossibleValuesForRecommendLabelActionActionSource() []string {
	return []string{
		string(RecommendLabelActionActionSource_Automatic),
		string(RecommendLabelActionActionSource_Default),
		string(RecommendLabelActionActionSource_Manual),
		string(RecommendLabelActionActionSource_Recommended),
	}
}

func (s *RecommendLabelActionActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendLabelActionActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendLabelActionActionSource(input string) (*RecommendLabelActionActionSource, error) {
	vals := map[string]RecommendLabelActionActionSource{
		"automatic":   RecommendLabelActionActionSource_Automatic,
		"default":     RecommendLabelActionActionSource_Default,
		"manual":      RecommendLabelActionActionSource_Manual,
		"recommended": RecommendLabelActionActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendLabelActionActionSource(input)
	return &out, nil
}
