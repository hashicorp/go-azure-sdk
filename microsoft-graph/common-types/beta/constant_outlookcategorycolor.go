package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookCategoryColor string

const (
	OutlookCategoryColor_None     OutlookCategoryColor = "none"
	OutlookCategoryColor_Preset0  OutlookCategoryColor = "preset0"
	OutlookCategoryColor_Preset1  OutlookCategoryColor = "preset1"
	OutlookCategoryColor_Preset10 OutlookCategoryColor = "preset10"
	OutlookCategoryColor_Preset11 OutlookCategoryColor = "preset11"
	OutlookCategoryColor_Preset12 OutlookCategoryColor = "preset12"
	OutlookCategoryColor_Preset13 OutlookCategoryColor = "preset13"
	OutlookCategoryColor_Preset14 OutlookCategoryColor = "preset14"
	OutlookCategoryColor_Preset15 OutlookCategoryColor = "preset15"
	OutlookCategoryColor_Preset16 OutlookCategoryColor = "preset16"
	OutlookCategoryColor_Preset17 OutlookCategoryColor = "preset17"
	OutlookCategoryColor_Preset18 OutlookCategoryColor = "preset18"
	OutlookCategoryColor_Preset19 OutlookCategoryColor = "preset19"
	OutlookCategoryColor_Preset2  OutlookCategoryColor = "preset2"
	OutlookCategoryColor_Preset20 OutlookCategoryColor = "preset20"
	OutlookCategoryColor_Preset21 OutlookCategoryColor = "preset21"
	OutlookCategoryColor_Preset22 OutlookCategoryColor = "preset22"
	OutlookCategoryColor_Preset23 OutlookCategoryColor = "preset23"
	OutlookCategoryColor_Preset24 OutlookCategoryColor = "preset24"
	OutlookCategoryColor_Preset3  OutlookCategoryColor = "preset3"
	OutlookCategoryColor_Preset4  OutlookCategoryColor = "preset4"
	OutlookCategoryColor_Preset5  OutlookCategoryColor = "preset5"
	OutlookCategoryColor_Preset6  OutlookCategoryColor = "preset6"
	OutlookCategoryColor_Preset7  OutlookCategoryColor = "preset7"
	OutlookCategoryColor_Preset8  OutlookCategoryColor = "preset8"
	OutlookCategoryColor_Preset9  OutlookCategoryColor = "preset9"
)

func PossibleValuesForOutlookCategoryColor() []string {
	return []string{
		string(OutlookCategoryColor_None),
		string(OutlookCategoryColor_Preset0),
		string(OutlookCategoryColor_Preset1),
		string(OutlookCategoryColor_Preset10),
		string(OutlookCategoryColor_Preset11),
		string(OutlookCategoryColor_Preset12),
		string(OutlookCategoryColor_Preset13),
		string(OutlookCategoryColor_Preset14),
		string(OutlookCategoryColor_Preset15),
		string(OutlookCategoryColor_Preset16),
		string(OutlookCategoryColor_Preset17),
		string(OutlookCategoryColor_Preset18),
		string(OutlookCategoryColor_Preset19),
		string(OutlookCategoryColor_Preset2),
		string(OutlookCategoryColor_Preset20),
		string(OutlookCategoryColor_Preset21),
		string(OutlookCategoryColor_Preset22),
		string(OutlookCategoryColor_Preset23),
		string(OutlookCategoryColor_Preset24),
		string(OutlookCategoryColor_Preset3),
		string(OutlookCategoryColor_Preset4),
		string(OutlookCategoryColor_Preset5),
		string(OutlookCategoryColor_Preset6),
		string(OutlookCategoryColor_Preset7),
		string(OutlookCategoryColor_Preset8),
		string(OutlookCategoryColor_Preset9),
	}
}

func (s *OutlookCategoryColor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookCategoryColor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookCategoryColor(input string) (*OutlookCategoryColor, error) {
	vals := map[string]OutlookCategoryColor{
		"none":     OutlookCategoryColor_None,
		"preset0":  OutlookCategoryColor_Preset0,
		"preset1":  OutlookCategoryColor_Preset1,
		"preset10": OutlookCategoryColor_Preset10,
		"preset11": OutlookCategoryColor_Preset11,
		"preset12": OutlookCategoryColor_Preset12,
		"preset13": OutlookCategoryColor_Preset13,
		"preset14": OutlookCategoryColor_Preset14,
		"preset15": OutlookCategoryColor_Preset15,
		"preset16": OutlookCategoryColor_Preset16,
		"preset17": OutlookCategoryColor_Preset17,
		"preset18": OutlookCategoryColor_Preset18,
		"preset19": OutlookCategoryColor_Preset19,
		"preset2":  OutlookCategoryColor_Preset2,
		"preset20": OutlookCategoryColor_Preset20,
		"preset21": OutlookCategoryColor_Preset21,
		"preset22": OutlookCategoryColor_Preset22,
		"preset23": OutlookCategoryColor_Preset23,
		"preset24": OutlookCategoryColor_Preset24,
		"preset3":  OutlookCategoryColor_Preset3,
		"preset4":  OutlookCategoryColor_Preset4,
		"preset5":  OutlookCategoryColor_Preset5,
		"preset6":  OutlookCategoryColor_Preset6,
		"preset7":  OutlookCategoryColor_Preset7,
		"preset8":  OutlookCategoryColor_Preset8,
		"preset9":  OutlookCategoryColor_Preset9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookCategoryColor(input)
	return &out, nil
}
