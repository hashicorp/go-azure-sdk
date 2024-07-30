package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionActivity string

const (
	ServicePrincipalRiskDetectionActivity_ServicePrincipal ServicePrincipalRiskDetectionActivity = "servicePrincipal"
	ServicePrincipalRiskDetectionActivity_Signin           ServicePrincipalRiskDetectionActivity = "signin"
	ServicePrincipalRiskDetectionActivity_User             ServicePrincipalRiskDetectionActivity = "user"
)

func PossibleValuesForServicePrincipalRiskDetectionActivity() []string {
	return []string{
		string(ServicePrincipalRiskDetectionActivity_ServicePrincipal),
		string(ServicePrincipalRiskDetectionActivity_Signin),
		string(ServicePrincipalRiskDetectionActivity_User),
	}
}

func (s *ServicePrincipalRiskDetectionActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionActivity(input string) (*ServicePrincipalRiskDetectionActivity, error) {
	vals := map[string]ServicePrincipalRiskDetectionActivity{
		"serviceprincipal": ServicePrincipalRiskDetectionActivity_ServicePrincipal,
		"signin":           ServicePrincipalRiskDetectionActivity_Signin,
		"user":             ServicePrincipalRiskDetectionActivity_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionActivity(input)
	return &out, nil
}
