package networksecurityperimeter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRuleDirection string

const (
	AccessRuleDirectionInbound  AccessRuleDirection = "Inbound"
	AccessRuleDirectionOutbound AccessRuleDirection = "Outbound"
)

func PossibleValuesForAccessRuleDirection() []string {
	return []string{
		string(AccessRuleDirectionInbound),
		string(AccessRuleDirectionOutbound),
	}
}

func (s *AccessRuleDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessRuleDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessRuleDirection(input string) (*AccessRuleDirection, error) {
	vals := map[string]AccessRuleDirection{
		"inbound":  AccessRuleDirectionInbound,
		"outbound": AccessRuleDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessRuleDirection(input)
	return &out, nil
}

type IssueType string

const (
	IssueTypeConfigurationPropagationFailure IssueType = "ConfigurationPropagationFailure"
	IssueTypeMissingIdentityConfiguration    IssueType = "MissingIdentityConfiguration"
	IssueTypeMissingPerimeterConfiguration   IssueType = "MissingPerimeterConfiguration"
	IssueTypeUnknown                         IssueType = "Unknown"
)

func PossibleValuesForIssueType() []string {
	return []string{
		string(IssueTypeConfigurationPropagationFailure),
		string(IssueTypeMissingIdentityConfiguration),
		string(IssueTypeMissingPerimeterConfiguration),
		string(IssueTypeUnknown),
	}
}

func (s *IssueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIssueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIssueType(input string) (*IssueType, error) {
	vals := map[string]IssueType{
		"configurationpropagationfailure": IssueTypeConfigurationPropagationFailure,
		"missingidentityconfiguration":    IssueTypeMissingIdentityConfiguration,
		"missingperimeterconfiguration":   IssueTypeMissingPerimeterConfiguration,
		"unknown":                         IssueTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IssueType(input)
	return &out, nil
}

type NetworkSecurityPerimeterConfigurationProvisioningState string

const (
	NetworkSecurityPerimeterConfigurationProvisioningStateAccepted  NetworkSecurityPerimeterConfigurationProvisioningState = "Accepted"
	NetworkSecurityPerimeterConfigurationProvisioningStateCanceled  NetworkSecurityPerimeterConfigurationProvisioningState = "Canceled"
	NetworkSecurityPerimeterConfigurationProvisioningStateCreating  NetworkSecurityPerimeterConfigurationProvisioningState = "Creating"
	NetworkSecurityPerimeterConfigurationProvisioningStateDeleting  NetworkSecurityPerimeterConfigurationProvisioningState = "Deleting"
	NetworkSecurityPerimeterConfigurationProvisioningStateFailed    NetworkSecurityPerimeterConfigurationProvisioningState = "Failed"
	NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded NetworkSecurityPerimeterConfigurationProvisioningState = "Succeeded"
	NetworkSecurityPerimeterConfigurationProvisioningStateUpdating  NetworkSecurityPerimeterConfigurationProvisioningState = "Updating"
)

func PossibleValuesForNetworkSecurityPerimeterConfigurationProvisioningState() []string {
	return []string{
		string(NetworkSecurityPerimeterConfigurationProvisioningStateAccepted),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateCanceled),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateCreating),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateDeleting),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateFailed),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateUpdating),
	}
}

func (s *NetworkSecurityPerimeterConfigurationProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterConfigurationProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterConfigurationProvisioningState(input string) (*NetworkSecurityPerimeterConfigurationProvisioningState, error) {
	vals := map[string]NetworkSecurityPerimeterConfigurationProvisioningState{
		"accepted":  NetworkSecurityPerimeterConfigurationProvisioningStateAccepted,
		"canceled":  NetworkSecurityPerimeterConfigurationProvisioningStateCanceled,
		"creating":  NetworkSecurityPerimeterConfigurationProvisioningStateCreating,
		"deleting":  NetworkSecurityPerimeterConfigurationProvisioningStateDeleting,
		"failed":    NetworkSecurityPerimeterConfigurationProvisioningStateFailed,
		"succeeded": NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded,
		"updating":  NetworkSecurityPerimeterConfigurationProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterConfigurationProvisioningState(input)
	return &out, nil
}

type ResourceAssociationAccessMode string

const (
	ResourceAssociationAccessModeAudit    ResourceAssociationAccessMode = "Audit"
	ResourceAssociationAccessModeEnforced ResourceAssociationAccessMode = "Enforced"
	ResourceAssociationAccessModeLearning ResourceAssociationAccessMode = "Learning"
)

func PossibleValuesForResourceAssociationAccessMode() []string {
	return []string{
		string(ResourceAssociationAccessModeAudit),
		string(ResourceAssociationAccessModeEnforced),
		string(ResourceAssociationAccessModeLearning),
	}
}

func (s *ResourceAssociationAccessMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceAssociationAccessMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceAssociationAccessMode(input string) (*ResourceAssociationAccessMode, error) {
	vals := map[string]ResourceAssociationAccessMode{
		"audit":    ResourceAssociationAccessModeAudit,
		"enforced": ResourceAssociationAccessModeEnforced,
		"learning": ResourceAssociationAccessModeLearning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceAssociationAccessMode(input)
	return &out, nil
}

type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityWarning Severity = "Warning"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityError),
		string(SeverityWarning),
	}
}

func (s *Severity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeverity(input string) (*Severity, error) {
	vals := map[string]Severity{
		"error":   SeverityError,
		"warning": SeverityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Severity(input)
	return &out, nil
}
