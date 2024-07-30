package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageHistoryItemActions string

const (
	ChatMessageHistoryItemActions_ActionUndefined ChatMessageHistoryItemActions = "actionUndefined"
	ChatMessageHistoryItemActions_ReactionAdded   ChatMessageHistoryItemActions = "reactionAdded"
	ChatMessageHistoryItemActions_ReactionRemoved ChatMessageHistoryItemActions = "reactionRemoved"
)

func PossibleValuesForChatMessageHistoryItemActions() []string {
	return []string{
		string(ChatMessageHistoryItemActions_ActionUndefined),
		string(ChatMessageHistoryItemActions_ReactionAdded),
		string(ChatMessageHistoryItemActions_ReactionRemoved),
	}
}

func (s *ChatMessageHistoryItemActions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageHistoryItemActions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageHistoryItemActions(input string) (*ChatMessageHistoryItemActions, error) {
	vals := map[string]ChatMessageHistoryItemActions{
		"actionundefined": ChatMessageHistoryItemActions_ActionUndefined,
		"reactionadded":   ChatMessageHistoryItemActions_ReactionAdded,
		"reactionremoved": ChatMessageHistoryItemActions_ReactionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageHistoryItemActions(input)
	return &out, nil
}
