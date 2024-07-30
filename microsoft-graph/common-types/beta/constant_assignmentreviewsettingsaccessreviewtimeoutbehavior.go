package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReviewSettingsAccessReviewTimeoutBehavior string

const (
	AssignmentReviewSettingsAccessReviewTimeoutBehavior_AcceptAccessRecommendation AssignmentReviewSettingsAccessReviewTimeoutBehavior = "acceptAccessRecommendation"
	AssignmentReviewSettingsAccessReviewTimeoutBehavior_KeepAccess                 AssignmentReviewSettingsAccessReviewTimeoutBehavior = "keepAccess"
	AssignmentReviewSettingsAccessReviewTimeoutBehavior_RemoveAccess               AssignmentReviewSettingsAccessReviewTimeoutBehavior = "removeAccess"
)

func PossibleValuesForAssignmentReviewSettingsAccessReviewTimeoutBehavior() []string {
	return []string{
		string(AssignmentReviewSettingsAccessReviewTimeoutBehavior_AcceptAccessRecommendation),
		string(AssignmentReviewSettingsAccessReviewTimeoutBehavior_KeepAccess),
		string(AssignmentReviewSettingsAccessReviewTimeoutBehavior_RemoveAccess),
	}
}

func (s *AssignmentReviewSettingsAccessReviewTimeoutBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentReviewSettingsAccessReviewTimeoutBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentReviewSettingsAccessReviewTimeoutBehavior(input string) (*AssignmentReviewSettingsAccessReviewTimeoutBehavior, error) {
	vals := map[string]AssignmentReviewSettingsAccessReviewTimeoutBehavior{
		"acceptaccessrecommendation": AssignmentReviewSettingsAccessReviewTimeoutBehavior_AcceptAccessRecommendation,
		"keepaccess":                 AssignmentReviewSettingsAccessReviewTimeoutBehavior_KeepAccess,
		"removeaccess":               AssignmentReviewSettingsAccessReviewTimeoutBehavior_RemoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentReviewSettingsAccessReviewTimeoutBehavior(input)
	return &out, nil
}
