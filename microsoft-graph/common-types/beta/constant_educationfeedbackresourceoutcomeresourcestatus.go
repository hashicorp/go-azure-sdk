package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedbackResourceOutcomeResourceStatus string

const (
	EducationFeedbackResourceOutcomeResourceStatus_FailedPublish  EducationFeedbackResourceOutcomeResourceStatus = "failedPublish"
	EducationFeedbackResourceOutcomeResourceStatus_NotPublished   EducationFeedbackResourceOutcomeResourceStatus = "notPublished"
	EducationFeedbackResourceOutcomeResourceStatus_PendingPublish EducationFeedbackResourceOutcomeResourceStatus = "pendingPublish"
	EducationFeedbackResourceOutcomeResourceStatus_Published      EducationFeedbackResourceOutcomeResourceStatus = "published"
)

func PossibleValuesForEducationFeedbackResourceOutcomeResourceStatus() []string {
	return []string{
		string(EducationFeedbackResourceOutcomeResourceStatus_FailedPublish),
		string(EducationFeedbackResourceOutcomeResourceStatus_NotPublished),
		string(EducationFeedbackResourceOutcomeResourceStatus_PendingPublish),
		string(EducationFeedbackResourceOutcomeResourceStatus_Published),
	}
}

func (s *EducationFeedbackResourceOutcomeResourceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationFeedbackResourceOutcomeResourceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationFeedbackResourceOutcomeResourceStatus(input string) (*EducationFeedbackResourceOutcomeResourceStatus, error) {
	vals := map[string]EducationFeedbackResourceOutcomeResourceStatus{
		"failedpublish":  EducationFeedbackResourceOutcomeResourceStatus_FailedPublish,
		"notpublished":   EducationFeedbackResourceOutcomeResourceStatus_NotPublished,
		"pendingpublish": EducationFeedbackResourceOutcomeResourceStatus_PendingPublish,
		"published":      EducationFeedbackResourceOutcomeResourceStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationFeedbackResourceOutcomeResourceStatus(input)
	return &out, nil
}
