package memanageddevice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange string

const (
	CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter         CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange = "After"
	CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore        CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange = "Before"
	CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange = "BeforeOrAfter"
)

func PossibleValuesForCreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange() []string {
	return []string{
		string(CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter),
		string(CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore),
		string(CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter),
	}
}

func (s *CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange(input string) (*CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange, error) {
	vals := map[string]CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange{
		"after":         CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangeafter,
		"before":        CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebefore,
		"beforeorafter": CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRangebeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange(input)
	return &out, nil
}

type CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState string

const (
	CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState = "BasedOnDeviceCompliancePolicy"
	CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant                  CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState = "NonCompliant"
)

func PossibleValuesForCreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState() []string {
	return []string{
		string(CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy),
		string(CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant),
	}
}

func (s *CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState(input string) (*CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState, error) {
	vals := map[string]CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState{
		"basedondevicecompliancepolicy": CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatebasedOnDeviceCompliancePolicy,
		"noncompliant":                  CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceStatenonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeManagedDeviceByIdOverrideComplianceStateRequestComplianceState(input)
	return &out, nil
}

type CreateMeManagedDeviceByIdWipeRequestObliterationBehavior string

const (
	CreateMeManagedDeviceByIdWipeRequestObliterationBehavioralways                CreateMeManagedDeviceByIdWipeRequestObliterationBehavior = "Always"
	CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordefault               CreateMeManagedDeviceByIdWipeRequestObliterationBehavior = "Default"
	CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate       CreateMeManagedDeviceByIdWipeRequestObliterationBehavior = "DoNotObliterate"
	CreateMeManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning CreateMeManagedDeviceByIdWipeRequestObliterationBehavior = "ObliterateWithWarning"
)

func PossibleValuesForCreateMeManagedDeviceByIdWipeRequestObliterationBehavior() []string {
	return []string{
		string(CreateMeManagedDeviceByIdWipeRequestObliterationBehavioralways),
		string(CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordefault),
		string(CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate),
		string(CreateMeManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning),
	}
}

func (s *CreateMeManagedDeviceByIdWipeRequestObliterationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeManagedDeviceByIdWipeRequestObliterationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeManagedDeviceByIdWipeRequestObliterationBehavior(input string) (*CreateMeManagedDeviceByIdWipeRequestObliterationBehavior, error) {
	vals := map[string]CreateMeManagedDeviceByIdWipeRequestObliterationBehavior{
		"always":                CreateMeManagedDeviceByIdWipeRequestObliterationBehavioralways,
		"default":               CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordefault,
		"donotobliterate":       CreateMeManagedDeviceByIdWipeRequestObliterationBehaviordoNotObliterate,
		"obliteratewithwarning": CreateMeManagedDeviceByIdWipeRequestObliterationBehaviorobliterateWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeManagedDeviceByIdWipeRequestObliterationBehavior(input)
	return &out, nil
}

type CreateMeManagedDeviceExecuteActionRequestActionName string

const (
	CreateMeManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim                        CreateMeManagedDeviceExecuteActionRequestActionName = "ActivateDeviceEsim"
	CreateMeManagedDeviceExecuteActionRequestActionNamecollectDiagnostics                        CreateMeManagedDeviceExecuteActionRequestActionName = "CollectDiagnostics"
	CreateMeManagedDeviceExecuteActionRequestActionNamecustomTextNotification                    CreateMeManagedDeviceExecuteActionRequestActionName = "CustomTextNotification"
	CreateMeManagedDeviceExecuteActionRequestActionNamedelete                                    CreateMeManagedDeviceExecuteActionRequestActionName = "Delete"
	CreateMeManagedDeviceExecuteActionRequestActionNamedeprovision                               CreateMeManagedDeviceExecuteActionRequestActionName = "Deprovision"
	CreateMeManagedDeviceExecuteActionRequestActionNamedisable                                   CreateMeManagedDeviceExecuteActionRequestActionName = "Disable"
	CreateMeManagedDeviceExecuteActionRequestActionNamefullScan                                  CreateMeManagedDeviceExecuteActionRequestActionName = "FullScan"
	CreateMeManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery CreateMeManagedDeviceExecuteActionRequestActionName = "InitiateMobileDeviceManagementKeyRecovery"
	CreateMeManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation      CreateMeManagedDeviceExecuteActionRequestActionName = "InitiateOnDemandProactiveRemediation"
	CreateMeManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit            CreateMeManagedDeviceExecuteActionRequestActionName = "MoveDeviceToOrganizationalUnit"
	CreateMeManagedDeviceExecuteActionRequestActionNamequickScan                                 CreateMeManagedDeviceExecuteActionRequestActionName = "QuickScan"
	CreateMeManagedDeviceExecuteActionRequestActionNamerebootNow                                 CreateMeManagedDeviceExecuteActionRequestActionName = "RebootNow"
	CreateMeManagedDeviceExecuteActionRequestActionNamereenable                                  CreateMeManagedDeviceExecuteActionRequestActionName = "Reenable"
	CreateMeManagedDeviceExecuteActionRequestActionNameretire                                    CreateMeManagedDeviceExecuteActionRequestActionName = "Retire"
	CreateMeManagedDeviceExecuteActionRequestActionNamesetDeviceName                             CreateMeManagedDeviceExecuteActionRequestActionName = "SetDeviceName"
	CreateMeManagedDeviceExecuteActionRequestActionNamesignatureUpdate                           CreateMeManagedDeviceExecuteActionRequestActionName = "SignatureUpdate"
	CreateMeManagedDeviceExecuteActionRequestActionNamesyncDevice                                CreateMeManagedDeviceExecuteActionRequestActionName = "SyncDevice"
	CreateMeManagedDeviceExecuteActionRequestActionNamewipe                                      CreateMeManagedDeviceExecuteActionRequestActionName = "Wipe"
)

func PossibleValuesForCreateMeManagedDeviceExecuteActionRequestActionName() []string {
	return []string{
		string(CreateMeManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamecollectDiagnostics),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamecustomTextNotification),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamedelete),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamedeprovision),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamedisable),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamefullScan),
		string(CreateMeManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery),
		string(CreateMeManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamequickScan),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamerebootNow),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamereenable),
		string(CreateMeManagedDeviceExecuteActionRequestActionNameretire),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamesetDeviceName),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamesignatureUpdate),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamesyncDevice),
		string(CreateMeManagedDeviceExecuteActionRequestActionNamewipe),
	}
}

func (s *CreateMeManagedDeviceExecuteActionRequestActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeManagedDeviceExecuteActionRequestActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeManagedDeviceExecuteActionRequestActionName(input string) (*CreateMeManagedDeviceExecuteActionRequestActionName, error) {
	vals := map[string]CreateMeManagedDeviceExecuteActionRequestActionName{
		"activatedeviceesim":     CreateMeManagedDeviceExecuteActionRequestActionNameactivateDeviceEsim,
		"collectdiagnostics":     CreateMeManagedDeviceExecuteActionRequestActionNamecollectDiagnostics,
		"customtextnotification": CreateMeManagedDeviceExecuteActionRequestActionNamecustomTextNotification,
		"delete":                 CreateMeManagedDeviceExecuteActionRequestActionNamedelete,
		"deprovision":            CreateMeManagedDeviceExecuteActionRequestActionNamedeprovision,
		"disable":                CreateMeManagedDeviceExecuteActionRequestActionNamedisable,
		"fullscan":               CreateMeManagedDeviceExecuteActionRequestActionNamefullScan,
		"initiatemobiledevicemanagementkeyrecovery": CreateMeManagedDeviceExecuteActionRequestActionNameinitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      CreateMeManagedDeviceExecuteActionRequestActionNameinitiateOnDemandProactiveRemediation,
		"movedevicetoorganizationalunit":            CreateMeManagedDeviceExecuteActionRequestActionNamemoveDeviceToOrganizationalUnit,
		"quickscan":                                 CreateMeManagedDeviceExecuteActionRequestActionNamequickScan,
		"rebootnow":                                 CreateMeManagedDeviceExecuteActionRequestActionNamerebootNow,
		"reenable":                                  CreateMeManagedDeviceExecuteActionRequestActionNamereenable,
		"retire":                                    CreateMeManagedDeviceExecuteActionRequestActionNameretire,
		"setdevicename":                             CreateMeManagedDeviceExecuteActionRequestActionNamesetDeviceName,
		"signatureupdate":                           CreateMeManagedDeviceExecuteActionRequestActionNamesignatureUpdate,
		"syncdevice":                                CreateMeManagedDeviceExecuteActionRequestActionNamesyncDevice,
		"wipe":                                      CreateMeManagedDeviceExecuteActionRequestActionNamewipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeManagedDeviceExecuteActionRequestActionName(input)
	return &out, nil
}
