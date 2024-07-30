package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionsConfigurationPolicySetItemErrorCode string

const (
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Deleted      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "deleted"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NoError      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "noError"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NotFound     EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "notFound"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Unauthorized EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForEnrollmentRestrictionsConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Deleted),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NoError),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NotFound),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *EnrollmentRestrictionsConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentRestrictionsConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentRestrictionsConfigurationPolicySetItemErrorCode(input string) (*EnrollmentRestrictionsConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]EnrollmentRestrictionsConfigurationPolicySetItemErrorCode{
		"deleted":      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": EnrollmentRestrictionsConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentRestrictionsConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
