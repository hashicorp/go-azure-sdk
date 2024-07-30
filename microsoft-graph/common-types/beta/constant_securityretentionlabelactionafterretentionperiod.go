package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelActionAfterRetentionPeriod string

const (
	SecurityRetentionLabelActionAfterRetentionPeriod_Delete                 SecurityRetentionLabelActionAfterRetentionPeriod = "delete"
	SecurityRetentionLabelActionAfterRetentionPeriod_None                   SecurityRetentionLabelActionAfterRetentionPeriod = "none"
	SecurityRetentionLabelActionAfterRetentionPeriod_Relabel                SecurityRetentionLabelActionAfterRetentionPeriod = "relabel"
	SecurityRetentionLabelActionAfterRetentionPeriod_StartDispositionReview SecurityRetentionLabelActionAfterRetentionPeriod = "startDispositionReview"
)

func PossibleValuesForSecurityRetentionLabelActionAfterRetentionPeriod() []string {
	return []string{
		string(SecurityRetentionLabelActionAfterRetentionPeriod_Delete),
		string(SecurityRetentionLabelActionAfterRetentionPeriod_None),
		string(SecurityRetentionLabelActionAfterRetentionPeriod_Relabel),
		string(SecurityRetentionLabelActionAfterRetentionPeriod_StartDispositionReview),
	}
}

func (s *SecurityRetentionLabelActionAfterRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionLabelActionAfterRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionLabelActionAfterRetentionPeriod(input string) (*SecurityRetentionLabelActionAfterRetentionPeriod, error) {
	vals := map[string]SecurityRetentionLabelActionAfterRetentionPeriod{
		"delete":                 SecurityRetentionLabelActionAfterRetentionPeriod_Delete,
		"none":                   SecurityRetentionLabelActionAfterRetentionPeriod_None,
		"relabel":                SecurityRetentionLabelActionAfterRetentionPeriod_Relabel,
		"startdispositionreview": SecurityRetentionLabelActionAfterRetentionPeriod_StartDispositionReview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelActionAfterRetentionPeriod(input)
	return &out, nil
}
