package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageInferenceClassification string

const (
	MessageInferenceClassification_Focused MessageInferenceClassification = "focused"
	MessageInferenceClassification_Other   MessageInferenceClassification = "other"
)

func PossibleValuesForMessageInferenceClassification() []string {
	return []string{
		string(MessageInferenceClassification_Focused),
		string(MessageInferenceClassification_Other),
	}
}

func (s *MessageInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageInferenceClassification(input string) (*MessageInferenceClassification, error) {
	vals := map[string]MessageInferenceClassification{
		"focused": MessageInferenceClassification_Focused,
		"other":   MessageInferenceClassification_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageInferenceClassification(input)
	return &out, nil
}
