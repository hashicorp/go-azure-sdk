package sourcecontrolconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceStateType string

const (
	ComplianceStateTypeCompliant    ComplianceStateType = "Compliant"
	ComplianceStateTypeFailed       ComplianceStateType = "Failed"
	ComplianceStateTypeInstalled    ComplianceStateType = "Installed"
	ComplianceStateTypeNoncompliant ComplianceStateType = "Noncompliant"
	ComplianceStateTypePending      ComplianceStateType = "Pending"
)

func PossibleValuesForComplianceStateType() []string {
	return []string{
		string(ComplianceStateTypeCompliant),
		string(ComplianceStateTypeFailed),
		string(ComplianceStateTypeInstalled),
		string(ComplianceStateTypeNoncompliant),
		string(ComplianceStateTypePending),
	}
}

func (s *ComplianceStateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceStateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceStateType(input string) (*ComplianceStateType, error) {
	vals := map[string]ComplianceStateType{
		"compliant":    ComplianceStateTypeCompliant,
		"failed":       ComplianceStateTypeFailed,
		"installed":    ComplianceStateTypeInstalled,
		"noncompliant": ComplianceStateTypeNoncompliant,
		"pending":      ComplianceStateTypePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceStateType(input)
	return &out, nil
}

type MessageLevelType string

const (
	MessageLevelTypeError       MessageLevelType = "Error"
	MessageLevelTypeInformation MessageLevelType = "Information"
	MessageLevelTypeWarning     MessageLevelType = "Warning"
)

func PossibleValuesForMessageLevelType() []string {
	return []string{
		string(MessageLevelTypeError),
		string(MessageLevelTypeInformation),
		string(MessageLevelTypeWarning),
	}
}

func (s *MessageLevelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageLevelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageLevelType(input string) (*MessageLevelType, error) {
	vals := map[string]MessageLevelType{
		"error":       MessageLevelTypeError,
		"information": MessageLevelTypeInformation,
		"warning":     MessageLevelTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageLevelType(input)
	return &out, nil
}

type OperatorScopeType string

const (
	OperatorScopeTypeCluster   OperatorScopeType = "cluster"
	OperatorScopeTypeNamespace OperatorScopeType = "namespace"
)

func PossibleValuesForOperatorScopeType() []string {
	return []string{
		string(OperatorScopeTypeCluster),
		string(OperatorScopeTypeNamespace),
	}
}

func (s *OperatorScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperatorScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperatorScopeType(input string) (*OperatorScopeType, error) {
	vals := map[string]OperatorScopeType{
		"cluster":   OperatorScopeTypeCluster,
		"namespace": OperatorScopeTypeNamespace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatorScopeType(input)
	return &out, nil
}

type OperatorType string

const (
	OperatorTypeFlux OperatorType = "Flux"
)

func PossibleValuesForOperatorType() []string {
	return []string{
		string(OperatorTypeFlux),
	}
}

func (s *OperatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperatorType(input string) (*OperatorType, error) {
	vals := map[string]OperatorType{
		"flux": OperatorTypeFlux,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatorType(input)
	return &out, nil
}

type ProvisioningStateType string

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = "Accepted"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeRunning   ProvisioningStateType = "Running"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
)

func PossibleValuesForProvisioningStateType() []string {
	return []string{
		string(ProvisioningStateTypeAccepted),
		string(ProvisioningStateTypeDeleting),
		string(ProvisioningStateTypeFailed),
		string(ProvisioningStateTypeRunning),
		string(ProvisioningStateTypeSucceeded),
	}
}

func (s *ProvisioningStateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStateType(input string) (*ProvisioningStateType, error) {
	vals := map[string]ProvisioningStateType{
		"accepted":  ProvisioningStateTypeAccepted,
		"deleting":  ProvisioningStateTypeDeleting,
		"failed":    ProvisioningStateTypeFailed,
		"running":   ProvisioningStateTypeRunning,
		"succeeded": ProvisioningStateTypeSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStateType(input)
	return &out, nil
}
