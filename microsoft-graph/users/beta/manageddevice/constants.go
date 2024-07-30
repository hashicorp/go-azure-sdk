package manageddevice

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

type CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange string

const (
	CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter         CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange = "after"
	CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore        CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange = "before"
	CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange = "beforeOrAfter"
)

func PossibleValuesForCreateManagedDeviceBulkRestoreCloudPCRequestTimeRange() []string {
	return []string{
		string(CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter),
		string(CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore),
		string(CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter),
	}
}

func (s *CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateManagedDeviceBulkRestoreCloudPCRequestTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateManagedDeviceBulkRestoreCloudPCRequestTimeRange(input string) (*CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange, error) {
	vals := map[string]CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange{
		"after":         CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeAfter,
		"before":        CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBefore,
		"beforeorafter": CreateManagedDeviceBulkRestoreCloudPCRequestTimeRangeBeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateManagedDeviceBulkRestoreCloudPCRequestTimeRange(input)
	return &out, nil
}

type CreateManagedDeviceExecuteActionRequestActionName string

const (
	CreateManagedDeviceExecuteActionRequestActionNameActivateDeviceEsim                        CreateManagedDeviceExecuteActionRequestActionName = "activateDeviceEsim"
	CreateManagedDeviceExecuteActionRequestActionNameCollectDiagnostics                        CreateManagedDeviceExecuteActionRequestActionName = "collectDiagnostics"
	CreateManagedDeviceExecuteActionRequestActionNameCustomTextNotification                    CreateManagedDeviceExecuteActionRequestActionName = "customTextNotification"
	CreateManagedDeviceExecuteActionRequestActionNameDelete                                    CreateManagedDeviceExecuteActionRequestActionName = "delete"
	CreateManagedDeviceExecuteActionRequestActionNameDeprovision                               CreateManagedDeviceExecuteActionRequestActionName = "deprovision"
	CreateManagedDeviceExecuteActionRequestActionNameDisable                                   CreateManagedDeviceExecuteActionRequestActionName = "disable"
	CreateManagedDeviceExecuteActionRequestActionNameFullScan                                  CreateManagedDeviceExecuteActionRequestActionName = "fullScan"
	CreateManagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery CreateManagedDeviceExecuteActionRequestActionName = "initiateMobileDeviceManagementKeyRecovery"
	CreateManagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation      CreateManagedDeviceExecuteActionRequestActionName = "initiateOnDemandProactiveRemediation"
	CreateManagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit            CreateManagedDeviceExecuteActionRequestActionName = "moveDeviceToOrganizationalUnit"
	CreateManagedDeviceExecuteActionRequestActionNameQuickScan                                 CreateManagedDeviceExecuteActionRequestActionName = "quickScan"
	CreateManagedDeviceExecuteActionRequestActionNameRebootNow                                 CreateManagedDeviceExecuteActionRequestActionName = "rebootNow"
	CreateManagedDeviceExecuteActionRequestActionNameReenable                                  CreateManagedDeviceExecuteActionRequestActionName = "reenable"
	CreateManagedDeviceExecuteActionRequestActionNameRetire                                    CreateManagedDeviceExecuteActionRequestActionName = "retire"
	CreateManagedDeviceExecuteActionRequestActionNameSetDeviceName                             CreateManagedDeviceExecuteActionRequestActionName = "setDeviceName"
	CreateManagedDeviceExecuteActionRequestActionNameSignatureUpdate                           CreateManagedDeviceExecuteActionRequestActionName = "signatureUpdate"
	CreateManagedDeviceExecuteActionRequestActionNameSyncDevice                                CreateManagedDeviceExecuteActionRequestActionName = "syncDevice"
	CreateManagedDeviceExecuteActionRequestActionNameWipe                                      CreateManagedDeviceExecuteActionRequestActionName = "wipe"
)

func PossibleValuesForCreateManagedDeviceExecuteActionRequestActionName() []string {
	return []string{
		string(CreateManagedDeviceExecuteActionRequestActionNameActivateDeviceEsim),
		string(CreateManagedDeviceExecuteActionRequestActionNameCollectDiagnostics),
		string(CreateManagedDeviceExecuteActionRequestActionNameCustomTextNotification),
		string(CreateManagedDeviceExecuteActionRequestActionNameDelete),
		string(CreateManagedDeviceExecuteActionRequestActionNameDeprovision),
		string(CreateManagedDeviceExecuteActionRequestActionNameDisable),
		string(CreateManagedDeviceExecuteActionRequestActionNameFullScan),
		string(CreateManagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery),
		string(CreateManagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation),
		string(CreateManagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit),
		string(CreateManagedDeviceExecuteActionRequestActionNameQuickScan),
		string(CreateManagedDeviceExecuteActionRequestActionNameRebootNow),
		string(CreateManagedDeviceExecuteActionRequestActionNameReenable),
		string(CreateManagedDeviceExecuteActionRequestActionNameRetire),
		string(CreateManagedDeviceExecuteActionRequestActionNameSetDeviceName),
		string(CreateManagedDeviceExecuteActionRequestActionNameSignatureUpdate),
		string(CreateManagedDeviceExecuteActionRequestActionNameSyncDevice),
		string(CreateManagedDeviceExecuteActionRequestActionNameWipe),
	}
}

func (s *CreateManagedDeviceExecuteActionRequestActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateManagedDeviceExecuteActionRequestActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateManagedDeviceExecuteActionRequestActionName(input string) (*CreateManagedDeviceExecuteActionRequestActionName, error) {
	vals := map[string]CreateManagedDeviceExecuteActionRequestActionName{
		"activatedeviceesim":     CreateManagedDeviceExecuteActionRequestActionNameActivateDeviceEsim,
		"collectdiagnostics":     CreateManagedDeviceExecuteActionRequestActionNameCollectDiagnostics,
		"customtextnotification": CreateManagedDeviceExecuteActionRequestActionNameCustomTextNotification,
		"delete":                 CreateManagedDeviceExecuteActionRequestActionNameDelete,
		"deprovision":            CreateManagedDeviceExecuteActionRequestActionNameDeprovision,
		"disable":                CreateManagedDeviceExecuteActionRequestActionNameDisable,
		"fullscan":               CreateManagedDeviceExecuteActionRequestActionNameFullScan,
		"initiatemobiledevicemanagementkeyrecovery": CreateManagedDeviceExecuteActionRequestActionNameInitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      CreateManagedDeviceExecuteActionRequestActionNameInitiateOnDemandProactiveRemediation,
		"movedevicetoorganizationalunit":            CreateManagedDeviceExecuteActionRequestActionNameMoveDeviceToOrganizationalUnit,
		"quickscan":                                 CreateManagedDeviceExecuteActionRequestActionNameQuickScan,
		"rebootnow":                                 CreateManagedDeviceExecuteActionRequestActionNameRebootNow,
		"reenable":                                  CreateManagedDeviceExecuteActionRequestActionNameReenable,
		"retire":                                    CreateManagedDeviceExecuteActionRequestActionNameRetire,
		"setdevicename":                             CreateManagedDeviceExecuteActionRequestActionNameSetDeviceName,
		"signatureupdate":                           CreateManagedDeviceExecuteActionRequestActionNameSignatureUpdate,
		"syncdevice":                                CreateManagedDeviceExecuteActionRequestActionNameSyncDevice,
		"wipe":                                      CreateManagedDeviceExecuteActionRequestActionNameWipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateManagedDeviceExecuteActionRequestActionName(input)
	return &out, nil
}

type CreateManagedDeviceOverrideComplianceStateRequestComplianceState string

const (
	CreateManagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy CreateManagedDeviceOverrideComplianceStateRequestComplianceState = "basedOnDeviceCompliancePolicy"
	CreateManagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant                  CreateManagedDeviceOverrideComplianceStateRequestComplianceState = "nonCompliant"
)

func PossibleValuesForCreateManagedDeviceOverrideComplianceStateRequestComplianceState() []string {
	return []string{
		string(CreateManagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy),
		string(CreateManagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant),
	}
}

func (s *CreateManagedDeviceOverrideComplianceStateRequestComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateManagedDeviceOverrideComplianceStateRequestComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateManagedDeviceOverrideComplianceStateRequestComplianceState(input string) (*CreateManagedDeviceOverrideComplianceStateRequestComplianceState, error) {
	vals := map[string]CreateManagedDeviceOverrideComplianceStateRequestComplianceState{
		"basedondevicecompliancepolicy": CreateManagedDeviceOverrideComplianceStateRequestComplianceStateBasedOnDeviceCompliancePolicy,
		"noncompliant":                  CreateManagedDeviceOverrideComplianceStateRequestComplianceStateNonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateManagedDeviceOverrideComplianceStateRequestComplianceState(input)
	return &out, nil
}

type CreateManagedDeviceWipeRequestObliterationBehavior string

const (
	CreateManagedDeviceWipeRequestObliterationBehaviorAlways                CreateManagedDeviceWipeRequestObliterationBehavior = "always"
	CreateManagedDeviceWipeRequestObliterationBehaviorDefault               CreateManagedDeviceWipeRequestObliterationBehavior = "default"
	CreateManagedDeviceWipeRequestObliterationBehaviorDoNotObliterate       CreateManagedDeviceWipeRequestObliterationBehavior = "doNotObliterate"
	CreateManagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning CreateManagedDeviceWipeRequestObliterationBehavior = "obliterateWithWarning"
)

func PossibleValuesForCreateManagedDeviceWipeRequestObliterationBehavior() []string {
	return []string{
		string(CreateManagedDeviceWipeRequestObliterationBehaviorAlways),
		string(CreateManagedDeviceWipeRequestObliterationBehaviorDefault),
		string(CreateManagedDeviceWipeRequestObliterationBehaviorDoNotObliterate),
		string(CreateManagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning),
	}
}

func (s *CreateManagedDeviceWipeRequestObliterationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateManagedDeviceWipeRequestObliterationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateManagedDeviceWipeRequestObliterationBehavior(input string) (*CreateManagedDeviceWipeRequestObliterationBehavior, error) {
	vals := map[string]CreateManagedDeviceWipeRequestObliterationBehavior{
		"always":                CreateManagedDeviceWipeRequestObliterationBehaviorAlways,
		"default":               CreateManagedDeviceWipeRequestObliterationBehaviorDefault,
		"donotobliterate":       CreateManagedDeviceWipeRequestObliterationBehaviorDoNotObliterate,
		"obliteratewithwarning": CreateManagedDeviceWipeRequestObliterationBehaviorObliterateWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateManagedDeviceWipeRequestObliterationBehavior(input)
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
