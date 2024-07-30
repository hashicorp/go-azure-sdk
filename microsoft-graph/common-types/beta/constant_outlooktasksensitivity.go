package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskSensitivity string

const (
	OutlookTaskSensitivity_Confidential OutlookTaskSensitivity = "confidential"
	OutlookTaskSensitivity_Normal       OutlookTaskSensitivity = "normal"
	OutlookTaskSensitivity_Personal     OutlookTaskSensitivity = "personal"
	OutlookTaskSensitivity_Private      OutlookTaskSensitivity = "private"
)

func PossibleValuesForOutlookTaskSensitivity() []string {
	return []string{
		string(OutlookTaskSensitivity_Confidential),
		string(OutlookTaskSensitivity_Normal),
		string(OutlookTaskSensitivity_Personal),
		string(OutlookTaskSensitivity_Private),
	}
}

func (s *OutlookTaskSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskSensitivity(input string) (*OutlookTaskSensitivity, error) {
	vals := map[string]OutlookTaskSensitivity{
		"confidential": OutlookTaskSensitivity_Confidential,
		"normal":       OutlookTaskSensitivity_Normal,
		"personal":     OutlookTaskSensitivity_Personal,
		"private":      OutlookTaskSensitivity_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskSensitivity(input)
	return &out, nil
}
