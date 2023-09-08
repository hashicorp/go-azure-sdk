package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionActivity string

const (
	ServicePrincipalRiskDetectionActivityservicePrincipal ServicePrincipalRiskDetectionActivity = "ServicePrincipal"
	ServicePrincipalRiskDetectionActivitysignin           ServicePrincipalRiskDetectionActivity = "Signin"
	ServicePrincipalRiskDetectionActivityuser             ServicePrincipalRiskDetectionActivity = "User"
)

func PossibleValuesForServicePrincipalRiskDetectionActivity() []string {
	return []string{
		string(ServicePrincipalRiskDetectionActivityservicePrincipal),
		string(ServicePrincipalRiskDetectionActivitysignin),
		string(ServicePrincipalRiskDetectionActivityuser),
	}
}

func parseServicePrincipalRiskDetectionActivity(input string) (*ServicePrincipalRiskDetectionActivity, error) {
	vals := map[string]ServicePrincipalRiskDetectionActivity{
		"serviceprincipal": ServicePrincipalRiskDetectionActivityservicePrincipal,
		"signin":           ServicePrincipalRiskDetectionActivitysignin,
		"user":             ServicePrincipalRiskDetectionActivityuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionActivity(input)
	return &out, nil
}
