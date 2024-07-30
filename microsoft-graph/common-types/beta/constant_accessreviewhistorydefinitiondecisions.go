package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinitionDecisions string

const (
	AccessReviewHistoryDefinitionDecisions_Approve     AccessReviewHistoryDefinitionDecisions = "approve"
	AccessReviewHistoryDefinitionDecisions_Deny        AccessReviewHistoryDefinitionDecisions = "deny"
	AccessReviewHistoryDefinitionDecisions_DontKnow    AccessReviewHistoryDefinitionDecisions = "dontKnow"
	AccessReviewHistoryDefinitionDecisions_NotNotified AccessReviewHistoryDefinitionDecisions = "notNotified"
	AccessReviewHistoryDefinitionDecisions_NotReviewed AccessReviewHistoryDefinitionDecisions = "notReviewed"
)

func PossibleValuesForAccessReviewHistoryDefinitionDecisions() []string {
	return []string{
		string(AccessReviewHistoryDefinitionDecisions_Approve),
		string(AccessReviewHistoryDefinitionDecisions_Deny),
		string(AccessReviewHistoryDefinitionDecisions_DontKnow),
		string(AccessReviewHistoryDefinitionDecisions_NotNotified),
		string(AccessReviewHistoryDefinitionDecisions_NotReviewed),
	}
}

func (s *AccessReviewHistoryDefinitionDecisions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewHistoryDefinitionDecisions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewHistoryDefinitionDecisions(input string) (*AccessReviewHistoryDefinitionDecisions, error) {
	vals := map[string]AccessReviewHistoryDefinitionDecisions{
		"approve":     AccessReviewHistoryDefinitionDecisions_Approve,
		"deny":        AccessReviewHistoryDefinitionDecisions_Deny,
		"dontknow":    AccessReviewHistoryDefinitionDecisions_DontKnow,
		"notnotified": AccessReviewHistoryDefinitionDecisions_NotNotified,
		"notreviewed": AccessReviewHistoryDefinitionDecisions_NotReviewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryDefinitionDecisions(input)
	return &out, nil
}
