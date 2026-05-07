package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceState string

const (
	ComplianceStateCompliant    ComplianceState = "Compliant"
	ComplianceStateNonCompliant ComplianceState = "NonCompliant"
	ComplianceStateUnknown      ComplianceState = "Unknown"
)

func PossibleValuesForComplianceState() []string {
	return []string{
		string(ComplianceStateCompliant),
		string(ComplianceStateNonCompliant),
		string(ComplianceStateUnknown),
	}
}

func (s *ComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceState(input string) (*ComplianceState, error) {
	vals := map[string]ComplianceState{
		"compliant":    ComplianceStateCompliant,
		"noncompliant": ComplianceStateNonCompliant,
		"unknown":      ComplianceStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceState(input)
	return &out, nil
}

type FieldRestrictionResult string

const (
	FieldRestrictionResultAudit    FieldRestrictionResult = "Audit"
	FieldRestrictionResultDeny     FieldRestrictionResult = "Deny"
	FieldRestrictionResultRemoved  FieldRestrictionResult = "Removed"
	FieldRestrictionResultRequired FieldRestrictionResult = "Required"
)

func PossibleValuesForFieldRestrictionResult() []string {
	return []string{
		string(FieldRestrictionResultAudit),
		string(FieldRestrictionResultDeny),
		string(FieldRestrictionResultRemoved),
		string(FieldRestrictionResultRequired),
	}
}

func (s *FieldRestrictionResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFieldRestrictionResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFieldRestrictionResult(input string) (*FieldRestrictionResult, error) {
	vals := map[string]FieldRestrictionResult{
		"audit":    FieldRestrictionResultAudit,
		"deny":     FieldRestrictionResultDeny,
		"removed":  FieldRestrictionResultRemoved,
		"required": FieldRestrictionResultRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FieldRestrictionResult(input)
	return &out, nil
}

type PolicyStatesResource string

const (
	PolicyStatesResourceDefault PolicyStatesResource = "default"
	PolicyStatesResourceLatest  PolicyStatesResource = "latest"
)

func PossibleValuesForPolicyStatesResource() []string {
	return []string{
		string(PolicyStatesResourceDefault),
		string(PolicyStatesResourceLatest),
	}
}

func (s *PolicyStatesResource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyStatesResource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyStatesResource(input string) (*PolicyStatesResource, error) {
	vals := map[string]PolicyStatesResource{
		"default": PolicyStatesResourceDefault,
		"latest":  PolicyStatesResourceLatest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyStatesResource(input)
	return &out, nil
}

type ResourceDiscoveryMode string

const (
	ResourceDiscoveryModeExistingNonCompliant ResourceDiscoveryMode = "ExistingNonCompliant"
	ResourceDiscoveryModeReEvaluateCompliance ResourceDiscoveryMode = "ReEvaluateCompliance"
)

func PossibleValuesForResourceDiscoveryMode() []string {
	return []string{
		string(ResourceDiscoveryModeExistingNonCompliant),
		string(ResourceDiscoveryModeReEvaluateCompliance),
	}
}

func (s *ResourceDiscoveryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceDiscoveryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceDiscoveryMode(input string) (*ResourceDiscoveryMode, error) {
	vals := map[string]ResourceDiscoveryMode{
		"existingnoncompliant": ResourceDiscoveryModeExistingNonCompliant,
		"reevaluatecompliance": ResourceDiscoveryModeReEvaluateCompliance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceDiscoveryMode(input)
	return &out, nil
}
