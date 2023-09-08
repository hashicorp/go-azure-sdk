package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswerInputType string

const (
	VirtualEventRegistrationQuestionAnswerInputTypeboolean       VirtualEventRegistrationQuestionAnswerInputType = "Boolean"
	VirtualEventRegistrationQuestionAnswerInputTypemultiChoice   VirtualEventRegistrationQuestionAnswerInputType = "MultiChoice"
	VirtualEventRegistrationQuestionAnswerInputTypemultilineText VirtualEventRegistrationQuestionAnswerInputType = "MultilineText"
	VirtualEventRegistrationQuestionAnswerInputTypesingleChoice  VirtualEventRegistrationQuestionAnswerInputType = "SingleChoice"
	VirtualEventRegistrationQuestionAnswerInputTypetext          VirtualEventRegistrationQuestionAnswerInputType = "Text"
)

func PossibleValuesForVirtualEventRegistrationQuestionAnswerInputType() []string {
	return []string{
		string(VirtualEventRegistrationQuestionAnswerInputTypeboolean),
		string(VirtualEventRegistrationQuestionAnswerInputTypemultiChoice),
		string(VirtualEventRegistrationQuestionAnswerInputTypemultilineText),
		string(VirtualEventRegistrationQuestionAnswerInputTypesingleChoice),
		string(VirtualEventRegistrationQuestionAnswerInputTypetext),
	}
}

func parseVirtualEventRegistrationQuestionAnswerInputType(input string) (*VirtualEventRegistrationQuestionAnswerInputType, error) {
	vals := map[string]VirtualEventRegistrationQuestionAnswerInputType{
		"boolean":       VirtualEventRegistrationQuestionAnswerInputTypeboolean,
		"multichoice":   VirtualEventRegistrationQuestionAnswerInputTypemultiChoice,
		"multilinetext": VirtualEventRegistrationQuestionAnswerInputTypemultilineText,
		"singlechoice":  VirtualEventRegistrationQuestionAnswerInputTypesingleChoice,
		"text":          VirtualEventRegistrationQuestionAnswerInputTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationQuestionAnswerInputType(input)
	return &out, nil
}
