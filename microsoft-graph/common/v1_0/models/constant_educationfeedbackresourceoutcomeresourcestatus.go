package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedbackResourceOutcomeResourceStatus string

const (
	EducationFeedbackResourceOutcomeResourceStatusfailedPublish  EducationFeedbackResourceOutcomeResourceStatus = "FailedPublish"
	EducationFeedbackResourceOutcomeResourceStatusnotPublished   EducationFeedbackResourceOutcomeResourceStatus = "NotPublished"
	EducationFeedbackResourceOutcomeResourceStatuspendingPublish EducationFeedbackResourceOutcomeResourceStatus = "PendingPublish"
	EducationFeedbackResourceOutcomeResourceStatuspublished      EducationFeedbackResourceOutcomeResourceStatus = "Published"
)

func PossibleValuesForEducationFeedbackResourceOutcomeResourceStatus() []string {
	return []string{
		string(EducationFeedbackResourceOutcomeResourceStatusfailedPublish),
		string(EducationFeedbackResourceOutcomeResourceStatusnotPublished),
		string(EducationFeedbackResourceOutcomeResourceStatuspendingPublish),
		string(EducationFeedbackResourceOutcomeResourceStatuspublished),
	}
}

func parseEducationFeedbackResourceOutcomeResourceStatus(input string) (*EducationFeedbackResourceOutcomeResourceStatus, error) {
	vals := map[string]EducationFeedbackResourceOutcomeResourceStatus{
		"failedpublish":  EducationFeedbackResourceOutcomeResourceStatusfailedPublish,
		"notpublished":   EducationFeedbackResourceOutcomeResourceStatusnotPublished,
		"pendingpublish": EducationFeedbackResourceOutcomeResourceStatuspendingPublish,
		"published":      EducationFeedbackResourceOutcomeResourceStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationFeedbackResourceOutcomeResourceStatus(input)
	return &out, nil
}
