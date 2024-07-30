package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestStatus string

const (
	MailAssessmentRequestStatus_Completed MailAssessmentRequestStatus = "completed"
	MailAssessmentRequestStatus_Pending   MailAssessmentRequestStatus = "pending"
)

func PossibleValuesForMailAssessmentRequestStatus() []string {
	return []string{
		string(MailAssessmentRequestStatus_Completed),
		string(MailAssessmentRequestStatus_Pending),
	}
}

func (s *MailAssessmentRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestStatus(input string) (*MailAssessmentRequestStatus, error) {
	vals := map[string]MailAssessmentRequestStatus{
		"completed": MailAssessmentRequestStatus_Completed,
		"pending":   MailAssessmentRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestStatus(input)
	return &out, nil
}
