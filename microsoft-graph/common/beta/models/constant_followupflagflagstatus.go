package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlagFlagStatus string

const (
	FollowupFlagFlagStatuscomplete   FollowupFlagFlagStatus = "Complete"
	FollowupFlagFlagStatusflagged    FollowupFlagFlagStatus = "Flagged"
	FollowupFlagFlagStatusnotFlagged FollowupFlagFlagStatus = "NotFlagged"
)

func PossibleValuesForFollowupFlagFlagStatus() []string {
	return []string{
		string(FollowupFlagFlagStatuscomplete),
		string(FollowupFlagFlagStatusflagged),
		string(FollowupFlagFlagStatusnotFlagged),
	}
}

func parseFollowupFlagFlagStatus(input string) (*FollowupFlagFlagStatus, error) {
	vals := map[string]FollowupFlagFlagStatus{
		"complete":   FollowupFlagFlagStatuscomplete,
		"flagged":    FollowupFlagFlagStatusflagged,
		"notflagged": FollowupFlagFlagStatusnotFlagged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FollowupFlagFlagStatus(input)
	return &out, nil
}
