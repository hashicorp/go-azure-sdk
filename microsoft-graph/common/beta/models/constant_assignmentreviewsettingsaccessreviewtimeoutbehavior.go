package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReviewSettingsAccessReviewTimeoutBehavior string

const (
	AssignmentReviewSettingsAccessReviewTimeoutBehavioracceptAccessRecommendation AssignmentReviewSettingsAccessReviewTimeoutBehavior = "AcceptAccessRecommendation"
	AssignmentReviewSettingsAccessReviewTimeoutBehaviorkeepAccess                 AssignmentReviewSettingsAccessReviewTimeoutBehavior = "KeepAccess"
	AssignmentReviewSettingsAccessReviewTimeoutBehaviorremoveAccess               AssignmentReviewSettingsAccessReviewTimeoutBehavior = "RemoveAccess"
)

func PossibleValuesForAssignmentReviewSettingsAccessReviewTimeoutBehavior() []string {
	return []string{
		string(AssignmentReviewSettingsAccessReviewTimeoutBehavioracceptAccessRecommendation),
		string(AssignmentReviewSettingsAccessReviewTimeoutBehaviorkeepAccess),
		string(AssignmentReviewSettingsAccessReviewTimeoutBehaviorremoveAccess),
	}
}

func parseAssignmentReviewSettingsAccessReviewTimeoutBehavior(input string) (*AssignmentReviewSettingsAccessReviewTimeoutBehavior, error) {
	vals := map[string]AssignmentReviewSettingsAccessReviewTimeoutBehavior{
		"acceptaccessrecommendation": AssignmentReviewSettingsAccessReviewTimeoutBehavioracceptAccessRecommendation,
		"keepaccess":                 AssignmentReviewSettingsAccessReviewTimeoutBehaviorkeepAccess,
		"removeaccess":               AssignmentReviewSettingsAccessReviewTimeoutBehaviorremoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentReviewSettingsAccessReviewTimeoutBehavior(input)
	return &out, nil
}
