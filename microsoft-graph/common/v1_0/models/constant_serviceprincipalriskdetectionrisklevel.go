package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskLevel string

const (
	ServicePrincipalRiskDetectionRiskLevelhidden ServicePrincipalRiskDetectionRiskLevel = "Hidden"
	ServicePrincipalRiskDetectionRiskLevelhigh   ServicePrincipalRiskDetectionRiskLevel = "High"
	ServicePrincipalRiskDetectionRiskLevellow    ServicePrincipalRiskDetectionRiskLevel = "Low"
	ServicePrincipalRiskDetectionRiskLevelmedium ServicePrincipalRiskDetectionRiskLevel = "Medium"
	ServicePrincipalRiskDetectionRiskLevelnone   ServicePrincipalRiskDetectionRiskLevel = "None"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskLevel() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskLevelhidden),
		string(ServicePrincipalRiskDetectionRiskLevelhigh),
		string(ServicePrincipalRiskDetectionRiskLevellow),
		string(ServicePrincipalRiskDetectionRiskLevelmedium),
		string(ServicePrincipalRiskDetectionRiskLevelnone),
	}
}

func parseServicePrincipalRiskDetectionRiskLevel(input string) (*ServicePrincipalRiskDetectionRiskLevel, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskLevel{
		"hidden": ServicePrincipalRiskDetectionRiskLevelhidden,
		"high":   ServicePrincipalRiskDetectionRiskLevelhigh,
		"low":    ServicePrincipalRiskDetectionRiskLevellow,
		"medium": ServicePrincipalRiskDetectionRiskLevelmedium,
		"none":   ServicePrincipalRiskDetectionRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskLevel(input)
	return &out, nil
}
