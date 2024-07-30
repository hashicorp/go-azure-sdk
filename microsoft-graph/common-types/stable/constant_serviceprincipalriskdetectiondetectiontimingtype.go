package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionDetectionTimingType string

const (
	ServicePrincipalRiskDetectionDetectionTimingType_NearRealtime ServicePrincipalRiskDetectionDetectionTimingType = "nearRealtime"
	ServicePrincipalRiskDetectionDetectionTimingType_NotDefined   ServicePrincipalRiskDetectionDetectionTimingType = "notDefined"
	ServicePrincipalRiskDetectionDetectionTimingType_Offline      ServicePrincipalRiskDetectionDetectionTimingType = "offline"
	ServicePrincipalRiskDetectionDetectionTimingType_Realtime     ServicePrincipalRiskDetectionDetectionTimingType = "realtime"
)

func PossibleValuesForServicePrincipalRiskDetectionDetectionTimingType() []string {
	return []string{
		string(ServicePrincipalRiskDetectionDetectionTimingType_NearRealtime),
		string(ServicePrincipalRiskDetectionDetectionTimingType_NotDefined),
		string(ServicePrincipalRiskDetectionDetectionTimingType_Offline),
		string(ServicePrincipalRiskDetectionDetectionTimingType_Realtime),
	}
}

func (s *ServicePrincipalRiskDetectionDetectionTimingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionDetectionTimingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionDetectionTimingType(input string) (*ServicePrincipalRiskDetectionDetectionTimingType, error) {
	vals := map[string]ServicePrincipalRiskDetectionDetectionTimingType{
		"nearrealtime": ServicePrincipalRiskDetectionDetectionTimingType_NearRealtime,
		"notdefined":   ServicePrincipalRiskDetectionDetectionTimingType_NotDefined,
		"offline":      ServicePrincipalRiskDetectionDetectionTimingType_Offline,
		"realtime":     ServicePrincipalRiskDetectionDetectionTimingType_Realtime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionDetectionTimingType(input)
	return &out, nil
}
