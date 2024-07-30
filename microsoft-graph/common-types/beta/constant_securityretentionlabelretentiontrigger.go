package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelRetentionTrigger string

const (
	SecurityRetentionLabelRetentionTrigger_DateCreated  SecurityRetentionLabelRetentionTrigger = "dateCreated"
	SecurityRetentionLabelRetentionTrigger_DateLabeled  SecurityRetentionLabelRetentionTrigger = "dateLabeled"
	SecurityRetentionLabelRetentionTrigger_DateModified SecurityRetentionLabelRetentionTrigger = "dateModified"
	SecurityRetentionLabelRetentionTrigger_DateOfEvent  SecurityRetentionLabelRetentionTrigger = "dateOfEvent"
)

func PossibleValuesForSecurityRetentionLabelRetentionTrigger() []string {
	return []string{
		string(SecurityRetentionLabelRetentionTrigger_DateCreated),
		string(SecurityRetentionLabelRetentionTrigger_DateLabeled),
		string(SecurityRetentionLabelRetentionTrigger_DateModified),
		string(SecurityRetentionLabelRetentionTrigger_DateOfEvent),
	}
}

func (s *SecurityRetentionLabelRetentionTrigger) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionLabelRetentionTrigger(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionLabelRetentionTrigger(input string) (*SecurityRetentionLabelRetentionTrigger, error) {
	vals := map[string]SecurityRetentionLabelRetentionTrigger{
		"datecreated":  SecurityRetentionLabelRetentionTrigger_DateCreated,
		"datelabeled":  SecurityRetentionLabelRetentionTrigger_DateLabeled,
		"datemodified": SecurityRetentionLabelRetentionTrigger_DateModified,
		"dateofevent":  SecurityRetentionLabelRetentionTrigger_DateOfEvent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelRetentionTrigger(input)
	return &out, nil
}
