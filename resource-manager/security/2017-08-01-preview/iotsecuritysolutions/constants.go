package iotsecuritysolutions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSource string

const (
	DataSourceTwinData DataSource = "TwinData"
)

func PossibleValuesForDataSource() []string {
	return []string{
		string(DataSourceTwinData),
	}
}

func (s *DataSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataSource(input string) (*DataSource, error) {
	vals := map[string]DataSource{
		"twindata": DataSourceTwinData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSource(input)
	return &out, nil
}

type ExportData string

const (
	ExportDataRawEvents ExportData = "RawEvents"
)

func PossibleValuesForExportData() []string {
	return []string{
		string(ExportDataRawEvents),
	}
}

func (s *ExportData) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExportData(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExportData(input string) (*ExportData, error) {
	vals := map[string]ExportData{
		"rawevents": ExportDataRawEvents,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportData(input)
	return &out, nil
}

type RecommendationConfigStatus string

const (
	RecommendationConfigStatusDisabled RecommendationConfigStatus = "Disabled"
	RecommendationConfigStatusEnabled  RecommendationConfigStatus = "Enabled"
)

func PossibleValuesForRecommendationConfigStatus() []string {
	return []string{
		string(RecommendationConfigStatusDisabled),
		string(RecommendationConfigStatusEnabled),
	}
}

func (s *RecommendationConfigStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationConfigStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationConfigStatus(input string) (*RecommendationConfigStatus, error) {
	vals := map[string]RecommendationConfigStatus{
		"disabled": RecommendationConfigStatusDisabled,
		"enabled":  RecommendationConfigStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationConfigStatus(input)
	return &out, nil
}

type RecommendationType string

const (
	RecommendationTypeIoTACRAuthentication             RecommendationType = "IoT_ACRAuthentication"
	RecommendationTypeIoTAgentSendsUnutilizedMessages  RecommendationType = "IoT_AgentSendsUnutilizedMessages"
	RecommendationTypeIoTBaseline                      RecommendationType = "IoT_Baseline"
	RecommendationTypeIoTEdgeHubMemOptimize            RecommendationType = "IoT_EdgeHubMemOptimize"
	RecommendationTypeIoTEdgeLoggingOptions            RecommendationType = "IoT_EdgeLoggingOptions"
	RecommendationTypeIoTIPFilterDenyAll               RecommendationType = "IoT_IPFilter_DenyAll"
	RecommendationTypeIoTIPFilterPermissiveRule        RecommendationType = "IoT_IPFilter_PermissiveRule"
	RecommendationTypeIoTInconsistentModuleSettings    RecommendationType = "IoT_InconsistentModuleSettings"
	RecommendationTypeIoTInstallAgent                  RecommendationType = "IoT_InstallAgent"
	RecommendationTypeIoTOpenPorts                     RecommendationType = "IoT_OpenPorts"
	RecommendationTypeIoTPermissiveFirewallPolicy      RecommendationType = "IoT_PermissiveFirewallPolicy"
	RecommendationTypeIoTPermissiveInputFirewallRules  RecommendationType = "IoT_PermissiveInputFirewallRules"
	RecommendationTypeIoTPermissiveOutputFirewallRules RecommendationType = "IoT_PermissiveOutputFirewallRules"
	RecommendationTypeIoTPrivilegedDockerOptions       RecommendationType = "IoT_PrivilegedDockerOptions"
	RecommendationTypeIoTSharedCredentials             RecommendationType = "IoT_SharedCredentials"
	RecommendationTypeIoTVulnerableTLSCipherSuite      RecommendationType = "IoT_VulnerableTLSCipherSuite"
)

func PossibleValuesForRecommendationType() []string {
	return []string{
		string(RecommendationTypeIoTACRAuthentication),
		string(RecommendationTypeIoTAgentSendsUnutilizedMessages),
		string(RecommendationTypeIoTBaseline),
		string(RecommendationTypeIoTEdgeHubMemOptimize),
		string(RecommendationTypeIoTEdgeLoggingOptions),
		string(RecommendationTypeIoTIPFilterDenyAll),
		string(RecommendationTypeIoTIPFilterPermissiveRule),
		string(RecommendationTypeIoTInconsistentModuleSettings),
		string(RecommendationTypeIoTInstallAgent),
		string(RecommendationTypeIoTOpenPorts),
		string(RecommendationTypeIoTPermissiveFirewallPolicy),
		string(RecommendationTypeIoTPermissiveInputFirewallRules),
		string(RecommendationTypeIoTPermissiveOutputFirewallRules),
		string(RecommendationTypeIoTPrivilegedDockerOptions),
		string(RecommendationTypeIoTSharedCredentials),
		string(RecommendationTypeIoTVulnerableTLSCipherSuite),
	}
}

func (s *RecommendationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationType(input string) (*RecommendationType, error) {
	vals := map[string]RecommendationType{
		"iot_acrauthentication":             RecommendationTypeIoTACRAuthentication,
		"iot_agentsendsunutilizedmessages":  RecommendationTypeIoTAgentSendsUnutilizedMessages,
		"iot_baseline":                      RecommendationTypeIoTBaseline,
		"iot_edgehubmemoptimize":            RecommendationTypeIoTEdgeHubMemOptimize,
		"iot_edgeloggingoptions":            RecommendationTypeIoTEdgeLoggingOptions,
		"iot_ipfilter_denyall":              RecommendationTypeIoTIPFilterDenyAll,
		"iot_ipfilter_permissiverule":       RecommendationTypeIoTIPFilterPermissiveRule,
		"iot_inconsistentmodulesettings":    RecommendationTypeIoTInconsistentModuleSettings,
		"iot_installagent":                  RecommendationTypeIoTInstallAgent,
		"iot_openports":                     RecommendationTypeIoTOpenPorts,
		"iot_permissivefirewallpolicy":      RecommendationTypeIoTPermissiveFirewallPolicy,
		"iot_permissiveinputfirewallrules":  RecommendationTypeIoTPermissiveInputFirewallRules,
		"iot_permissiveoutputfirewallrules": RecommendationTypeIoTPermissiveOutputFirewallRules,
		"iot_privilegeddockeroptions":       RecommendationTypeIoTPrivilegedDockerOptions,
		"iot_sharedcredentials":             RecommendationTypeIoTSharedCredentials,
		"iot_vulnerabletlsciphersuite":      RecommendationTypeIoTVulnerableTLSCipherSuite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationType(input)
	return &out, nil
}

type SecuritySolutionStatus string

const (
	SecuritySolutionStatusDisabled SecuritySolutionStatus = "Disabled"
	SecuritySolutionStatusEnabled  SecuritySolutionStatus = "Enabled"
)

func PossibleValuesForSecuritySolutionStatus() []string {
	return []string{
		string(SecuritySolutionStatusDisabled),
		string(SecuritySolutionStatusEnabled),
	}
}

func (s *SecuritySolutionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySolutionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySolutionStatus(input string) (*SecuritySolutionStatus, error) {
	vals := map[string]SecuritySolutionStatus{
		"disabled": SecuritySolutionStatusDisabled,
		"enabled":  SecuritySolutionStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySolutionStatus(input)
	return &out, nil
}
