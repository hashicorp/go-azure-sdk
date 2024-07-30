package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionsConfigurationPolicySetItemStatus string

const (
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_Error          EnrollmentRestrictionsConfigurationPolicySetItemStatus = "error"
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_NotAssigned    EnrollmentRestrictionsConfigurationPolicySetItemStatus = "notAssigned"
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_PartialSuccess EnrollmentRestrictionsConfigurationPolicySetItemStatus = "partialSuccess"
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_Success        EnrollmentRestrictionsConfigurationPolicySetItemStatus = "success"
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_Unknown        EnrollmentRestrictionsConfigurationPolicySetItemStatus = "unknown"
	EnrollmentRestrictionsConfigurationPolicySetItemStatus_Validating     EnrollmentRestrictionsConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForEnrollmentRestrictionsConfigurationPolicySetItemStatus() []string {
	return []string{
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_Error),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_NotAssigned),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_PartialSuccess),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_Success),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_Unknown),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *EnrollmentRestrictionsConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentRestrictionsConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentRestrictionsConfigurationPolicySetItemStatus(input string) (*EnrollmentRestrictionsConfigurationPolicySetItemStatus, error) {
	vals := map[string]EnrollmentRestrictionsConfigurationPolicySetItemStatus{
		"error":          EnrollmentRestrictionsConfigurationPolicySetItemStatus_Error,
		"notassigned":    EnrollmentRestrictionsConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": EnrollmentRestrictionsConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        EnrollmentRestrictionsConfigurationPolicySetItemStatus_Success,
		"unknown":        EnrollmentRestrictionsConfigurationPolicySetItemStatus_Unknown,
		"validating":     EnrollmentRestrictionsConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentRestrictionsConfigurationPolicySetItemStatus(input)
	return &out, nil
}
