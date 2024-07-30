package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentReviewSettingsExpirationBehavior string

const (
	AccessPackageAssignmentReviewSettingsExpirationBehavior_AcceptAccessRecommendation AccessPackageAssignmentReviewSettingsExpirationBehavior = "acceptAccessRecommendation"
	AccessPackageAssignmentReviewSettingsExpirationBehavior_KeepAccess                 AccessPackageAssignmentReviewSettingsExpirationBehavior = "keepAccess"
	AccessPackageAssignmentReviewSettingsExpirationBehavior_RemoveAccess               AccessPackageAssignmentReviewSettingsExpirationBehavior = "removeAccess"
)

func PossibleValuesForAccessPackageAssignmentReviewSettingsExpirationBehavior() []string {
	return []string{
		string(AccessPackageAssignmentReviewSettingsExpirationBehavior_AcceptAccessRecommendation),
		string(AccessPackageAssignmentReviewSettingsExpirationBehavior_KeepAccess),
		string(AccessPackageAssignmentReviewSettingsExpirationBehavior_RemoveAccess),
	}
}

func (s *AccessPackageAssignmentReviewSettingsExpirationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentReviewSettingsExpirationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentReviewSettingsExpirationBehavior(input string) (*AccessPackageAssignmentReviewSettingsExpirationBehavior, error) {
	vals := map[string]AccessPackageAssignmentReviewSettingsExpirationBehavior{
		"acceptaccessrecommendation": AccessPackageAssignmentReviewSettingsExpirationBehavior_AcceptAccessRecommendation,
		"keepaccess":                 AccessPackageAssignmentReviewSettingsExpirationBehavior_KeepAccess,
		"removeaccess":               AccessPackageAssignmentReviewSettingsExpirationBehavior_RemoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentReviewSettingsExpirationBehavior(input)
	return &out, nil
}
