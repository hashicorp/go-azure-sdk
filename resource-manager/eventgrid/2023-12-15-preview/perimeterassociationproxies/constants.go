package perimeterassociationproxies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterAssociationAccessMode string

const (
	NetworkSecurityPerimeterAssociationAccessModeAudit    NetworkSecurityPerimeterAssociationAccessMode = "Audit"
	NetworkSecurityPerimeterAssociationAccessModeEnforced NetworkSecurityPerimeterAssociationAccessMode = "Enforced"
	NetworkSecurityPerimeterAssociationAccessModeLearning NetworkSecurityPerimeterAssociationAccessMode = "Learning"
)

func PossibleValuesForNetworkSecurityPerimeterAssociationAccessMode() []string {
	return []string{
		string(NetworkSecurityPerimeterAssociationAccessModeAudit),
		string(NetworkSecurityPerimeterAssociationAccessModeEnforced),
		string(NetworkSecurityPerimeterAssociationAccessModeLearning),
	}
}

func (s *NetworkSecurityPerimeterAssociationAccessMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterAssociationAccessMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterAssociationAccessMode(input string) (*NetworkSecurityPerimeterAssociationAccessMode, error) {
	vals := map[string]NetworkSecurityPerimeterAssociationAccessMode{
		"audit":    NetworkSecurityPerimeterAssociationAccessModeAudit,
		"enforced": NetworkSecurityPerimeterAssociationAccessModeEnforced,
		"learning": NetworkSecurityPerimeterAssociationAccessModeLearning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterAssociationAccessMode(input)
	return &out, nil
}

type NetworkSecurityPerimeterConfigProvisioningState string

const (
	NetworkSecurityPerimeterConfigProvisioningStateAccepted  NetworkSecurityPerimeterConfigProvisioningState = "Accepted"
	NetworkSecurityPerimeterConfigProvisioningStateCanceled  NetworkSecurityPerimeterConfigProvisioningState = "Canceled"
	NetworkSecurityPerimeterConfigProvisioningStateCreating  NetworkSecurityPerimeterConfigProvisioningState = "Creating"
	NetworkSecurityPerimeterConfigProvisioningStateDeleted   NetworkSecurityPerimeterConfigProvisioningState = "Deleted"
	NetworkSecurityPerimeterConfigProvisioningStateDeleting  NetworkSecurityPerimeterConfigProvisioningState = "Deleting"
	NetworkSecurityPerimeterConfigProvisioningStateFailed    NetworkSecurityPerimeterConfigProvisioningState = "Failed"
	NetworkSecurityPerimeterConfigProvisioningStateSucceeded NetworkSecurityPerimeterConfigProvisioningState = "Succeeded"
	NetworkSecurityPerimeterConfigProvisioningStateUpdating  NetworkSecurityPerimeterConfigProvisioningState = "Updating"
)

func PossibleValuesForNetworkSecurityPerimeterConfigProvisioningState() []string {
	return []string{
		string(NetworkSecurityPerimeterConfigProvisioningStateAccepted),
		string(NetworkSecurityPerimeterConfigProvisioningStateCanceled),
		string(NetworkSecurityPerimeterConfigProvisioningStateCreating),
		string(NetworkSecurityPerimeterConfigProvisioningStateDeleted),
		string(NetworkSecurityPerimeterConfigProvisioningStateDeleting),
		string(NetworkSecurityPerimeterConfigProvisioningStateFailed),
		string(NetworkSecurityPerimeterConfigProvisioningStateSucceeded),
		string(NetworkSecurityPerimeterConfigProvisioningStateUpdating),
	}
}

func (s *NetworkSecurityPerimeterConfigProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterConfigProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterConfigProvisioningState(input string) (*NetworkSecurityPerimeterConfigProvisioningState, error) {
	vals := map[string]NetworkSecurityPerimeterConfigProvisioningState{
		"accepted":  NetworkSecurityPerimeterConfigProvisioningStateAccepted,
		"canceled":  NetworkSecurityPerimeterConfigProvisioningStateCanceled,
		"creating":  NetworkSecurityPerimeterConfigProvisioningStateCreating,
		"deleted":   NetworkSecurityPerimeterConfigProvisioningStateDeleted,
		"deleting":  NetworkSecurityPerimeterConfigProvisioningStateDeleting,
		"failed":    NetworkSecurityPerimeterConfigProvisioningStateFailed,
		"succeeded": NetworkSecurityPerimeterConfigProvisioningStateSucceeded,
		"updating":  NetworkSecurityPerimeterConfigProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterConfigProvisioningState(input)
	return &out, nil
}

type NetworkSecurityPerimeterConfigurationIssueSeverity string

const (
	NetworkSecurityPerimeterConfigurationIssueSeverityError   NetworkSecurityPerimeterConfigurationIssueSeverity = "Error"
	NetworkSecurityPerimeterConfigurationIssueSeverityWarning NetworkSecurityPerimeterConfigurationIssueSeverity = "Warning"
)

func PossibleValuesForNetworkSecurityPerimeterConfigurationIssueSeverity() []string {
	return []string{
		string(NetworkSecurityPerimeterConfigurationIssueSeverityError),
		string(NetworkSecurityPerimeterConfigurationIssueSeverityWarning),
	}
}

func (s *NetworkSecurityPerimeterConfigurationIssueSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterConfigurationIssueSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterConfigurationIssueSeverity(input string) (*NetworkSecurityPerimeterConfigurationIssueSeverity, error) {
	vals := map[string]NetworkSecurityPerimeterConfigurationIssueSeverity{
		"error":   NetworkSecurityPerimeterConfigurationIssueSeverityError,
		"warning": NetworkSecurityPerimeterConfigurationIssueSeverityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterConfigurationIssueSeverity(input)
	return &out, nil
}

type NetworkSecurityPerimeterConfigurationIssueType string

const (
	NetworkSecurityPerimeterConfigurationIssueTypeConfigurationPropagationFailure NetworkSecurityPerimeterConfigurationIssueType = "ConfigurationPropagationFailure"
	NetworkSecurityPerimeterConfigurationIssueTypeMissingIdentityConfiguration    NetworkSecurityPerimeterConfigurationIssueType = "MissingIdentityConfiguration"
	NetworkSecurityPerimeterConfigurationIssueTypeMissingPerimeterConfiguration   NetworkSecurityPerimeterConfigurationIssueType = "MissingPerimeterConfiguration"
	NetworkSecurityPerimeterConfigurationIssueTypeOther                           NetworkSecurityPerimeterConfigurationIssueType = "Other"
)

func PossibleValuesForNetworkSecurityPerimeterConfigurationIssueType() []string {
	return []string{
		string(NetworkSecurityPerimeterConfigurationIssueTypeConfigurationPropagationFailure),
		string(NetworkSecurityPerimeterConfigurationIssueTypeMissingIdentityConfiguration),
		string(NetworkSecurityPerimeterConfigurationIssueTypeMissingPerimeterConfiguration),
		string(NetworkSecurityPerimeterConfigurationIssueTypeOther),
	}
}

func (s *NetworkSecurityPerimeterConfigurationIssueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterConfigurationIssueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterConfigurationIssueType(input string) (*NetworkSecurityPerimeterConfigurationIssueType, error) {
	vals := map[string]NetworkSecurityPerimeterConfigurationIssueType{
		"configurationpropagationfailure": NetworkSecurityPerimeterConfigurationIssueTypeConfigurationPropagationFailure,
		"missingidentityconfiguration":    NetworkSecurityPerimeterConfigurationIssueTypeMissingIdentityConfiguration,
		"missingperimeterconfiguration":   NetworkSecurityPerimeterConfigurationIssueTypeMissingPerimeterConfiguration,
		"other":                           NetworkSecurityPerimeterConfigurationIssueTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterConfigurationIssueType(input)
	return &out, nil
}

type NetworkSecurityPerimeterProfileAccessRuleDirection string

const (
	NetworkSecurityPerimeterProfileAccessRuleDirectionInbound  NetworkSecurityPerimeterProfileAccessRuleDirection = "Inbound"
	NetworkSecurityPerimeterProfileAccessRuleDirectionOutbound NetworkSecurityPerimeterProfileAccessRuleDirection = "Outbound"
)

func PossibleValuesForNetworkSecurityPerimeterProfileAccessRuleDirection() []string {
	return []string{
		string(NetworkSecurityPerimeterProfileAccessRuleDirectionInbound),
		string(NetworkSecurityPerimeterProfileAccessRuleDirectionOutbound),
	}
}

func (s *NetworkSecurityPerimeterProfileAccessRuleDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityPerimeterProfileAccessRuleDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityPerimeterProfileAccessRuleDirection(input string) (*NetworkSecurityPerimeterProfileAccessRuleDirection, error) {
	vals := map[string]NetworkSecurityPerimeterProfileAccessRuleDirection{
		"inbound":  NetworkSecurityPerimeterProfileAccessRuleDirectionInbound,
		"outbound": NetworkSecurityPerimeterProfileAccessRuleDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterProfileAccessRuleDirection(input)
	return &out, nil
}
