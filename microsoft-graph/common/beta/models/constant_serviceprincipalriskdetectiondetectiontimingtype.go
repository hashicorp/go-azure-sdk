package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionDetectionTimingType string

const (
	ServicePrincipalRiskDetectionDetectionTimingTypenearRealtime ServicePrincipalRiskDetectionDetectionTimingType = "NearRealtime"
	ServicePrincipalRiskDetectionDetectionTimingTypenotDefined   ServicePrincipalRiskDetectionDetectionTimingType = "NotDefined"
	ServicePrincipalRiskDetectionDetectionTimingTypeoffline      ServicePrincipalRiskDetectionDetectionTimingType = "Offline"
	ServicePrincipalRiskDetectionDetectionTimingTyperealtime     ServicePrincipalRiskDetectionDetectionTimingType = "Realtime"
)

func PossibleValuesForServicePrincipalRiskDetectionDetectionTimingType() []string {
	return []string{
		string(ServicePrincipalRiskDetectionDetectionTimingTypenearRealtime),
		string(ServicePrincipalRiskDetectionDetectionTimingTypenotDefined),
		string(ServicePrincipalRiskDetectionDetectionTimingTypeoffline),
		string(ServicePrincipalRiskDetectionDetectionTimingTyperealtime),
	}
}

func parseServicePrincipalRiskDetectionDetectionTimingType(input string) (*ServicePrincipalRiskDetectionDetectionTimingType, error) {
	vals := map[string]ServicePrincipalRiskDetectionDetectionTimingType{
		"nearrealtime": ServicePrincipalRiskDetectionDetectionTimingTypenearRealtime,
		"notdefined":   ServicePrincipalRiskDetectionDetectionTimingTypenotDefined,
		"offline":      ServicePrincipalRiskDetectionDetectionTimingTypeoffline,
		"realtime":     ServicePrincipalRiskDetectionDetectionTimingTyperealtime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionDetectionTimingType(input)
	return &out, nil
}
