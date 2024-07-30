package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesMessageActionFlag string

const (
	MessageRulePredicatesMessageActionFlag_Any                 MessageRulePredicatesMessageActionFlag = "any"
	MessageRulePredicatesMessageActionFlag_Call                MessageRulePredicatesMessageActionFlag = "call"
	MessageRulePredicatesMessageActionFlag_DoNotForward        MessageRulePredicatesMessageActionFlag = "doNotForward"
	MessageRulePredicatesMessageActionFlag_FollowUp            MessageRulePredicatesMessageActionFlag = "followUp"
	MessageRulePredicatesMessageActionFlag_Forward             MessageRulePredicatesMessageActionFlag = "forward"
	MessageRulePredicatesMessageActionFlag_Fyi                 MessageRulePredicatesMessageActionFlag = "fyi"
	MessageRulePredicatesMessageActionFlag_NoResponseNecessary MessageRulePredicatesMessageActionFlag = "noResponseNecessary"
	MessageRulePredicatesMessageActionFlag_Read                MessageRulePredicatesMessageActionFlag = "read"
	MessageRulePredicatesMessageActionFlag_Reply               MessageRulePredicatesMessageActionFlag = "reply"
	MessageRulePredicatesMessageActionFlag_ReplyToAll          MessageRulePredicatesMessageActionFlag = "replyToAll"
	MessageRulePredicatesMessageActionFlag_Review              MessageRulePredicatesMessageActionFlag = "review"
)

func PossibleValuesForMessageRulePredicatesMessageActionFlag() []string {
	return []string{
		string(MessageRulePredicatesMessageActionFlag_Any),
		string(MessageRulePredicatesMessageActionFlag_Call),
		string(MessageRulePredicatesMessageActionFlag_DoNotForward),
		string(MessageRulePredicatesMessageActionFlag_FollowUp),
		string(MessageRulePredicatesMessageActionFlag_Forward),
		string(MessageRulePredicatesMessageActionFlag_Fyi),
		string(MessageRulePredicatesMessageActionFlag_NoResponseNecessary),
		string(MessageRulePredicatesMessageActionFlag_Read),
		string(MessageRulePredicatesMessageActionFlag_Reply),
		string(MessageRulePredicatesMessageActionFlag_ReplyToAll),
		string(MessageRulePredicatesMessageActionFlag_Review),
	}
}

func (s *MessageRulePredicatesMessageActionFlag) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesMessageActionFlag(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesMessageActionFlag(input string) (*MessageRulePredicatesMessageActionFlag, error) {
	vals := map[string]MessageRulePredicatesMessageActionFlag{
		"any":                 MessageRulePredicatesMessageActionFlag_Any,
		"call":                MessageRulePredicatesMessageActionFlag_Call,
		"donotforward":        MessageRulePredicatesMessageActionFlag_DoNotForward,
		"followup":            MessageRulePredicatesMessageActionFlag_FollowUp,
		"forward":             MessageRulePredicatesMessageActionFlag_Forward,
		"fyi":                 MessageRulePredicatesMessageActionFlag_Fyi,
		"noresponsenecessary": MessageRulePredicatesMessageActionFlag_NoResponseNecessary,
		"read":                MessageRulePredicatesMessageActionFlag_Read,
		"reply":               MessageRulePredicatesMessageActionFlag_Reply,
		"replytoall":          MessageRulePredicatesMessageActionFlag_ReplyToAll,
		"review":              MessageRulePredicatesMessageActionFlag_Review,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesMessageActionFlag(input)
	return &out, nil
}
