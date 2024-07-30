package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstanceStage string

const (
	CustomExtensionHandlerInstanceStage_AssignmentFourteenDaysBeforeExpiration CustomExtensionHandlerInstanceStage = "assignmentFourteenDaysBeforeExpiration"
	CustomExtensionHandlerInstanceStage_AssignmentOneDayBeforeExpiration       CustomExtensionHandlerInstanceStage = "assignmentOneDayBeforeExpiration"
	CustomExtensionHandlerInstanceStage_AssignmentRequestApproved              CustomExtensionHandlerInstanceStage = "assignmentRequestApproved"
	CustomExtensionHandlerInstanceStage_AssignmentRequestCreated               CustomExtensionHandlerInstanceStage = "assignmentRequestCreated"
	CustomExtensionHandlerInstanceStage_AssignmentRequestGranted               CustomExtensionHandlerInstanceStage = "assignmentRequestGranted"
	CustomExtensionHandlerInstanceStage_AssignmentRequestRemoved               CustomExtensionHandlerInstanceStage = "assignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionHandlerInstanceStage() []string {
	return []string{
		string(CustomExtensionHandlerInstanceStage_AssignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionHandlerInstanceStage_AssignmentOneDayBeforeExpiration),
		string(CustomExtensionHandlerInstanceStage_AssignmentRequestApproved),
		string(CustomExtensionHandlerInstanceStage_AssignmentRequestCreated),
		string(CustomExtensionHandlerInstanceStage_AssignmentRequestGranted),
		string(CustomExtensionHandlerInstanceStage_AssignmentRequestRemoved),
	}
}

func (s *CustomExtensionHandlerInstanceStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomExtensionHandlerInstanceStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomExtensionHandlerInstanceStage(input string) (*CustomExtensionHandlerInstanceStage, error) {
	vals := map[string]CustomExtensionHandlerInstanceStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionHandlerInstanceStage_AssignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionHandlerInstanceStage_AssignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionHandlerInstanceStage_AssignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionHandlerInstanceStage_AssignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionHandlerInstanceStage_AssignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionHandlerInstanceStage_AssignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerInstanceStage(input)
	return &out, nil
}
