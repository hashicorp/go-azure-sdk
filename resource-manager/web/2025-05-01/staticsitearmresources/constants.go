package staticsitearmresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseGradeCdnStatus string

const (
	EnterpriseGradeCdnStatusDisabled  EnterpriseGradeCdnStatus = "Disabled"
	EnterpriseGradeCdnStatusDisabling EnterpriseGradeCdnStatus = "Disabling"
	EnterpriseGradeCdnStatusEnabled   EnterpriseGradeCdnStatus = "Enabled"
	EnterpriseGradeCdnStatusEnabling  EnterpriseGradeCdnStatus = "Enabling"
)

func PossibleValuesForEnterpriseGradeCdnStatus() []string {
	return []string{
		string(EnterpriseGradeCdnStatusDisabled),
		string(EnterpriseGradeCdnStatusDisabling),
		string(EnterpriseGradeCdnStatusEnabled),
		string(EnterpriseGradeCdnStatusEnabling),
	}
}

func (s *EnterpriseGradeCdnStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnterpriseGradeCdnStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnterpriseGradeCdnStatus(input string) (*EnterpriseGradeCdnStatus, error) {
	vals := map[string]EnterpriseGradeCdnStatus{
		"disabled":  EnterpriseGradeCdnStatusDisabled,
		"disabling": EnterpriseGradeCdnStatusDisabling,
		"enabled":   EnterpriseGradeCdnStatusEnabled,
		"enabling":  EnterpriseGradeCdnStatusEnabling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnterpriseGradeCdnStatus(input)
	return &out, nil
}

type StagingEnvironmentPolicy string

const (
	StagingEnvironmentPolicyDisabled StagingEnvironmentPolicy = "Disabled"
	StagingEnvironmentPolicyEnabled  StagingEnvironmentPolicy = "Enabled"
)

func PossibleValuesForStagingEnvironmentPolicy() []string {
	return []string{
		string(StagingEnvironmentPolicyDisabled),
		string(StagingEnvironmentPolicyEnabled),
	}
}

func (s *StagingEnvironmentPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStagingEnvironmentPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStagingEnvironmentPolicy(input string) (*StagingEnvironmentPolicy, error) {
	vals := map[string]StagingEnvironmentPolicy{
		"disabled": StagingEnvironmentPolicyDisabled,
		"enabled":  StagingEnvironmentPolicyEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StagingEnvironmentPolicy(input)
	return &out, nil
}

type TriggerTypes string

const (
	TriggerTypesHTTPTrigger TriggerTypes = "HttpTrigger"
	TriggerTypesUnknown     TriggerTypes = "Unknown"
)

func PossibleValuesForTriggerTypes() []string {
	return []string{
		string(TriggerTypesHTTPTrigger),
		string(TriggerTypesUnknown),
	}
}

func (s *TriggerTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerTypes(input string) (*TriggerTypes, error) {
	vals := map[string]TriggerTypes{
		"httptrigger": TriggerTypesHTTPTrigger,
		"unknown":     TriggerTypesUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerTypes(input)
	return &out, nil
}
