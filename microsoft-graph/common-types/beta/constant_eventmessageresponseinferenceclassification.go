package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseInferenceClassification string

const (
	EventMessageResponseInferenceClassification_Focused EventMessageResponseInferenceClassification = "focused"
	EventMessageResponseInferenceClassification_Other   EventMessageResponseInferenceClassification = "other"
)

func PossibleValuesForEventMessageResponseInferenceClassification() []string {
	return []string{
		string(EventMessageResponseInferenceClassification_Focused),
		string(EventMessageResponseInferenceClassification_Other),
	}
}

func (s *EventMessageResponseInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageResponseInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageResponseInferenceClassification(input string) (*EventMessageResponseInferenceClassification, error) {
	vals := map[string]EventMessageResponseInferenceClassification{
		"focused": EventMessageResponseInferenceClassification_Focused,
		"other":   EventMessageResponseInferenceClassification_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseInferenceClassification(input)
	return &out, nil
}
