package securityconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudName string

const (
	CloudNameAWS         CloudName = "AWS"
	CloudNameAzure       CloudName = "Azure"
	CloudNameAzureDevOps CloudName = "AzureDevOps"
	CloudNameGCP         CloudName = "GCP"
	CloudNameGitLab      CloudName = "GitLab"
	CloudNameGithub      CloudName = "Github"
)

func PossibleValuesForCloudName() []string {
	return []string{
		string(CloudNameAWS),
		string(CloudNameAzure),
		string(CloudNameAzureDevOps),
		string(CloudNameGCP),
		string(CloudNameGitLab),
		string(CloudNameGithub),
	}
}

func (s *CloudName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudName(input string) (*CloudName, error) {
	vals := map[string]CloudName{
		"aws":         CloudNameAWS,
		"azure":       CloudNameAzure,
		"azuredevops": CloudNameAzureDevOps,
		"gcp":         CloudNameGCP,
		"gitlab":      CloudNameGitLab,
		"github":      CloudNameGithub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudName(input)
	return &out, nil
}

type EnvironmentType string

const (
	EnvironmentTypeAwsAccount       EnvironmentType = "AwsAccount"
	EnvironmentTypeAzureDevOpsScope EnvironmentType = "AzureDevOpsScope"
	EnvironmentTypeGcpProject       EnvironmentType = "GcpProject"
	EnvironmentTypeGithubScope      EnvironmentType = "GithubScope"
	EnvironmentTypeGitlabScope      EnvironmentType = "GitlabScope"
)

func PossibleValuesForEnvironmentType() []string {
	return []string{
		string(EnvironmentTypeAwsAccount),
		string(EnvironmentTypeAzureDevOpsScope),
		string(EnvironmentTypeGcpProject),
		string(EnvironmentTypeGithubScope),
		string(EnvironmentTypeGitlabScope),
	}
}

func (s *EnvironmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnvironmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnvironmentType(input string) (*EnvironmentType, error) {
	vals := map[string]EnvironmentType{
		"awsaccount":       EnvironmentTypeAwsAccount,
		"azuredevopsscope": EnvironmentTypeAzureDevOpsScope,
		"gcpproject":       EnvironmentTypeGcpProject,
		"githubscope":      EnvironmentTypeGithubScope,
		"gitlabscope":      EnvironmentTypeGitlabScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnvironmentType(input)
	return &out, nil
}

type OfferingType string

const (
	OfferingTypeCspmMonitorAws               OfferingType = "CspmMonitorAws"
	OfferingTypeCspmMonitorAzureDevOps       OfferingType = "CspmMonitorAzureDevOps"
	OfferingTypeCspmMonitorGcp               OfferingType = "CspmMonitorGcp"
	OfferingTypeCspmMonitorGitLab            OfferingType = "CspmMonitorGitLab"
	OfferingTypeCspmMonitorGithub            OfferingType = "CspmMonitorGithub"
	OfferingTypeDefenderCspmAws              OfferingType = "DefenderCspmAws"
	OfferingTypeDefenderCspmGcp              OfferingType = "DefenderCspmGcp"
	OfferingTypeDefenderForContainersAws     OfferingType = "DefenderForContainersAws"
	OfferingTypeDefenderForContainersGcp     OfferingType = "DefenderForContainersGcp"
	OfferingTypeDefenderForDatabasesAws      OfferingType = "DefenderForDatabasesAws"
	OfferingTypeDefenderForDatabasesGcp      OfferingType = "DefenderForDatabasesGcp"
	OfferingTypeDefenderForDevOpsAzureDevOps OfferingType = "DefenderForDevOpsAzureDevOps"
	OfferingTypeDefenderForDevOpsGitLab      OfferingType = "DefenderForDevOpsGitLab"
	OfferingTypeDefenderForDevOpsGithub      OfferingType = "DefenderForDevOpsGithub"
	OfferingTypeDefenderForServersAws        OfferingType = "DefenderForServersAws"
	OfferingTypeDefenderForServersGcp        OfferingType = "DefenderForServersGcp"
	OfferingTypeInformationProtectionAws     OfferingType = "InformationProtectionAws"
)

func PossibleValuesForOfferingType() []string {
	return []string{
		string(OfferingTypeCspmMonitorAws),
		string(OfferingTypeCspmMonitorAzureDevOps),
		string(OfferingTypeCspmMonitorGcp),
		string(OfferingTypeCspmMonitorGitLab),
		string(OfferingTypeCspmMonitorGithub),
		string(OfferingTypeDefenderCspmAws),
		string(OfferingTypeDefenderCspmGcp),
		string(OfferingTypeDefenderForContainersAws),
		string(OfferingTypeDefenderForContainersGcp),
		string(OfferingTypeDefenderForDatabasesAws),
		string(OfferingTypeDefenderForDatabasesGcp),
		string(OfferingTypeDefenderForDevOpsAzureDevOps),
		string(OfferingTypeDefenderForDevOpsGitLab),
		string(OfferingTypeDefenderForDevOpsGithub),
		string(OfferingTypeDefenderForServersAws),
		string(OfferingTypeDefenderForServersGcp),
		string(OfferingTypeInformationProtectionAws),
	}
}

func (s *OfferingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfferingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfferingType(input string) (*OfferingType, error) {
	vals := map[string]OfferingType{
		"cspmmonitoraws":               OfferingTypeCspmMonitorAws,
		"cspmmonitorazuredevops":       OfferingTypeCspmMonitorAzureDevOps,
		"cspmmonitorgcp":               OfferingTypeCspmMonitorGcp,
		"cspmmonitorgitlab":            OfferingTypeCspmMonitorGitLab,
		"cspmmonitorgithub":            OfferingTypeCspmMonitorGithub,
		"defendercspmaws":              OfferingTypeDefenderCspmAws,
		"defendercspmgcp":              OfferingTypeDefenderCspmGcp,
		"defenderforcontainersaws":     OfferingTypeDefenderForContainersAws,
		"defenderforcontainersgcp":     OfferingTypeDefenderForContainersGcp,
		"defenderfordatabasesaws":      OfferingTypeDefenderForDatabasesAws,
		"defenderfordatabasesgcp":      OfferingTypeDefenderForDatabasesGcp,
		"defenderfordevopsazuredevops": OfferingTypeDefenderForDevOpsAzureDevOps,
		"defenderfordevopsgitlab":      OfferingTypeDefenderForDevOpsGitLab,
		"defenderfordevopsgithub":      OfferingTypeDefenderForDevOpsGithub,
		"defenderforserversaws":        OfferingTypeDefenderForServersAws,
		"defenderforserversgcp":        OfferingTypeDefenderForServersGcp,
		"informationprotectionaws":     OfferingTypeInformationProtectionAws,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferingType(input)
	return &out, nil
}

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       OrganizationMembershipType = "Member"
	OrganizationMembershipTypeOrganization OrganizationMembershipType = "Organization"
)

func PossibleValuesForOrganizationMembershipType() []string {
	return []string{
		string(OrganizationMembershipTypeMember),
		string(OrganizationMembershipTypeOrganization),
	}
}

func (s *OrganizationMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrganizationMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrganizationMembershipType(input string) (*OrganizationMembershipType, error) {
	vals := map[string]OrganizationMembershipType{
		"member":       OrganizationMembershipTypeMember,
		"organization": OrganizationMembershipTypeOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationMembershipType(input)
	return &out, nil
}

type ScanningMode string

const (
	ScanningModeDefault ScanningMode = "Default"
)

func PossibleValuesForScanningMode() []string {
	return []string{
		string(ScanningModeDefault),
	}
}

func (s *ScanningMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScanningMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScanningMode(input string) (*ScanningMode, error) {
	vals := map[string]ScanningMode{
		"default": ScanningModeDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScanningMode(input)
	return &out, nil
}

type SubPlan string

const (
	SubPlanPOne SubPlan = "P1"
	SubPlanPTwo SubPlan = "P2"
)

func PossibleValuesForSubPlan() []string {
	return []string{
		string(SubPlanPOne),
		string(SubPlanPTwo),
	}
}

func (s *SubPlan) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubPlan(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubPlan(input string) (*SubPlan, error) {
	vals := map[string]SubPlan{
		"p1": SubPlanPOne,
		"p2": SubPlanPTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubPlan(input)
	return &out, nil
}

type Type string

const (
	TypeQualys Type = "Qualys"
	TypeTVM    Type = "TVM"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeQualys),
		string(TypeTVM),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"qualys": TypeQualys,
		"tvm":    TypeTVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
