package locations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaEnabled string

const (
	QuotaEnabledDisabled QuotaEnabled = "Disabled"
	QuotaEnabledEnabled  QuotaEnabled = "Enabled"
)

func PossibleValuesForQuotaEnabled() []string {
	return []string{
		string(QuotaEnabledDisabled),
		string(QuotaEnabledEnabled),
	}
}

func (s *QuotaEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQuotaEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQuotaEnabled(input string) (*QuotaEnabled, error) {
	vals := map[string]QuotaEnabled{
		"disabled": QuotaEnabledDisabled,
		"enabled":  QuotaEnabledEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QuotaEnabled(input)
	return &out, nil
}

type TrialStatus string

const (
	TrialStatusTrialAvailable TrialStatus = "TrialAvailable"
	TrialStatusTrialDisabled  TrialStatus = "TrialDisabled"
	TrialStatusTrialUsed      TrialStatus = "TrialUsed"
)

func PossibleValuesForTrialStatus() []string {
	return []string{
		string(TrialStatusTrialAvailable),
		string(TrialStatusTrialDisabled),
		string(TrialStatusTrialUsed),
	}
}

func (s *TrialStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrialStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrialStatus(input string) (*TrialStatus, error) {
	vals := map[string]TrialStatus{
		"trialavailable": TrialStatusTrialAvailable,
		"trialdisabled":  TrialStatusTrialDisabled,
		"trialused":      TrialStatusTrialUsed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrialStatus(input)
	return &out, nil
}
