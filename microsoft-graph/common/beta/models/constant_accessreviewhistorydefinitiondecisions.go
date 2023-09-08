package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinitionDecisions string

const (
	AccessReviewHistoryDefinitionDecisionsapprove     AccessReviewHistoryDefinitionDecisions = "Approve"
	AccessReviewHistoryDefinitionDecisionsdeny        AccessReviewHistoryDefinitionDecisions = "Deny"
	AccessReviewHistoryDefinitionDecisionsdontKnow    AccessReviewHistoryDefinitionDecisions = "DontKnow"
	AccessReviewHistoryDefinitionDecisionsnotNotified AccessReviewHistoryDefinitionDecisions = "NotNotified"
	AccessReviewHistoryDefinitionDecisionsnotReviewed AccessReviewHistoryDefinitionDecisions = "NotReviewed"
)

func PossibleValuesForAccessReviewHistoryDefinitionDecisions() []string {
	return []string{
		string(AccessReviewHistoryDefinitionDecisionsapprove),
		string(AccessReviewHistoryDefinitionDecisionsdeny),
		string(AccessReviewHistoryDefinitionDecisionsdontKnow),
		string(AccessReviewHistoryDefinitionDecisionsnotNotified),
		string(AccessReviewHistoryDefinitionDecisionsnotReviewed),
	}
}

func parseAccessReviewHistoryDefinitionDecisions(input string) (*AccessReviewHistoryDefinitionDecisions, error) {
	vals := map[string]AccessReviewHistoryDefinitionDecisions{
		"approve":     AccessReviewHistoryDefinitionDecisionsapprove,
		"deny":        AccessReviewHistoryDefinitionDecisionsdeny,
		"dontknow":    AccessReviewHistoryDefinitionDecisionsdontKnow,
		"notnotified": AccessReviewHistoryDefinitionDecisionsnotNotified,
		"notreviewed": AccessReviewHistoryDefinitionDecisionsnotReviewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryDefinitionDecisions(input)
	return &out, nil
}
