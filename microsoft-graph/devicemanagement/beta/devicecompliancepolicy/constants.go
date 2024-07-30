package devicecompliancepolicy

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

type DeviceComplianceActionItemActionType string

const (
	DeviceComplianceActionItemActionTypeBlock                        DeviceComplianceActionItemActionType = "block"
	DeviceComplianceActionItemActionTypeNoAction                     DeviceComplianceActionItemActionType = "noAction"
	DeviceComplianceActionItemActionTypeNotification                 DeviceComplianceActionItemActionType = "notification"
	DeviceComplianceActionItemActionTypePushNotification             DeviceComplianceActionItemActionType = "pushNotification"
	DeviceComplianceActionItemActionTypeRemoteLock                   DeviceComplianceActionItemActionType = "remoteLock"
	DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles DeviceComplianceActionItemActionType = "removeResourceAccessProfiles"
	DeviceComplianceActionItemActionTypeRetire                       DeviceComplianceActionItemActionType = "retire"
	DeviceComplianceActionItemActionTypeWipe                         DeviceComplianceActionItemActionType = "wipe"
)

func PossibleValuesForDeviceComplianceActionItemActionType() []string {
	return []string{
		string(DeviceComplianceActionItemActionTypeBlock),
		string(DeviceComplianceActionItemActionTypeNoAction),
		string(DeviceComplianceActionItemActionTypeNotification),
		string(DeviceComplianceActionItemActionTypePushNotification),
		string(DeviceComplianceActionItemActionTypeRemoteLock),
		string(DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles),
		string(DeviceComplianceActionItemActionTypeRetire),
		string(DeviceComplianceActionItemActionTypeWipe),
	}
}

func (s *DeviceComplianceActionItemActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceActionItemActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceActionItemActionType(input string) (*DeviceComplianceActionItemActionType, error) {
	vals := map[string]DeviceComplianceActionItemActionType{
		"block":                        DeviceComplianceActionItemActionTypeBlock,
		"noaction":                     DeviceComplianceActionItemActionTypeNoAction,
		"notification":                 DeviceComplianceActionItemActionTypeNotification,
		"pushnotification":             DeviceComplianceActionItemActionTypePushNotification,
		"remotelock":                   DeviceComplianceActionItemActionTypeRemoteLock,
		"removeresourceaccessprofiles": DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles,
		"retire":                       DeviceComplianceActionItemActionTypeRetire,
		"wipe":                         DeviceComplianceActionItemActionTypeWipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceActionItemActionType(input)
	return &out, nil
}

type DeviceCompliancePolicyAssignmentSource string

const (
	DeviceCompliancePolicyAssignmentSourceDirect     DeviceCompliancePolicyAssignmentSource = "direct"
	DeviceCompliancePolicyAssignmentSourcePolicySets DeviceCompliancePolicyAssignmentSource = "policySets"
)

func PossibleValuesForDeviceCompliancePolicyAssignmentSource() []string {
	return []string{
		string(DeviceCompliancePolicyAssignmentSourceDirect),
		string(DeviceCompliancePolicyAssignmentSourcePolicySets),
	}
}

func (s *DeviceCompliancePolicyAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyAssignmentSource(input string) (*DeviceCompliancePolicyAssignmentSource, error) {
	vals := map[string]DeviceCompliancePolicyAssignmentSource{
		"direct":     DeviceCompliancePolicyAssignmentSourceDirect,
		"policysets": DeviceCompliancePolicyAssignmentSourcePolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyAssignmentSource(input)
	return &out, nil
}

type SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState string

const (
	SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateCancelRetire  SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState = "cancelRetire"
	SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateConfirmRetire SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState = "confirmRetire"
)

func PossibleValuesForSetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState() []string {
	return []string{
		string(SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateCancelRetire),
		string(SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateConfirmRetire),
	}
}

func (s *SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState(input string) (*SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState, error) {
	vals := map[string]SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState{
		"cancelretire":  SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateCancelRetire,
		"confirmretire": SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestStateConfirmRetire,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState(input)
	return &out, nil
}
