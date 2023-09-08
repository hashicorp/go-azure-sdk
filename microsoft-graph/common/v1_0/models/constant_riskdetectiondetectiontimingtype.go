package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionDetectionTimingType string

const (
	RiskDetectionDetectionTimingTypenearRealtime RiskDetectionDetectionTimingType = "NearRealtime"
	RiskDetectionDetectionTimingTypenotDefined   RiskDetectionDetectionTimingType = "NotDefined"
	RiskDetectionDetectionTimingTypeoffline      RiskDetectionDetectionTimingType = "Offline"
	RiskDetectionDetectionTimingTyperealtime     RiskDetectionDetectionTimingType = "Realtime"
)

func PossibleValuesForRiskDetectionDetectionTimingType() []string {
	return []string{
		string(RiskDetectionDetectionTimingTypenearRealtime),
		string(RiskDetectionDetectionTimingTypenotDefined),
		string(RiskDetectionDetectionTimingTypeoffline),
		string(RiskDetectionDetectionTimingTyperealtime),
	}
}

func parseRiskDetectionDetectionTimingType(input string) (*RiskDetectionDetectionTimingType, error) {
	vals := map[string]RiskDetectionDetectionTimingType{
		"nearrealtime": RiskDetectionDetectionTimingTypenearRealtime,
		"notdefined":   RiskDetectionDetectionTimingTypenotDefined,
		"offline":      RiskDetectionDetectionTimingTypeoffline,
		"realtime":     RiskDetectionDetectionTimingTyperealtime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionDetectionTimingType(input)
	return &out, nil
}
