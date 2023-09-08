package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesMessageActionFlag string

const (
	MessageRulePredicatesMessageActionFlagany                 MessageRulePredicatesMessageActionFlag = "Any"
	MessageRulePredicatesMessageActionFlagcall                MessageRulePredicatesMessageActionFlag = "Call"
	MessageRulePredicatesMessageActionFlagdoNotForward        MessageRulePredicatesMessageActionFlag = "DoNotForward"
	MessageRulePredicatesMessageActionFlagfollowUp            MessageRulePredicatesMessageActionFlag = "FollowUp"
	MessageRulePredicatesMessageActionFlagforward             MessageRulePredicatesMessageActionFlag = "Forward"
	MessageRulePredicatesMessageActionFlagfyi                 MessageRulePredicatesMessageActionFlag = "Fyi"
	MessageRulePredicatesMessageActionFlagnoResponseNecessary MessageRulePredicatesMessageActionFlag = "NoResponseNecessary"
	MessageRulePredicatesMessageActionFlagread                MessageRulePredicatesMessageActionFlag = "Read"
	MessageRulePredicatesMessageActionFlagreply               MessageRulePredicatesMessageActionFlag = "Reply"
	MessageRulePredicatesMessageActionFlagreplyToAll          MessageRulePredicatesMessageActionFlag = "ReplyToAll"
	MessageRulePredicatesMessageActionFlagreview              MessageRulePredicatesMessageActionFlag = "Review"
)

func PossibleValuesForMessageRulePredicatesMessageActionFlag() []string {
	return []string{
		string(MessageRulePredicatesMessageActionFlagany),
		string(MessageRulePredicatesMessageActionFlagcall),
		string(MessageRulePredicatesMessageActionFlagdoNotForward),
		string(MessageRulePredicatesMessageActionFlagfollowUp),
		string(MessageRulePredicatesMessageActionFlagforward),
		string(MessageRulePredicatesMessageActionFlagfyi),
		string(MessageRulePredicatesMessageActionFlagnoResponseNecessary),
		string(MessageRulePredicatesMessageActionFlagread),
		string(MessageRulePredicatesMessageActionFlagreply),
		string(MessageRulePredicatesMessageActionFlagreplyToAll),
		string(MessageRulePredicatesMessageActionFlagreview),
	}
}

func parseMessageRulePredicatesMessageActionFlag(input string) (*MessageRulePredicatesMessageActionFlag, error) {
	vals := map[string]MessageRulePredicatesMessageActionFlag{
		"any":                 MessageRulePredicatesMessageActionFlagany,
		"call":                MessageRulePredicatesMessageActionFlagcall,
		"donotforward":        MessageRulePredicatesMessageActionFlagdoNotForward,
		"followup":            MessageRulePredicatesMessageActionFlagfollowUp,
		"forward":             MessageRulePredicatesMessageActionFlagforward,
		"fyi":                 MessageRulePredicatesMessageActionFlagfyi,
		"noresponsenecessary": MessageRulePredicatesMessageActionFlagnoResponseNecessary,
		"read":                MessageRulePredicatesMessageActionFlagread,
		"reply":               MessageRulePredicatesMessageActionFlagreply,
		"replytoall":          MessageRulePredicatesMessageActionFlagreplyToAll,
		"review":              MessageRulePredicatesMessageActionFlagreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesMessageActionFlag(input)
	return &out, nil
}
