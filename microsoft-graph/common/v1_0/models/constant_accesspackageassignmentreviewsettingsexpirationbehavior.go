package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentReviewSettingsExpirationBehavior string

const (
	AccessPackageAssignmentReviewSettingsExpirationBehavioracceptAccessRecommendation AccessPackageAssignmentReviewSettingsExpirationBehavior = "AcceptAccessRecommendation"
	AccessPackageAssignmentReviewSettingsExpirationBehaviorkeepAccess                 AccessPackageAssignmentReviewSettingsExpirationBehavior = "KeepAccess"
	AccessPackageAssignmentReviewSettingsExpirationBehaviorremoveAccess               AccessPackageAssignmentReviewSettingsExpirationBehavior = "RemoveAccess"
)

func PossibleValuesForAccessPackageAssignmentReviewSettingsExpirationBehavior() []string {
	return []string{
		string(AccessPackageAssignmentReviewSettingsExpirationBehavioracceptAccessRecommendation),
		string(AccessPackageAssignmentReviewSettingsExpirationBehaviorkeepAccess),
		string(AccessPackageAssignmentReviewSettingsExpirationBehaviorremoveAccess),
	}
}

func parseAccessPackageAssignmentReviewSettingsExpirationBehavior(input string) (*AccessPackageAssignmentReviewSettingsExpirationBehavior, error) {
	vals := map[string]AccessPackageAssignmentReviewSettingsExpirationBehavior{
		"acceptaccessrecommendation": AccessPackageAssignmentReviewSettingsExpirationBehavioracceptAccessRecommendation,
		"keepaccess":                 AccessPackageAssignmentReviewSettingsExpirationBehaviorkeepAccess,
		"removeaccess":               AccessPackageAssignmentReviewSettingsExpirationBehaviorremoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentReviewSettingsExpirationBehavior(input)
	return &out, nil
}
