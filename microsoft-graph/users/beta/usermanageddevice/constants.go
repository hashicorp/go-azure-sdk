package usermanageddevice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange string

const (
	CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter         CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange = "After"
	CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore        CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange = "Before"
	CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange = "BeforeOrAfter"
)

func PossibleValuesForCreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange() []string {
	return []string{
		string(CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter),
		string(CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore),
		string(CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter),
	}
}

func (s *CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange(input string) (*CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange, error) {
	vals := map[string]CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange{
		"after":         CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter,
		"before":        CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore,
		"beforeorafter": CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange(input)
	return &out, nil
}

type CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState string

const (
	CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState = "BasedOnDeviceCompliancePolicy"
	CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant                  CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState = "NonCompliant"
)

func PossibleValuesForCreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState() []string {
	return []string{
		string(CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy),
		string(CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant),
	}
}

func (s *CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState(input string) (*CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState, error) {
	vals := map[string]CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState{
		"basedondevicecompliancepolicy": CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy,
		"noncompliant":                  CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState(input)
	return &out, nil
}

type CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior string

const (
	CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavioralways                CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior = "Always"
	CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordefault               CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior = "Default"
	CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate       CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior = "DoNotObliterate"
	CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior = "ObliterateWithWarning"
)

func PossibleValuesForCreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior() []string {
	return []string{
		string(CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavioralways),
		string(CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordefault),
		string(CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate),
		string(CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning),
	}
}

func (s *CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior(input string) (*CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior, error) {
	vals := map[string]CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior{
		"always":                CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavioralways,
		"default":               CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordefault,
		"donotobliterate":       CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate,
		"obliteratewithwarning": CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdManagedDeviceByIdWipeRequestObliterationBehavior(input)
	return &out, nil
}

type CreateUserByIdManagedDeviceExecuteActionRequestActionName string

const (
	CreateUserByIdManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim                        CreateUserByIdManagedDeviceExecuteActionRequestActionName = "ActivateDeviceEsim"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamecollectDiagnostics                        CreateUserByIdManagedDeviceExecuteActionRequestActionName = "CollectDiagnostics"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamecustomTextNotification                    CreateUserByIdManagedDeviceExecuteActionRequestActionName = "CustomTextNotification"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamedelete                                    CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Delete"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamedeprovision                               CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Deprovision"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamedisable                                   CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Disable"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamefullScan                                  CreateUserByIdManagedDeviceExecuteActionRequestActionName = "FullScan"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery CreateUserByIdManagedDeviceExecuteActionRequestActionName = "InitiateMobileDeviceManagementKeyRecovery"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation      CreateUserByIdManagedDeviceExecuteActionRequestActionName = "InitiateOnDemandProactiveRemediation"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit            CreateUserByIdManagedDeviceExecuteActionRequestActionName = "MoveDeviceToOrganizationalUnit"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamequickScan                                 CreateUserByIdManagedDeviceExecuteActionRequestActionName = "QuickScan"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamerebootNow                                 CreateUserByIdManagedDeviceExecuteActionRequestActionName = "RebootNow"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamereenable                                  CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Reenable"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNameretire                                    CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Retire"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamesetDeviceName                             CreateUserByIdManagedDeviceExecuteActionRequestActionName = "SetDeviceName"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamesignatureUpdate                           CreateUserByIdManagedDeviceExecuteActionRequestActionName = "SignatureUpdate"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamesyncDevice                                CreateUserByIdManagedDeviceExecuteActionRequestActionName = "SyncDevice"
	CreateUserByIdManagedDeviceExecuteActionRequestActionNamewipe                                      CreateUserByIdManagedDeviceExecuteActionRequestActionName = "Wipe"
)

func PossibleValuesForCreateUserByIdManagedDeviceExecuteActionRequestActionName() []string {
	return []string{
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamecollectDiagnostics),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamecustomTextNotification),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamedelete),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamedeprovision),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamedisable),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamefullScan),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamequickScan),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamerebootNow),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamereenable),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNameretire),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamesetDeviceName),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamesignatureUpdate),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamesyncDevice),
		string(CreateUserByIdManagedDeviceExecuteActionRequestActionNamewipe),
	}
}

func (s *CreateUserByIdManagedDeviceExecuteActionRequestActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdManagedDeviceExecuteActionRequestActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdManagedDeviceExecuteActionRequestActionName(input string) (*CreateUserByIdManagedDeviceExecuteActionRequestActionName, error) {
	vals := map[string]CreateUserByIdManagedDeviceExecuteActionRequestActionName{
		"activatedeviceesim":     CreateUserByIdManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim,
		"collectdiagnostics":     CreateUserByIdManagedDeviceExecuteActionRequestActionNamecollectDiagnostics,
		"customtextnotification": CreateUserByIdManagedDeviceExecuteActionRequestActionNamecustomTextNotification,
		"delete":                 CreateUserByIdManagedDeviceExecuteActionRequestActionNamedelete,
		"deprovision":            CreateUserByIdManagedDeviceExecuteActionRequestActionNamedeprovision,
		"disable":                CreateUserByIdManagedDeviceExecuteActionRequestActionNamedisable,
		"fullscan":               CreateUserByIdManagedDeviceExecuteActionRequestActionNamefullScan,
		"initiatemobiledevicemanagementkeyrecovery": CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      CreateUserByIdManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation,
		"movedevicetoorganizationalunit":            CreateUserByIdManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit,
		"quickscan":                                 CreateUserByIdManagedDeviceExecuteActionRequestActionNamequickScan,
		"rebootnow":                                 CreateUserByIdManagedDeviceExecuteActionRequestActionNamerebootNow,
		"reenable":                                  CreateUserByIdManagedDeviceExecuteActionRequestActionNamereenable,
		"retire":                                    CreateUserByIdManagedDeviceExecuteActionRequestActionNameretire,
		"setdevicename":                             CreateUserByIdManagedDeviceExecuteActionRequestActionNamesetDeviceName,
		"signatureupdate":                           CreateUserByIdManagedDeviceExecuteActionRequestActionNamesignatureUpdate,
		"syncdevice":                                CreateUserByIdManagedDeviceExecuteActionRequestActionNamesyncDevice,
		"wipe":                                      CreateUserByIdManagedDeviceExecuteActionRequestActionNamewipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdManagedDeviceExecuteActionRequestActionName(input)
	return &out, nil
}
