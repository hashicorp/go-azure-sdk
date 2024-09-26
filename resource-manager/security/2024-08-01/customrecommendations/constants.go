package customrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationSupportedClouds string

const (
	RecommendationSupportedCloudsAWS   RecommendationSupportedClouds = "AWS"
	RecommendationSupportedCloudsAzure RecommendationSupportedClouds = "Azure"
	RecommendationSupportedCloudsGCP   RecommendationSupportedClouds = "GCP"
)

func PossibleValuesForRecommendationSupportedClouds() []string {
	return []string{
		string(RecommendationSupportedCloudsAWS),
		string(RecommendationSupportedCloudsAzure),
		string(RecommendationSupportedCloudsGCP),
	}
}

func (s *RecommendationSupportedClouds) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationSupportedClouds(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationSupportedClouds(input string) (*RecommendationSupportedClouds, error) {
	vals := map[string]RecommendationSupportedClouds{
		"aws":   RecommendationSupportedCloudsAWS,
		"azure": RecommendationSupportedCloudsAzure,
		"gcp":   RecommendationSupportedCloudsGCP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationSupportedClouds(input)
	return &out, nil
}

type SecurityIssue string

const (
	SecurityIssueAnonymousAccess      SecurityIssue = "AnonymousAccess"
	SecurityIssueBestPractices        SecurityIssue = "BestPractices"
	SecurityIssueExcessivePermissions SecurityIssue = "ExcessivePermissions"
	SecurityIssueNetworkExposure      SecurityIssue = "NetworkExposure"
	SecurityIssueTrafficEncryption    SecurityIssue = "TrafficEncryption"
	SecurityIssueVulnerability        SecurityIssue = "Vulnerability"
)

func PossibleValuesForSecurityIssue() []string {
	return []string{
		string(SecurityIssueAnonymousAccess),
		string(SecurityIssueBestPractices),
		string(SecurityIssueExcessivePermissions),
		string(SecurityIssueNetworkExposure),
		string(SecurityIssueTrafficEncryption),
		string(SecurityIssueVulnerability),
	}
}

func (s *SecurityIssue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIssue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIssue(input string) (*SecurityIssue, error) {
	vals := map[string]SecurityIssue{
		"anonymousaccess":      SecurityIssueAnonymousAccess,
		"bestpractices":        SecurityIssueBestPractices,
		"excessivepermissions": SecurityIssueExcessivePermissions,
		"networkexposure":      SecurityIssueNetworkExposure,
		"trafficencryption":    SecurityIssueTrafficEncryption,
		"vulnerability":        SecurityIssueVulnerability,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIssue(input)
	return &out, nil
}

type SeverityEnum string

const (
	SeverityEnumHigh   SeverityEnum = "High"
	SeverityEnumLow    SeverityEnum = "Low"
	SeverityEnumMedium SeverityEnum = "Medium"
)

func PossibleValuesForSeverityEnum() []string {
	return []string{
		string(SeverityEnumHigh),
		string(SeverityEnumLow),
		string(SeverityEnumMedium),
	}
}

func (s *SeverityEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeverityEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeverityEnum(input string) (*SeverityEnum, error) {
	vals := map[string]SeverityEnum{
		"high":   SeverityEnumHigh,
		"low":    SeverityEnumLow,
		"medium": SeverityEnumMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SeverityEnum(input)
	return &out, nil
}
