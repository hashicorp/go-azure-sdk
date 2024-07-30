package comanageddevice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatusUserAccessLevel string

const (
	CloudPCReviewStatusUserAccessLevelRestricted   CloudPCReviewStatusUserAccessLevel = "restricted"
	CloudPCReviewStatusUserAccessLevelUnrestricted CloudPCReviewStatusUserAccessLevel = "unrestricted"
)

func PossibleValuesForCloudPCReviewStatusUserAccessLevel() []string {
	return []string{
		string(CloudPCReviewStatusUserAccessLevelRestricted),
		string(CloudPCReviewStatusUserAccessLevelUnrestricted),
	}
}

func (s *CloudPCReviewStatusUserAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCReviewStatusUserAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCReviewStatusUserAccessLevel(input string) (*CloudPCReviewStatusUserAccessLevel, error) {
	vals := map[string]CloudPCReviewStatusUserAccessLevel{
		"restricted":   CloudPCReviewStatusUserAccessLevelRestricted,
		"unrestricted": CloudPCReviewStatusUserAccessLevelUnrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCReviewStatusUserAccessLevel(input)
	return &out, nil
}

type ConfigurationManagerActionAction string

const (
	ConfigurationManagerActionActionAppEvaluation                   ConfigurationManagerActionAction = "appEvaluation"
	ConfigurationManagerActionActionFullScan                        ConfigurationManagerActionAction = "fullScan"
	ConfigurationManagerActionActionQuickScan                       ConfigurationManagerActionAction = "quickScan"
	ConfigurationManagerActionActionRefreshMachinePolicy            ConfigurationManagerActionAction = "refreshMachinePolicy"
	ConfigurationManagerActionActionRefreshUserPolicy               ConfigurationManagerActionAction = "refreshUserPolicy"
	ConfigurationManagerActionActionWakeUpClient                    ConfigurationManagerActionAction = "wakeUpClient"
	ConfigurationManagerActionActionWindowsDefenderUpdateSignatures ConfigurationManagerActionAction = "windowsDefenderUpdateSignatures"
)

func PossibleValuesForConfigurationManagerActionAction() []string {
	return []string{
		string(ConfigurationManagerActionActionAppEvaluation),
		string(ConfigurationManagerActionActionFullScan),
		string(ConfigurationManagerActionActionQuickScan),
		string(ConfigurationManagerActionActionRefreshMachinePolicy),
		string(ConfigurationManagerActionActionRefreshUserPolicy),
		string(ConfigurationManagerActionActionWakeUpClient),
		string(ConfigurationManagerActionActionWindowsDefenderUpdateSignatures),
	}
}

func (s *ConfigurationManagerActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionAction(input string) (*ConfigurationManagerActionAction, error) {
	vals := map[string]ConfigurationManagerActionAction{
		"appevaluation":                   ConfigurationManagerActionActionAppEvaluation,
		"fullscan":                        ConfigurationManagerActionActionFullScan,
		"quickscan":                       ConfigurationManagerActionActionQuickScan,
		"refreshmachinepolicy":            ConfigurationManagerActionActionRefreshMachinePolicy,
		"refreshuserpolicy":               ConfigurationManagerActionActionRefreshUserPolicy,
		"wakeupclient":                    ConfigurationManagerActionActionWakeUpClient,
		"windowsdefenderupdatesignatures": ConfigurationManagerActionActionWindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionAction(input)
	return &out, nil
}

type CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange string

const (
	CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter         CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange = "after"
	CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore        CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange = "before"
	CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange = "beforeOrAfter"
)

func PossibleValuesForCreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange() []string {
	return []string{
		string(CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter),
		string(CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore),
		string(CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter),
	}
}

func (s *CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange(input string) (*CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange, error) {
	vals := map[string]CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange{
		"after":         CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter,
		"before":        CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore,
		"beforeorafter": CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange(input)
	return &out, nil
}

type CreateComanagedDeviceExecuteActionRequestActionName string

const (
	CreateComanagedDeviceExecuteActionRequestActionNameActivateDeviceEsim                        CreateComanagedDeviceExecuteActionRequestActionName = "activateDeviceEsim"
	CreateComanagedDeviceExecuteActionRequestActionNameCollectDiagnostics                        CreateComanagedDeviceExecuteActionRequestActionName = "collectDiagnostics"
	CreateComanagedDeviceExecuteActionRequestActionNameCustomTextNotification                    CreateComanagedDeviceExecuteActionRequestActionName = "customTextNotification"
	CreateComanagedDeviceExecuteActionRequestActionNameDelete                                    CreateComanagedDeviceExecuteActionRequestActionName = "delete"
	CreateComanagedDeviceExecuteActionRequestActionNameDeprovision                               CreateComanagedDeviceExecuteActionRequestActionName = "deprovision"
	CreateComanagedDeviceExecuteActionRequestActionNameDisable                                   CreateComanagedDeviceExecuteActionRequestActionName = "disable"
	CreateComanagedDeviceExecuteActionRequestActionNameFullScan                                  CreateComanagedDeviceExecuteActionRequestActionName = "fullScan"
	CreateComanagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery CreateComanagedDeviceExecuteActionRequestActionName = "initiateMobileDeviceManagementKeyRecovery"
	CreateComanagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation      CreateComanagedDeviceExecuteActionRequestActionName = "initiateOnDemandProactiveRemediation"
	CreateComanagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit            CreateComanagedDeviceExecuteActionRequestActionName = "moveDeviceToOrganizationalUnit"
	CreateComanagedDeviceExecuteActionRequestActionNameQuickScan                                 CreateComanagedDeviceExecuteActionRequestActionName = "quickScan"
	CreateComanagedDeviceExecuteActionRequestActionNameRebootNow                                 CreateComanagedDeviceExecuteActionRequestActionName = "rebootNow"
	CreateComanagedDeviceExecuteActionRequestActionNameReenable                                  CreateComanagedDeviceExecuteActionRequestActionName = "reenable"
	CreateComanagedDeviceExecuteActionRequestActionNameRetire                                    CreateComanagedDeviceExecuteActionRequestActionName = "retire"
	CreateComanagedDeviceExecuteActionRequestActionNameSetDeviceName                             CreateComanagedDeviceExecuteActionRequestActionName = "setDeviceName"
	CreateComanagedDeviceExecuteActionRequestActionNameSignatureUpdate                           CreateComanagedDeviceExecuteActionRequestActionName = "signatureUpdate"
	CreateComanagedDeviceExecuteActionRequestActionNameSyncDevice                                CreateComanagedDeviceExecuteActionRequestActionName = "syncDevice"
	CreateComanagedDeviceExecuteActionRequestActionNameWipe                                      CreateComanagedDeviceExecuteActionRequestActionName = "wipe"
)

func PossibleValuesForCreateComanagedDeviceExecuteActionRequestActionName() []string {
	return []string{
		string(CreateComanagedDeviceExecuteActionRequestActionNameActivateDeviceEsim),
		string(CreateComanagedDeviceExecuteActionRequestActionNameCollectDiagnostics),
		string(CreateComanagedDeviceExecuteActionRequestActionNameCustomTextNotification),
		string(CreateComanagedDeviceExecuteActionRequestActionNameDelete),
		string(CreateComanagedDeviceExecuteActionRequestActionNameDeprovision),
		string(CreateComanagedDeviceExecuteActionRequestActionNameDisable),
		string(CreateComanagedDeviceExecuteActionRequestActionNameFullScan),
		string(CreateComanagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery),
		string(CreateComanagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation),
		string(CreateComanagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit),
		string(CreateComanagedDeviceExecuteActionRequestActionNameQuickScan),
		string(CreateComanagedDeviceExecuteActionRequestActionNameRebootNow),
		string(CreateComanagedDeviceExecuteActionRequestActionNameReenable),
		string(CreateComanagedDeviceExecuteActionRequestActionNameRetire),
		string(CreateComanagedDeviceExecuteActionRequestActionNameSetDeviceName),
		string(CreateComanagedDeviceExecuteActionRequestActionNameSignatureUpdate),
		string(CreateComanagedDeviceExecuteActionRequestActionNameSyncDevice),
		string(CreateComanagedDeviceExecuteActionRequestActionNameWipe),
	}
}

func (s *CreateComanagedDeviceExecuteActionRequestActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateComanagedDeviceExecuteActionRequestActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateComanagedDeviceExecuteActionRequestActionName(input string) (*CreateComanagedDeviceExecuteActionRequestActionName, error) {
	vals := map[string]CreateComanagedDeviceExecuteActionRequestActionName{
		"activatedeviceesim":     CreateComanagedDeviceExecuteActionRequestActionNameActivateDeviceEsim,
		"collectdiagnostics":     CreateComanagedDeviceExecuteActionRequestActionNameCollectDiagnostics,
		"customtextnotification": CreateComanagedDeviceExecuteActionRequestActionNameCustomTextNotification,
		"delete":                 CreateComanagedDeviceExecuteActionRequestActionNameDelete,
		"deprovision":            CreateComanagedDeviceExecuteActionRequestActionNameDeprovision,
		"disable":                CreateComanagedDeviceExecuteActionRequestActionNameDisable,
		"fullscan":               CreateComanagedDeviceExecuteActionRequestActionNameFullScan,
		"initiatemobiledevicemanagementkeyrecovery": CreateComanagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      CreateComanagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation,
		"movedevicetoorganizationalunit":            CreateComanagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit,
		"quickscan":                                 CreateComanagedDeviceExecuteActionRequestActionNameQuickScan,
		"rebootnow":                                 CreateComanagedDeviceExecuteActionRequestActionNameRebootNow,
		"reenable":                                  CreateComanagedDeviceExecuteActionRequestActionNameReenable,
		"retire":                                    CreateComanagedDeviceExecuteActionRequestActionNameRetire,
		"setdevicename":                             CreateComanagedDeviceExecuteActionRequestActionNameSetDeviceName,
		"signatureupdate":                           CreateComanagedDeviceExecuteActionRequestActionNameSignatureUpdate,
		"syncdevice":                                CreateComanagedDeviceExecuteActionRequestActionNameSyncDevice,
		"wipe":                                      CreateComanagedDeviceExecuteActionRequestActionNameWipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateComanagedDeviceExecuteActionRequestActionName(input)
	return &out, nil
}

type CreateComanagedDeviceOverrideComplianceStateRequestComplianceState string

const (
	CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy CreateComanagedDeviceOverrideComplianceStateRequestComplianceState = "basedOnDeviceCompliancePolicy"
	CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant                  CreateComanagedDeviceOverrideComplianceStateRequestComplianceState = "nonCompliant"
)

func PossibleValuesForCreateComanagedDeviceOverrideComplianceStateRequestComplianceState() []string {
	return []string{
		string(CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy),
		string(CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant),
	}
}

func (s *CreateComanagedDeviceOverrideComplianceStateRequestComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateComanagedDeviceOverrideComplianceStateRequestComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateComanagedDeviceOverrideComplianceStateRequestComplianceState(input string) (*CreateComanagedDeviceOverrideComplianceStateRequestComplianceState, error) {
	vals := map[string]CreateComanagedDeviceOverrideComplianceStateRequestComplianceState{
		"basedondevicecompliancepolicy": CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy,
		"noncompliant":                  CreateComanagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateComanagedDeviceOverrideComplianceStateRequestComplianceState(input)
	return &out, nil
}

type CreateComanagedDeviceWipeRequestObliterationBehavior string

const (
	CreateComanagedDeviceWipeRequestObliterationBehaviorAlways                CreateComanagedDeviceWipeRequestObliterationBehavior = "always"
	CreateComanagedDeviceWipeRequestObliterationBehaviorDefault               CreateComanagedDeviceWipeRequestObliterationBehavior = "default"
	CreateComanagedDeviceWipeRequestObliterationBehaviorDoNotObliterate       CreateComanagedDeviceWipeRequestObliterationBehavior = "doNotObliterate"
	CreateComanagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning CreateComanagedDeviceWipeRequestObliterationBehavior = "obliterateWithWarning"
)

func PossibleValuesForCreateComanagedDeviceWipeRequestObliterationBehavior() []string {
	return []string{
		string(CreateComanagedDeviceWipeRequestObliterationBehaviorAlways),
		string(CreateComanagedDeviceWipeRequestObliterationBehaviorDefault),
		string(CreateComanagedDeviceWipeRequestObliterationBehaviorDoNotObliterate),
		string(CreateComanagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning),
	}
}

func (s *CreateComanagedDeviceWipeRequestObliterationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateComanagedDeviceWipeRequestObliterationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateComanagedDeviceWipeRequestObliterationBehavior(input string) (*CreateComanagedDeviceWipeRequestObliterationBehavior, error) {
	vals := map[string]CreateComanagedDeviceWipeRequestObliterationBehavior{
		"always":                CreateComanagedDeviceWipeRequestObliterationBehaviorAlways,
		"default":               CreateComanagedDeviceWipeRequestObliterationBehaviorDefault,
		"donotobliterate":       CreateComanagedDeviceWipeRequestObliterationBehaviorDoNotObliterate,
		"obliteratewithwarning": CreateComanagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateComanagedDeviceWipeRequestObliterationBehavior(input)
	return &out, nil
}

type DeviceLogCollectionRequestTemplateType string

const (
	DeviceLogCollectionRequestTemplateTypePredefined DeviceLogCollectionRequestTemplateType = "predefined"
)

func PossibleValuesForDeviceLogCollectionRequestTemplateType() []string {
	return []string{
		string(DeviceLogCollectionRequestTemplateTypePredefined),
	}
}

func (s *DeviceLogCollectionRequestTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLogCollectionRequestTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLogCollectionRequestTemplateType(input string) (*DeviceLogCollectionRequestTemplateType, error) {
	vals := map[string]DeviceLogCollectionRequestTemplateType{
		"predefined": DeviceLogCollectionRequestTemplateTypePredefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionRequestTemplateType(input)
	return &out, nil
}
