package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationOverrideClassifyAs string

const (
	InferenceClassificationOverrideClassifyAsfocused InferenceClassificationOverrideClassifyAs = "Focused"
	InferenceClassificationOverrideClassifyAsother   InferenceClassificationOverrideClassifyAs = "Other"
)

func PossibleValuesForInferenceClassificationOverrideClassifyAs() []string {
	return []string{
		string(InferenceClassificationOverrideClassifyAsfocused),
		string(InferenceClassificationOverrideClassifyAsother),
	}
}

func parseInferenceClassificationOverrideClassifyAs(input string) (*InferenceClassificationOverrideClassifyAs, error) {
	vals := map[string]InferenceClassificationOverrideClassifyAs{
		"focused": InferenceClassificationOverrideClassifyAsfocused,
		"other":   InferenceClassificationOverrideClassifyAsother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InferenceClassificationOverrideClassifyAs(input)
	return &out, nil
}
