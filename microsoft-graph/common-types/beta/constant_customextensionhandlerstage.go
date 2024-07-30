package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerStage string

const (
	CustomExtensionHandlerStage_AssignmentFourteenDaysBeforeExpiration CustomExtensionHandlerStage = "assignmentFourteenDaysBeforeExpiration"
	CustomExtensionHandlerStage_AssignmentOneDayBeforeExpiration       CustomExtensionHandlerStage = "assignmentOneDayBeforeExpiration"
	CustomExtensionHandlerStage_AssignmentRequestApproved              CustomExtensionHandlerStage = "assignmentRequestApproved"
	CustomExtensionHandlerStage_AssignmentRequestCreated               CustomExtensionHandlerStage = "assignmentRequestCreated"
	CustomExtensionHandlerStage_AssignmentRequestGranted               CustomExtensionHandlerStage = "assignmentRequestGranted"
	CustomExtensionHandlerStage_AssignmentRequestRemoved               CustomExtensionHandlerStage = "assignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionHandlerStage() []string {
	return []string{
		string(CustomExtensionHandlerStage_AssignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionHandlerStage_AssignmentOneDayBeforeExpiration),
		string(CustomExtensionHandlerStage_AssignmentRequestApproved),
		string(CustomExtensionHandlerStage_AssignmentRequestCreated),
		string(CustomExtensionHandlerStage_AssignmentRequestGranted),
		string(CustomExtensionHandlerStage_AssignmentRequestRemoved),
	}
}

func (s *CustomExtensionHandlerStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomExtensionHandlerStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomExtensionHandlerStage(input string) (*CustomExtensionHandlerStage, error) {
	vals := map[string]CustomExtensionHandlerStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionHandlerStage_AssignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionHandlerStage_AssignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionHandlerStage_AssignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionHandlerStage_AssignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionHandlerStage_AssignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionHandlerStage_AssignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerStage(input)
	return &out, nil
}
