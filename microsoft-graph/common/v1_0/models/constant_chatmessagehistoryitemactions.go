package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageHistoryItemActions string

const (
	ChatMessageHistoryItemActionsactionUndefined ChatMessageHistoryItemActions = "ActionUndefined"
	ChatMessageHistoryItemActionsreactionAdded   ChatMessageHistoryItemActions = "ReactionAdded"
	ChatMessageHistoryItemActionsreactionRemoved ChatMessageHistoryItemActions = "ReactionRemoved"
)

func PossibleValuesForChatMessageHistoryItemActions() []string {
	return []string{
		string(ChatMessageHistoryItemActionsactionUndefined),
		string(ChatMessageHistoryItemActionsreactionAdded),
		string(ChatMessageHistoryItemActionsreactionRemoved),
	}
}

func parseChatMessageHistoryItemActions(input string) (*ChatMessageHistoryItemActions, error) {
	vals := map[string]ChatMessageHistoryItemActions{
		"actionundefined": ChatMessageHistoryItemActionsactionUndefined,
		"reactionadded":   ChatMessageHistoryItemActionsreactionAdded,
		"reactionremoved": ChatMessageHistoryItemActionsreactionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageHistoryItemActions(input)
	return &out, nil
}
