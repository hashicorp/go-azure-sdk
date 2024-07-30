package grouppolicyconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone),
	}
}

func (s *DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}

type GroupPolicyCategoryIngestionSource string

const (
	GroupPolicyCategoryIngestionSourceBuiltIn GroupPolicyCategoryIngestionSource = "builtIn"
	GroupPolicyCategoryIngestionSourceCustom  GroupPolicyCategoryIngestionSource = "custom"
	GroupPolicyCategoryIngestionSourceUnknown GroupPolicyCategoryIngestionSource = "unknown"
)

func PossibleValuesForGroupPolicyCategoryIngestionSource() []string {
	return []string{
		string(GroupPolicyCategoryIngestionSourceBuiltIn),
		string(GroupPolicyCategoryIngestionSourceCustom),
		string(GroupPolicyCategoryIngestionSourceUnknown),
	}
}

func (s *GroupPolicyCategoryIngestionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyCategoryIngestionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyCategoryIngestionSource(input string) (*GroupPolicyCategoryIngestionSource, error) {
	vals := map[string]GroupPolicyCategoryIngestionSource{
		"builtin": GroupPolicyCategoryIngestionSourceBuiltIn,
		"custom":  GroupPolicyCategoryIngestionSourceCustom,
		"unknown": GroupPolicyCategoryIngestionSourceUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyCategoryIngestionSource(input)
	return &out, nil
}

type GroupPolicyDefinitionClassType string

const (
	GroupPolicyDefinitionClassTypeMachine GroupPolicyDefinitionClassType = "machine"
	GroupPolicyDefinitionClassTypeUser    GroupPolicyDefinitionClassType = "user"
)

func PossibleValuesForGroupPolicyDefinitionClassType() []string {
	return []string{
		string(GroupPolicyDefinitionClassTypeMachine),
		string(GroupPolicyDefinitionClassTypeUser),
	}
}

func (s *GroupPolicyDefinitionClassType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionClassType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionClassType(input string) (*GroupPolicyDefinitionClassType, error) {
	vals := map[string]GroupPolicyDefinitionClassType{
		"machine": GroupPolicyDefinitionClassTypeMachine,
		"user":    GroupPolicyDefinitionClassTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionClassType(input)
	return &out, nil
}

type GroupPolicyDefinitionFilePolicyType string

const (
	GroupPolicyDefinitionFilePolicyTypeAdmxBacked   GroupPolicyDefinitionFilePolicyType = "admxBacked"
	GroupPolicyDefinitionFilePolicyTypeAdmxIngested GroupPolicyDefinitionFilePolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyDefinitionFilePolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionFilePolicyTypeAdmxBacked),
		string(GroupPolicyDefinitionFilePolicyTypeAdmxIngested),
	}
}

func (s *GroupPolicyDefinitionFilePolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionFilePolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionFilePolicyType(input string) (*GroupPolicyDefinitionFilePolicyType, error) {
	vals := map[string]GroupPolicyDefinitionFilePolicyType{
		"admxbacked":   GroupPolicyDefinitionFilePolicyTypeAdmxBacked,
		"admxingested": GroupPolicyDefinitionFilePolicyTypeAdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionFilePolicyType(input)
	return &out, nil
}

type GroupPolicyDefinitionPolicyType string

const (
	GroupPolicyDefinitionPolicyTypeAdmxBacked   GroupPolicyDefinitionPolicyType = "admxBacked"
	GroupPolicyDefinitionPolicyTypeAdmxIngested GroupPolicyDefinitionPolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyDefinitionPolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionPolicyTypeAdmxBacked),
		string(GroupPolicyDefinitionPolicyTypeAdmxIngested),
	}
}

func (s *GroupPolicyDefinitionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionPolicyType(input string) (*GroupPolicyDefinitionPolicyType, error) {
	vals := map[string]GroupPolicyDefinitionPolicyType{
		"admxbacked":   GroupPolicyDefinitionPolicyTypeAdmxBacked,
		"admxingested": GroupPolicyDefinitionPolicyTypeAdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionPolicyType(input)
	return &out, nil
}

type GroupPolicyDefinitionValueConfigurationType string

const (
	GroupPolicyDefinitionValueConfigurationTypePolicy     GroupPolicyDefinitionValueConfigurationType = "policy"
	GroupPolicyDefinitionValueConfigurationTypePreference GroupPolicyDefinitionValueConfigurationType = "preference"
)

func PossibleValuesForGroupPolicyDefinitionValueConfigurationType() []string {
	return []string{
		string(GroupPolicyDefinitionValueConfigurationTypePolicy),
		string(GroupPolicyDefinitionValueConfigurationTypePreference),
	}
}

func (s *GroupPolicyDefinitionValueConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionValueConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionValueConfigurationType(input string) (*GroupPolicyDefinitionValueConfigurationType, error) {
	vals := map[string]GroupPolicyDefinitionValueConfigurationType{
		"policy":     GroupPolicyDefinitionValueConfigurationTypePolicy,
		"preference": GroupPolicyDefinitionValueConfigurationTypePreference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionValueConfigurationType(input)
	return &out, nil
}
