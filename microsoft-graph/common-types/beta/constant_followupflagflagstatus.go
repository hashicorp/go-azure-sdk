package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlagFlagStatus string

const (
	FollowupFlagFlagStatus_Complete   FollowupFlagFlagStatus = "complete"
	FollowupFlagFlagStatus_Flagged    FollowupFlagFlagStatus = "flagged"
	FollowupFlagFlagStatus_NotFlagged FollowupFlagFlagStatus = "notFlagged"
)

func PossibleValuesForFollowupFlagFlagStatus() []string {
	return []string{
		string(FollowupFlagFlagStatus_Complete),
		string(FollowupFlagFlagStatus_Flagged),
		string(FollowupFlagFlagStatus_NotFlagged),
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
		"complete":   FollowupFlagFlagStatus_Complete,
		"flagged":    FollowupFlagFlagStatus_Flagged,
		"notflagged": FollowupFlagFlagStatus_NotFlagged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FollowupFlagFlagStatus(input)
	return &out, nil
}
