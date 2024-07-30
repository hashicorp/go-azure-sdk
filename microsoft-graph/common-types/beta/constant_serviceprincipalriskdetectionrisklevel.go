package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskLevel string

const (
	ServicePrincipalRiskDetectionRiskLevel_Hidden ServicePrincipalRiskDetectionRiskLevel = "hidden"
	ServicePrincipalRiskDetectionRiskLevel_High   ServicePrincipalRiskDetectionRiskLevel = "high"
	ServicePrincipalRiskDetectionRiskLevel_Low    ServicePrincipalRiskDetectionRiskLevel = "low"
	ServicePrincipalRiskDetectionRiskLevel_Medium ServicePrincipalRiskDetectionRiskLevel = "medium"
	ServicePrincipalRiskDetectionRiskLevel_None   ServicePrincipalRiskDetectionRiskLevel = "none"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskLevel() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskLevel_Hidden),
		string(ServicePrincipalRiskDetectionRiskLevel_High),
		string(ServicePrincipalRiskDetectionRiskLevel_Low),
		string(ServicePrincipalRiskDetectionRiskLevel_Medium),
		string(ServicePrincipalRiskDetectionRiskLevel_None),
	}
}

func (s *ServicePrincipalRiskDetectionRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionRiskLevel(input string) (*ServicePrincipalRiskDetectionRiskLevel, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskLevel{
		"hidden": ServicePrincipalRiskDetectionRiskLevel_Hidden,
		"high":   ServicePrincipalRiskDetectionRiskLevel_High,
		"low":    ServicePrincipalRiskDetectionRiskLevel_Low,
		"medium": ServicePrincipalRiskDetectionRiskLevel_Medium,
		"none":   ServicePrincipalRiskDetectionRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskLevel(input)
	return &out, nil
}
