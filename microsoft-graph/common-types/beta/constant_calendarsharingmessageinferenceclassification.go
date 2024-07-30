package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageInferenceClassification string

const (
	CalendarSharingMessageInferenceClassification_Focused CalendarSharingMessageInferenceClassification = "focused"
	CalendarSharingMessageInferenceClassification_Other   CalendarSharingMessageInferenceClassification = "other"
)

func PossibleValuesForCalendarSharingMessageInferenceClassification() []string {
	return []string{
		string(CalendarSharingMessageInferenceClassification_Focused),
		string(CalendarSharingMessageInferenceClassification_Other),
	}
}

func (s *CalendarSharingMessageInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingMessageInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingMessageInferenceClassification(input string) (*CalendarSharingMessageInferenceClassification, error) {
	vals := map[string]CalendarSharingMessageInferenceClassification{
		"focused": CalendarSharingMessageInferenceClassification_Focused,
		"other":   CalendarSharingMessageInferenceClassification_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageInferenceClassification(input)
	return &out, nil
}
