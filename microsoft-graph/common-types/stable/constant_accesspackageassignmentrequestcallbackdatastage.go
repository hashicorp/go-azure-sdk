package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestCallbackDataStage string

const (
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentFourteenDaysBeforeExpiration AccessPackageAssignmentRequestCallbackDataStage = "assignmentFourteenDaysBeforeExpiration"
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentOneDayBeforeExpiration       AccessPackageAssignmentRequestCallbackDataStage = "assignmentOneDayBeforeExpiration"
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestApproved              AccessPackageAssignmentRequestCallbackDataStage = "assignmentRequestApproved"
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestCreated               AccessPackageAssignmentRequestCallbackDataStage = "assignmentRequestCreated"
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestGranted               AccessPackageAssignmentRequestCallbackDataStage = "assignmentRequestGranted"
	AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestRemoved               AccessPackageAssignmentRequestCallbackDataStage = "assignmentRequestRemoved"
)

func PossibleValuesForAccessPackageAssignmentRequestCallbackDataStage() []string {
	return []string{
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentFourteenDaysBeforeExpiration),
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentOneDayBeforeExpiration),
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestApproved),
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestCreated),
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestGranted),
		string(AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestRemoved),
	}
}

func (s *AccessPackageAssignmentRequestCallbackDataStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentRequestCallbackDataStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentRequestCallbackDataStage(input string) (*AccessPackageAssignmentRequestCallbackDataStage, error) {
	vals := map[string]AccessPackageAssignmentRequestCallbackDataStage{
		"assignmentfourteendaysbeforeexpiration": AccessPackageAssignmentRequestCallbackDataStage_AssignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       AccessPackageAssignmentRequestCallbackDataStage_AssignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestApproved,
		"assignmentrequestcreated":               AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestCreated,
		"assignmentrequestgranted":               AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestGranted,
		"assignmentrequestremoved":               AccessPackageAssignmentRequestCallbackDataStage_AssignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestCallbackDataStage(input)
	return &out, nil
}
