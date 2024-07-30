package deviceconfiguration

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

type DeviceConfigurationAssignmentIntent string

const (
	DeviceConfigurationAssignmentIntentApply  DeviceConfigurationAssignmentIntent = "apply"
	DeviceConfigurationAssignmentIntentRemove DeviceConfigurationAssignmentIntent = "remove"
)

func PossibleValuesForDeviceConfigurationAssignmentIntent() []string {
	return []string{
		string(DeviceConfigurationAssignmentIntentApply),
		string(DeviceConfigurationAssignmentIntentRemove),
	}
}

func (s *DeviceConfigurationAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationAssignmentIntent(input string) (*DeviceConfigurationAssignmentIntent, error) {
	vals := map[string]DeviceConfigurationAssignmentIntent{
		"apply":  DeviceConfigurationAssignmentIntentApply,
		"remove": DeviceConfigurationAssignmentIntentRemove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentIntent(input)
	return &out, nil
}

type DeviceConfigurationAssignmentSource string

const (
	DeviceConfigurationAssignmentSourceDirect     DeviceConfigurationAssignmentSource = "direct"
	DeviceConfigurationAssignmentSourcePolicySets DeviceConfigurationAssignmentSource = "policySets"
)

func PossibleValuesForDeviceConfigurationAssignmentSource() []string {
	return []string{
		string(DeviceConfigurationAssignmentSourceDirect),
		string(DeviceConfigurationAssignmentSourcePolicySets),
	}
}

func (s *DeviceConfigurationAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationAssignmentSource(input string) (*DeviceConfigurationAssignmentSource, error) {
	vals := map[string]DeviceConfigurationAssignmentSource{
		"direct":     DeviceConfigurationAssignmentSourceDirect,
		"policysets": DeviceConfigurationAssignmentSourcePolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentSource(input)
	return &out, nil
}

type DeviceConfigurationDeviceStatusStatus string

const (
	DeviceConfigurationDeviceStatusStatusCompliant     DeviceConfigurationDeviceStatusStatus = "compliant"
	DeviceConfigurationDeviceStatusStatusConflict      DeviceConfigurationDeviceStatusStatus = "conflict"
	DeviceConfigurationDeviceStatusStatusError         DeviceConfigurationDeviceStatusStatus = "error"
	DeviceConfigurationDeviceStatusStatusNonCompliant  DeviceConfigurationDeviceStatusStatus = "nonCompliant"
	DeviceConfigurationDeviceStatusStatusNotApplicable DeviceConfigurationDeviceStatusStatus = "notApplicable"
	DeviceConfigurationDeviceStatusStatusNotAssigned   DeviceConfigurationDeviceStatusStatus = "notAssigned"
	DeviceConfigurationDeviceStatusStatusRemediated    DeviceConfigurationDeviceStatusStatus = "remediated"
	DeviceConfigurationDeviceStatusStatusUnknown       DeviceConfigurationDeviceStatusStatus = "unknown"
)

func PossibleValuesForDeviceConfigurationDeviceStatusStatus() []string {
	return []string{
		string(DeviceConfigurationDeviceStatusStatusCompliant),
		string(DeviceConfigurationDeviceStatusStatusConflict),
		string(DeviceConfigurationDeviceStatusStatusError),
		string(DeviceConfigurationDeviceStatusStatusNonCompliant),
		string(DeviceConfigurationDeviceStatusStatusNotApplicable),
		string(DeviceConfigurationDeviceStatusStatusNotAssigned),
		string(DeviceConfigurationDeviceStatusStatusRemediated),
		string(DeviceConfigurationDeviceStatusStatusUnknown),
	}
}

func (s *DeviceConfigurationDeviceStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationDeviceStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationDeviceStatusStatus(input string) (*DeviceConfigurationDeviceStatusStatus, error) {
	vals := map[string]DeviceConfigurationDeviceStatusStatus{
		"compliant":     DeviceConfigurationDeviceStatusStatusCompliant,
		"conflict":      DeviceConfigurationDeviceStatusStatusConflict,
		"error":         DeviceConfigurationDeviceStatusStatusError,
		"noncompliant":  DeviceConfigurationDeviceStatusStatusNonCompliant,
		"notapplicable": DeviceConfigurationDeviceStatusStatusNotApplicable,
		"notassigned":   DeviceConfigurationDeviceStatusStatusNotAssigned,
		"remediated":    DeviceConfigurationDeviceStatusStatusRemediated,
		"unknown":       DeviceConfigurationDeviceStatusStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationDeviceStatusStatus(input)
	return &out, nil
}

type DeviceConfigurationUserStatusStatus string

const (
	DeviceConfigurationUserStatusStatusCompliant     DeviceConfigurationUserStatusStatus = "compliant"
	DeviceConfigurationUserStatusStatusConflict      DeviceConfigurationUserStatusStatus = "conflict"
	DeviceConfigurationUserStatusStatusError         DeviceConfigurationUserStatusStatus = "error"
	DeviceConfigurationUserStatusStatusNonCompliant  DeviceConfigurationUserStatusStatus = "nonCompliant"
	DeviceConfigurationUserStatusStatusNotApplicable DeviceConfigurationUserStatusStatus = "notApplicable"
	DeviceConfigurationUserStatusStatusNotAssigned   DeviceConfigurationUserStatusStatus = "notAssigned"
	DeviceConfigurationUserStatusStatusRemediated    DeviceConfigurationUserStatusStatus = "remediated"
	DeviceConfigurationUserStatusStatusUnknown       DeviceConfigurationUserStatusStatus = "unknown"
)

func PossibleValuesForDeviceConfigurationUserStatusStatus() []string {
	return []string{
		string(DeviceConfigurationUserStatusStatusCompliant),
		string(DeviceConfigurationUserStatusStatusConflict),
		string(DeviceConfigurationUserStatusStatusError),
		string(DeviceConfigurationUserStatusStatusNonCompliant),
		string(DeviceConfigurationUserStatusStatusNotApplicable),
		string(DeviceConfigurationUserStatusStatusNotAssigned),
		string(DeviceConfigurationUserStatusStatusRemediated),
		string(DeviceConfigurationUserStatusStatusUnknown),
	}
}

func (s *DeviceConfigurationUserStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationUserStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationUserStatusStatus(input string) (*DeviceConfigurationUserStatusStatus, error) {
	vals := map[string]DeviceConfigurationUserStatusStatus{
		"compliant":     DeviceConfigurationUserStatusStatusCompliant,
		"conflict":      DeviceConfigurationUserStatusStatusConflict,
		"error":         DeviceConfigurationUserStatusStatusError,
		"noncompliant":  DeviceConfigurationUserStatusStatusNonCompliant,
		"notapplicable": DeviceConfigurationUserStatusStatusNotApplicable,
		"notassigned":   DeviceConfigurationUserStatusStatusNotAssigned,
		"remediated":    DeviceConfigurationUserStatusStatusRemediated,
		"unknown":       DeviceConfigurationUserStatusStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationUserStatusStatus(input)
	return &out, nil
}

type DeviceManagementApplicabilityRuleDeviceModeDeviceMode string

const (
	DeviceManagementApplicabilityRuleDeviceModeDeviceModeSModeConfiguration    DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "sModeConfiguration"
	DeviceManagementApplicabilityRuleDeviceModeDeviceModeStandardConfiguration DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "standardConfiguration"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeDeviceMode() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceModeSModeConfiguration),
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceModeStandardConfiguration),
	}
}

func (s *DeviceManagementApplicabilityRuleDeviceModeDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleDeviceModeDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleDeviceModeDeviceMode(input string) (*DeviceManagementApplicabilityRuleDeviceModeDeviceMode, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeDeviceMode{
		"smodeconfiguration":    DeviceManagementApplicabilityRuleDeviceModeDeviceModeSModeConfiguration,
		"standardconfiguration": DeviceManagementApplicabilityRuleDeviceModeDeviceModeStandardConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeDeviceMode(input)
	return &out, nil
}

type DeviceManagementApplicabilityRuleDeviceModeRuleType string

const (
	DeviceManagementApplicabilityRuleDeviceModeRuleTypeExclude DeviceManagementApplicabilityRuleDeviceModeRuleType = "exclude"
	DeviceManagementApplicabilityRuleDeviceModeRuleTypeInclude DeviceManagementApplicabilityRuleDeviceModeRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeRuleTypeExclude),
		string(DeviceManagementApplicabilityRuleDeviceModeRuleTypeInclude),
	}
}

func (s *DeviceManagementApplicabilityRuleDeviceModeRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleDeviceModeRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleDeviceModeRuleType(input string) (*DeviceManagementApplicabilityRuleDeviceModeRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeRuleType{
		"exclude": DeviceManagementApplicabilityRuleDeviceModeRuleTypeExclude,
		"include": DeviceManagementApplicabilityRuleDeviceModeRuleTypeInclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeRuleType(input)
	return &out, nil
}

type DeviceManagementApplicabilityRuleOsEditionOsEditionTypes string

const (
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesNotConfigured                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "notConfigured"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Education                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Education"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EducationN               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10EducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Enterprise               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Enterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EnterpriseN              DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10EnterpriseN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HolographicEnterprise    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HolographicEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Home                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Home"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeChina                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeChina"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeN                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeSingleLanguage       DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeSingleLanguage"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCore                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10IoTCore"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCoreCommercial        DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10IoTCoreCommercial"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Mobile                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Mobile"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10MobileEnterprise         DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10MobileEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Professional             DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Professional"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducation    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalEducation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducationN   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalEducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalN            DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstation  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalWorkstation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstationN DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalWorkstationN"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionOsEditionTypes() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesNotConfigured),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Education),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Enterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EnterpriseN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HolographicEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Home),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeChina),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeSingleLanguage),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCore),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCoreCommercial),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Mobile),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10MobileEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Professional),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstationN),
	}
}

func (s *DeviceManagementApplicabilityRuleOsEditionOsEditionTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsEditionOsEditionTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input string) (*DeviceManagementApplicabilityRuleOsEditionOsEditionTypes, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionOsEditionTypes{
		"notconfigured":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypesNotConfigured,
		"windows10education":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Education,
		"windows10educationn":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EducationN,
		"windows10enterprise":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Enterprise,
		"windows10enterprisen":              DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10EnterpriseN,
		"windows10holographicenterprise":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HolographicEnterprise,
		"windows10home":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Home,
		"windows10homechina":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeChina,
		"windows10homen":                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeN,
		"windows10homesinglelanguage":       DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10HomeSingleLanguage,
		"windows10iotcore":                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCore,
		"windows10iotcorecommercial":        DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10IoTCoreCommercial,
		"windows10mobile":                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Mobile,
		"windows10mobileenterprise":         DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10MobileEnterprise,
		"windows10professional":             DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10Professional,
		"windows10professionaleducation":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducation,
		"windows10professionaleducationn":   DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalEducationN,
		"windows10professionaln":            DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalN,
		"windows10professionalworkstation":  DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstation,
		"windows10professionalworkstationn": DeviceManagementApplicabilityRuleOsEditionOsEditionTypesWindows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input)
	return &out, nil
}

type DeviceManagementApplicabilityRuleOsEditionRuleType string

const (
	DeviceManagementApplicabilityRuleOsEditionRuleTypeExclude DeviceManagementApplicabilityRuleOsEditionRuleType = "exclude"
	DeviceManagementApplicabilityRuleOsEditionRuleTypeInclude DeviceManagementApplicabilityRuleOsEditionRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionRuleTypeExclude),
		string(DeviceManagementApplicabilityRuleOsEditionRuleTypeInclude),
	}
}

func (s *DeviceManagementApplicabilityRuleOsEditionRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsEditionRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsEditionRuleType(input string) (*DeviceManagementApplicabilityRuleOsEditionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsEditionRuleTypeExclude,
		"include": DeviceManagementApplicabilityRuleOsEditionRuleTypeInclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionRuleType(input)
	return &out, nil
}

type DeviceManagementApplicabilityRuleOsVersionRuleType string

const (
	DeviceManagementApplicabilityRuleOsVersionRuleTypeExclude DeviceManagementApplicabilityRuleOsVersionRuleType = "exclude"
	DeviceManagementApplicabilityRuleOsVersionRuleTypeInclude DeviceManagementApplicabilityRuleOsVersionRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsVersionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsVersionRuleTypeExclude),
		string(DeviceManagementApplicabilityRuleOsVersionRuleTypeInclude),
	}
}

func (s *DeviceManagementApplicabilityRuleOsVersionRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsVersionRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsVersionRuleType(input string) (*DeviceManagementApplicabilityRuleOsVersionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsVersionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsVersionRuleTypeExclude,
		"include": DeviceManagementApplicabilityRuleOsVersionRuleTypeInclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsVersionRuleType(input)
	return &out, nil
}

type WindowsPrivacyDataAccessControlItemAccessLevel string

const (
	WindowsPrivacyDataAccessControlItemAccessLevelForceAllow    WindowsPrivacyDataAccessControlItemAccessLevel = "forceAllow"
	WindowsPrivacyDataAccessControlItemAccessLevelForceDeny     WindowsPrivacyDataAccessControlItemAccessLevel = "forceDeny"
	WindowsPrivacyDataAccessControlItemAccessLevelNotConfigured WindowsPrivacyDataAccessControlItemAccessLevel = "notConfigured"
	WindowsPrivacyDataAccessControlItemAccessLevelUserInControl WindowsPrivacyDataAccessControlItemAccessLevel = "userInControl"
)

func PossibleValuesForWindowsPrivacyDataAccessControlItemAccessLevel() []string {
	return []string{
		string(WindowsPrivacyDataAccessControlItemAccessLevelForceAllow),
		string(WindowsPrivacyDataAccessControlItemAccessLevelForceDeny),
		string(WindowsPrivacyDataAccessControlItemAccessLevelNotConfigured),
		string(WindowsPrivacyDataAccessControlItemAccessLevelUserInControl),
	}
}

func (s *WindowsPrivacyDataAccessControlItemAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataAccessControlItemAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataAccessControlItemAccessLevel(input string) (*WindowsPrivacyDataAccessControlItemAccessLevel, error) {
	vals := map[string]WindowsPrivacyDataAccessControlItemAccessLevel{
		"forceallow":    WindowsPrivacyDataAccessControlItemAccessLevelForceAllow,
		"forcedeny":     WindowsPrivacyDataAccessControlItemAccessLevelForceDeny,
		"notconfigured": WindowsPrivacyDataAccessControlItemAccessLevelNotConfigured,
		"userincontrol": WindowsPrivacyDataAccessControlItemAccessLevelUserInControl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessControlItemAccessLevel(input)
	return &out, nil
}

type WindowsPrivacyDataAccessControlItemDataCategory string

const (
	WindowsPrivacyDataAccessControlItemDataCategoryAccountInfo         WindowsPrivacyDataAccessControlItemDataCategory = "accountInfo"
	WindowsPrivacyDataAccessControlItemDataCategoryAppsRunInBackground WindowsPrivacyDataAccessControlItemDataCategory = "appsRunInBackground"
	WindowsPrivacyDataAccessControlItemDataCategoryCalendar            WindowsPrivacyDataAccessControlItemDataCategory = "calendar"
	WindowsPrivacyDataAccessControlItemDataCategoryCallHistory         WindowsPrivacyDataAccessControlItemDataCategory = "callHistory"
	WindowsPrivacyDataAccessControlItemDataCategoryCamera              WindowsPrivacyDataAccessControlItemDataCategory = "camera"
	WindowsPrivacyDataAccessControlItemDataCategoryContacts            WindowsPrivacyDataAccessControlItemDataCategory = "contacts"
	WindowsPrivacyDataAccessControlItemDataCategoryDiagnosticsInfo     WindowsPrivacyDataAccessControlItemDataCategory = "diagnosticsInfo"
	WindowsPrivacyDataAccessControlItemDataCategoryEmail               WindowsPrivacyDataAccessControlItemDataCategory = "email"
	WindowsPrivacyDataAccessControlItemDataCategoryLocation            WindowsPrivacyDataAccessControlItemDataCategory = "location"
	WindowsPrivacyDataAccessControlItemDataCategoryMessaging           WindowsPrivacyDataAccessControlItemDataCategory = "messaging"
	WindowsPrivacyDataAccessControlItemDataCategoryMicrophone          WindowsPrivacyDataAccessControlItemDataCategory = "microphone"
	WindowsPrivacyDataAccessControlItemDataCategoryMotion              WindowsPrivacyDataAccessControlItemDataCategory = "motion"
	WindowsPrivacyDataAccessControlItemDataCategoryNotConfigured       WindowsPrivacyDataAccessControlItemDataCategory = "notConfigured"
	WindowsPrivacyDataAccessControlItemDataCategoryNotifications       WindowsPrivacyDataAccessControlItemDataCategory = "notifications"
	WindowsPrivacyDataAccessControlItemDataCategoryPhone               WindowsPrivacyDataAccessControlItemDataCategory = "phone"
	WindowsPrivacyDataAccessControlItemDataCategoryRadios              WindowsPrivacyDataAccessControlItemDataCategory = "radios"
	WindowsPrivacyDataAccessControlItemDataCategorySyncWithDevices     WindowsPrivacyDataAccessControlItemDataCategory = "syncWithDevices"
	WindowsPrivacyDataAccessControlItemDataCategoryTasks               WindowsPrivacyDataAccessControlItemDataCategory = "tasks"
	WindowsPrivacyDataAccessControlItemDataCategoryTrustedDevices      WindowsPrivacyDataAccessControlItemDataCategory = "trustedDevices"
)

func PossibleValuesForWindowsPrivacyDataAccessControlItemDataCategory() []string {
	return []string{
		string(WindowsPrivacyDataAccessControlItemDataCategoryAccountInfo),
		string(WindowsPrivacyDataAccessControlItemDataCategoryAppsRunInBackground),
		string(WindowsPrivacyDataAccessControlItemDataCategoryCalendar),
		string(WindowsPrivacyDataAccessControlItemDataCategoryCallHistory),
		string(WindowsPrivacyDataAccessControlItemDataCategoryCamera),
		string(WindowsPrivacyDataAccessControlItemDataCategoryContacts),
		string(WindowsPrivacyDataAccessControlItemDataCategoryDiagnosticsInfo),
		string(WindowsPrivacyDataAccessControlItemDataCategoryEmail),
		string(WindowsPrivacyDataAccessControlItemDataCategoryLocation),
		string(WindowsPrivacyDataAccessControlItemDataCategoryMessaging),
		string(WindowsPrivacyDataAccessControlItemDataCategoryMicrophone),
		string(WindowsPrivacyDataAccessControlItemDataCategoryMotion),
		string(WindowsPrivacyDataAccessControlItemDataCategoryNotConfigured),
		string(WindowsPrivacyDataAccessControlItemDataCategoryNotifications),
		string(WindowsPrivacyDataAccessControlItemDataCategoryPhone),
		string(WindowsPrivacyDataAccessControlItemDataCategoryRadios),
		string(WindowsPrivacyDataAccessControlItemDataCategorySyncWithDevices),
		string(WindowsPrivacyDataAccessControlItemDataCategoryTasks),
		string(WindowsPrivacyDataAccessControlItemDataCategoryTrustedDevices),
	}
}

func (s *WindowsPrivacyDataAccessControlItemDataCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataAccessControlItemDataCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataAccessControlItemDataCategory(input string) (*WindowsPrivacyDataAccessControlItemDataCategory, error) {
	vals := map[string]WindowsPrivacyDataAccessControlItemDataCategory{
		"accountinfo":         WindowsPrivacyDataAccessControlItemDataCategoryAccountInfo,
		"appsruninbackground": WindowsPrivacyDataAccessControlItemDataCategoryAppsRunInBackground,
		"calendar":            WindowsPrivacyDataAccessControlItemDataCategoryCalendar,
		"callhistory":         WindowsPrivacyDataAccessControlItemDataCategoryCallHistory,
		"camera":              WindowsPrivacyDataAccessControlItemDataCategoryCamera,
		"contacts":            WindowsPrivacyDataAccessControlItemDataCategoryContacts,
		"diagnosticsinfo":     WindowsPrivacyDataAccessControlItemDataCategoryDiagnosticsInfo,
		"email":               WindowsPrivacyDataAccessControlItemDataCategoryEmail,
		"location":            WindowsPrivacyDataAccessControlItemDataCategoryLocation,
		"messaging":           WindowsPrivacyDataAccessControlItemDataCategoryMessaging,
		"microphone":          WindowsPrivacyDataAccessControlItemDataCategoryMicrophone,
		"motion":              WindowsPrivacyDataAccessControlItemDataCategoryMotion,
		"notconfigured":       WindowsPrivacyDataAccessControlItemDataCategoryNotConfigured,
		"notifications":       WindowsPrivacyDataAccessControlItemDataCategoryNotifications,
		"phone":               WindowsPrivacyDataAccessControlItemDataCategoryPhone,
		"radios":              WindowsPrivacyDataAccessControlItemDataCategoryRadios,
		"syncwithdevices":     WindowsPrivacyDataAccessControlItemDataCategorySyncWithDevices,
		"tasks":               WindowsPrivacyDataAccessControlItemDataCategoryTasks,
		"trusteddevices":      WindowsPrivacyDataAccessControlItemDataCategoryTrustedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessControlItemDataCategory(input)
	return &out, nil
}
