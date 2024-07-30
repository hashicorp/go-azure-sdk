package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlayPromptOperationCompletionReason string

const (
	PlayPromptOperationCompletionReason_CompletedSuccessfully  PlayPromptOperationCompletionReason = "completedSuccessfully"
	PlayPromptOperationCompletionReason_MediaOperationCanceled PlayPromptOperationCompletionReason = "mediaOperationCanceled"
	PlayPromptOperationCompletionReason_Unknown                PlayPromptOperationCompletionReason = "unknown"
)

func PossibleValuesForPlayPromptOperationCompletionReason() []string {
	return []string{
		string(PlayPromptOperationCompletionReason_CompletedSuccessfully),
		string(PlayPromptOperationCompletionReason_MediaOperationCanceled),
		string(PlayPromptOperationCompletionReason_Unknown),
	}
}

func (s *PlayPromptOperationCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlayPromptOperationCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlayPromptOperationCompletionReason(input string) (*PlayPromptOperationCompletionReason, error) {
	vals := map[string]PlayPromptOperationCompletionReason{
		"completedsuccessfully":  PlayPromptOperationCompletionReason_CompletedSuccessfully,
		"mediaoperationcanceled": PlayPromptOperationCompletionReason_MediaOperationCanceled,
		"unknown":                PlayPromptOperationCompletionReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlayPromptOperationCompletionReason(input)
	return &out, nil
}
