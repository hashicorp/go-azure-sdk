package analysisdefinitionoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IssueType string

const (
	IssueTypeAppCrash             IssueType = "AppCrash"
	IssueTypeAppDeployment        IssueType = "AppDeployment"
	IssueTypeAseDeployment        IssueType = "AseDeployment"
	IssueTypeOther                IssueType = "Other"
	IssueTypePlatformIssue        IssueType = "PlatformIssue"
	IssueTypeRuntimeIssueDetected IssueType = "RuntimeIssueDetected"
	IssueTypeServiceIncident      IssueType = "ServiceIncident"
	IssueTypeUserIssue            IssueType = "UserIssue"
)

func PossibleValuesForIssueType() []string {
	return []string{
		string(IssueTypeAppCrash),
		string(IssueTypeAppDeployment),
		string(IssueTypeAseDeployment),
		string(IssueTypeOther),
		string(IssueTypePlatformIssue),
		string(IssueTypeRuntimeIssueDetected),
		string(IssueTypeServiceIncident),
		string(IssueTypeUserIssue),
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
		"appcrash":             IssueTypeAppCrash,
		"appdeployment":        IssueTypeAppDeployment,
		"asedeployment":        IssueTypeAseDeployment,
		"other":                IssueTypeOther,
		"platformissue":        IssueTypePlatformIssue,
		"runtimeissuedetected": IssueTypeRuntimeIssueDetected,
		"serviceincident":      IssueTypeServiceIncident,
		"userissue":            IssueTypeUserIssue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IssueType(input)
	return &out, nil
}

type SolutionType string

const (
	SolutionTypeBestPractices     SolutionType = "BestPractices"
	SolutionTypeDeepInvestigation SolutionType = "DeepInvestigation"
	SolutionTypeQuickSolution     SolutionType = "QuickSolution"
)

func PossibleValuesForSolutionType() []string {
	return []string{
		string(SolutionTypeBestPractices),
		string(SolutionTypeDeepInvestigation),
		string(SolutionTypeQuickSolution),
	}
}

func (s *SolutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSolutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSolutionType(input string) (*SolutionType, error) {
	vals := map[string]SolutionType{
		"bestpractices":     SolutionTypeBestPractices,
		"deepinvestigation": SolutionTypeDeepInvestigation,
		"quicksolution":     SolutionTypeQuickSolution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SolutionType(input)
	return &out, nil
}
