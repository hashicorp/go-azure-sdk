package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageInferenceClassification string

const (
	EventMessageInferenceClassification_Focused EventMessageInferenceClassification = "focused"
	EventMessageInferenceClassification_Other   EventMessageInferenceClassification = "other"
)

func PossibleValuesForEventMessageInferenceClassification() []string {
	return []string{
		string(EventMessageInferenceClassification_Focused),
		string(EventMessageInferenceClassification_Other),
	}
}

func (s *EventMessageInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageInferenceClassification(input string) (*EventMessageInferenceClassification, error) {
	vals := map[string]EventMessageInferenceClassification{
		"focused": EventMessageInferenceClassification_Focused,
		"other":   EventMessageInferenceClassification_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageInferenceClassification(input)
	return &out, nil
}
