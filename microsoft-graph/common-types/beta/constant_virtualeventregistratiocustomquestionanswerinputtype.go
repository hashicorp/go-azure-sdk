package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistratioCustomQuestionAnswerInputType string

const (
	VirtualEventRegistratioCustomQuestionAnswerInputType_Boolean       VirtualEventRegistratioCustomQuestionAnswerInputType = "boolean"
	VirtualEventRegistratioCustomQuestionAnswerInputType_MultiChoice   VirtualEventRegistratioCustomQuestionAnswerInputType = "multiChoice"
	VirtualEventRegistratioCustomQuestionAnswerInputType_MultilineText VirtualEventRegistratioCustomQuestionAnswerInputType = "multilineText"
	VirtualEventRegistratioCustomQuestionAnswerInputType_SingleChoice  VirtualEventRegistratioCustomQuestionAnswerInputType = "singleChoice"
	VirtualEventRegistratioCustomQuestionAnswerInputType_Text          VirtualEventRegistratioCustomQuestionAnswerInputType = "text"
)

func PossibleValuesForVirtualEventRegistratioCustomQuestionAnswerInputType() []string {
	return []string{
		string(VirtualEventRegistratioCustomQuestionAnswerInputType_Boolean),
		string(VirtualEventRegistratioCustomQuestionAnswerInputType_MultiChoice),
		string(VirtualEventRegistratioCustomQuestionAnswerInputType_MultilineText),
		string(VirtualEventRegistratioCustomQuestionAnswerInputType_SingleChoice),
		string(VirtualEventRegistratioCustomQuestionAnswerInputType_Text),
	}
}

func (s *VirtualEventRegistratioCustomQuestionAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventRegistratioCustomQuestionAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventRegistratioCustomQuestionAnswerInputType(input string) (*VirtualEventRegistratioCustomQuestionAnswerInputType, error) {
	vals := map[string]VirtualEventRegistratioCustomQuestionAnswerInputType{
		"boolean":       VirtualEventRegistratioCustomQuestionAnswerInputType_Boolean,
		"multichoice":   VirtualEventRegistratioCustomQuestionAnswerInputType_MultiChoice,
		"multilinetext": VirtualEventRegistratioCustomQuestionAnswerInputType_MultilineText,
		"singlechoice":  VirtualEventRegistratioCustomQuestionAnswerInputType_SingleChoice,
		"text":          VirtualEventRegistratioCustomQuestionAnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistratioCustomQuestionAnswerInputType(input)
	return &out, nil
}
