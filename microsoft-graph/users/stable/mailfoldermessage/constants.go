package mailfoldermessage

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlagFlagStatus string

const (
	FollowupFlagFlagStatusComplete   FollowupFlagFlagStatus = "complete"
	FollowupFlagFlagStatusFlagged    FollowupFlagFlagStatus = "flagged"
	FollowupFlagFlagStatusNotFlagged FollowupFlagFlagStatus = "notFlagged"
)

func PossibleValuesForFollowupFlagFlagStatus() []string {
	return []string{
		string(FollowupFlagFlagStatusComplete),
		string(FollowupFlagFlagStatusFlagged),
		string(FollowupFlagFlagStatusNotFlagged),
	}
}

func (s *FollowupFlagFlagStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFollowupFlagFlagStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFollowupFlagFlagStatus(input string) (*FollowupFlagFlagStatus, error) {
	vals := map[string]FollowupFlagFlagStatus{
		"complete":   FollowupFlagFlagStatusComplete,
		"flagged":    FollowupFlagFlagStatusFlagged,
		"notflagged": FollowupFlagFlagStatusNotFlagged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FollowupFlagFlagStatus(input)
	return &out, nil
}

type ItemBodyContentType string

const (
	ItemBodyContentTypeHtml ItemBodyContentType = "html"
	ItemBodyContentTypeText ItemBodyContentType = "text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentTypeHtml),
		string(ItemBodyContentTypeText),
	}
}

func (s *ItemBodyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemBodyContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentTypeHtml,
		"text": ItemBodyContentTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}

type MessageImportance string

const (
	MessageImportanceHigh   MessageImportance = "high"
	MessageImportanceLow    MessageImportance = "low"
	MessageImportanceNormal MessageImportance = "normal"
)

func PossibleValuesForMessageImportance() []string {
	return []string{
		string(MessageImportanceHigh),
		string(MessageImportanceLow),
		string(MessageImportanceNormal),
	}
}

func (s *MessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageImportance(input string) (*MessageImportance, error) {
	vals := map[string]MessageImportance{
		"high":   MessageImportanceHigh,
		"low":    MessageImportanceLow,
		"normal": MessageImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageImportance(input)
	return &out, nil
}

type MessageInferenceClassification string

const (
	MessageInferenceClassificationFocused MessageInferenceClassification = "focused"
	MessageInferenceClassificationOther   MessageInferenceClassification = "other"
)

func PossibleValuesForMessageInferenceClassification() []string {
	return []string{
		string(MessageInferenceClassificationFocused),
		string(MessageInferenceClassificationOther),
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
		"focused": MessageInferenceClassificationFocused,
		"other":   MessageInferenceClassificationOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageInferenceClassification(input)
	return &out, nil
}
