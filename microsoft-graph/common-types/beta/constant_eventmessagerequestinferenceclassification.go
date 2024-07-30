package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestInferenceClassification string

const (
	EventMessageRequestInferenceClassification_Focused EventMessageRequestInferenceClassification = "focused"
	EventMessageRequestInferenceClassification_Other   EventMessageRequestInferenceClassification = "other"
)

func PossibleValuesForEventMessageRequestInferenceClassification() []string {
	return []string{
		string(EventMessageRequestInferenceClassification_Focused),
		string(EventMessageRequestInferenceClassification_Other),
	}
}

func (s *EventMessageRequestInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageRequestInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageRequestInferenceClassification(input string) (*EventMessageRequestInferenceClassification, error) {
	vals := map[string]EventMessageRequestInferenceClassification{
		"focused": EventMessageRequestInferenceClassification_Focused,
		"other":   EventMessageRequestInferenceClassification_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestInferenceClassification(input)
	return &out, nil
}
