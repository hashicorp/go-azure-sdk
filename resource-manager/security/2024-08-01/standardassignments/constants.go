package standardassignments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationComplianceState string

const (
	AttestationComplianceStateCompliant    AttestationComplianceState = "compliant"
	AttestationComplianceStateNonCompliant AttestationComplianceState = "nonCompliant"
	AttestationComplianceStateUnknown      AttestationComplianceState = "unknown"
)

func PossibleValuesForAttestationComplianceState() []string {
	return []string{
		string(AttestationComplianceStateCompliant),
		string(AttestationComplianceStateNonCompliant),
		string(AttestationComplianceStateUnknown),
	}
}

func (s *AttestationComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttestationComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttestationComplianceState(input string) (*AttestationComplianceState, error) {
	vals := map[string]AttestationComplianceState{
		"compliant":    AttestationComplianceStateCompliant,
		"noncompliant": AttestationComplianceStateNonCompliant,
		"unknown":      AttestationComplianceStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttestationComplianceState(input)
	return &out, nil
}

type Effect string

const (
	EffectAttest Effect = "Attest"
	EffectAudit  Effect = "Audit"
	EffectExempt Effect = "Exempt"
)

func PossibleValuesForEffect() []string {
	return []string{
		string(EffectAttest),
		string(EffectAudit),
		string(EffectExempt),
	}
}

func (s *Effect) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEffect(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEffect(input string) (*Effect, error) {
	vals := map[string]Effect{
		"attest": EffectAttest,
		"audit":  EffectAudit,
		"exempt": EffectExempt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Effect(input)
	return &out, nil
}

type ExemptionCategory string

const (
	ExemptionCategoryMitigated ExemptionCategory = "mitigated"
	ExemptionCategoryWaiver    ExemptionCategory = "waiver"
)

func PossibleValuesForExemptionCategory() []string {
	return []string{
		string(ExemptionCategoryMitigated),
		string(ExemptionCategoryWaiver),
	}
}

func (s *ExemptionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExemptionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExemptionCategory(input string) (*ExemptionCategory, error) {
	vals := map[string]ExemptionCategory{
		"mitigated": ExemptionCategoryMitigated,
		"waiver":    ExemptionCategoryWaiver,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExemptionCategory(input)
	return &out, nil
}
