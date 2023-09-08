package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelRetentionTrigger string

const (
	SecurityRetentionLabelRetentionTriggerdateCreated  SecurityRetentionLabelRetentionTrigger = "DateCreated"
	SecurityRetentionLabelRetentionTriggerdateLabeled  SecurityRetentionLabelRetentionTrigger = "DateLabeled"
	SecurityRetentionLabelRetentionTriggerdateModified SecurityRetentionLabelRetentionTrigger = "DateModified"
	SecurityRetentionLabelRetentionTriggerdateOfEvent  SecurityRetentionLabelRetentionTrigger = "DateOfEvent"
)

func PossibleValuesForSecurityRetentionLabelRetentionTrigger() []string {
	return []string{
		string(SecurityRetentionLabelRetentionTriggerdateCreated),
		string(SecurityRetentionLabelRetentionTriggerdateLabeled),
		string(SecurityRetentionLabelRetentionTriggerdateModified),
		string(SecurityRetentionLabelRetentionTriggerdateOfEvent),
	}
}

func parseSecurityRetentionLabelRetentionTrigger(input string) (*SecurityRetentionLabelRetentionTrigger, error) {
	vals := map[string]SecurityRetentionLabelRetentionTrigger{
		"datecreated":  SecurityRetentionLabelRetentionTriggerdateCreated,
		"datelabeled":  SecurityRetentionLabelRetentionTriggerdateLabeled,
		"datemodified": SecurityRetentionLabelRetentionTriggerdateModified,
		"dateofevent":  SecurityRetentionLabelRetentionTriggerdateOfEvent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelRetentionTrigger(input)
	return &out, nil
}
