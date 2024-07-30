package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTroubleshootingEventFailureCategory string

const (
	EnrollmentTroubleshootingEventFailureCategory_AccountValidation              EnrollmentTroubleshootingEventFailureCategory = "accountValidation"
	EnrollmentTroubleshootingEventFailureCategory_Authentication                 EnrollmentTroubleshootingEventFailureCategory = "authentication"
	EnrollmentTroubleshootingEventFailureCategory_Authorization                  EnrollmentTroubleshootingEventFailureCategory = "authorization"
	EnrollmentTroubleshootingEventFailureCategory_BadRequest                     EnrollmentTroubleshootingEventFailureCategory = "badRequest"
	EnrollmentTroubleshootingEventFailureCategory_ClientDisconnected             EnrollmentTroubleshootingEventFailureCategory = "clientDisconnected"
	EnrollmentTroubleshootingEventFailureCategory_DeviceNotSupported             EnrollmentTroubleshootingEventFailureCategory = "deviceNotSupported"
	EnrollmentTroubleshootingEventFailureCategory_EnrollmentRestrictionsEnforced EnrollmentTroubleshootingEventFailureCategory = "enrollmentRestrictionsEnforced"
	EnrollmentTroubleshootingEventFailureCategory_FeatureNotSupported            EnrollmentTroubleshootingEventFailureCategory = "featureNotSupported"
	EnrollmentTroubleshootingEventFailureCategory_InMaintenance                  EnrollmentTroubleshootingEventFailureCategory = "inMaintenance"
	EnrollmentTroubleshootingEventFailureCategory_Unknown                        EnrollmentTroubleshootingEventFailureCategory = "unknown"
	EnrollmentTroubleshootingEventFailureCategory_UserAbandonment                EnrollmentTroubleshootingEventFailureCategory = "userAbandonment"
	EnrollmentTroubleshootingEventFailureCategory_UserValidation                 EnrollmentTroubleshootingEventFailureCategory = "userValidation"
)

func PossibleValuesForEnrollmentTroubleshootingEventFailureCategory() []string {
	return []string{
		string(EnrollmentTroubleshootingEventFailureCategory_AccountValidation),
		string(EnrollmentTroubleshootingEventFailureCategory_Authentication),
		string(EnrollmentTroubleshootingEventFailureCategory_Authorization),
		string(EnrollmentTroubleshootingEventFailureCategory_BadRequest),
		string(EnrollmentTroubleshootingEventFailureCategory_ClientDisconnected),
		string(EnrollmentTroubleshootingEventFailureCategory_DeviceNotSupported),
		string(EnrollmentTroubleshootingEventFailureCategory_EnrollmentRestrictionsEnforced),
		string(EnrollmentTroubleshootingEventFailureCategory_FeatureNotSupported),
		string(EnrollmentTroubleshootingEventFailureCategory_InMaintenance),
		string(EnrollmentTroubleshootingEventFailureCategory_Unknown),
		string(EnrollmentTroubleshootingEventFailureCategory_UserAbandonment),
		string(EnrollmentTroubleshootingEventFailureCategory_UserValidation),
	}
}

func (s *EnrollmentTroubleshootingEventFailureCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentTroubleshootingEventFailureCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentTroubleshootingEventFailureCategory(input string) (*EnrollmentTroubleshootingEventFailureCategory, error) {
	vals := map[string]EnrollmentTroubleshootingEventFailureCategory{
		"accountvalidation":              EnrollmentTroubleshootingEventFailureCategory_AccountValidation,
		"authentication":                 EnrollmentTroubleshootingEventFailureCategory_Authentication,
		"authorization":                  EnrollmentTroubleshootingEventFailureCategory_Authorization,
		"badrequest":                     EnrollmentTroubleshootingEventFailureCategory_BadRequest,
		"clientdisconnected":             EnrollmentTroubleshootingEventFailureCategory_ClientDisconnected,
		"devicenotsupported":             EnrollmentTroubleshootingEventFailureCategory_DeviceNotSupported,
		"enrollmentrestrictionsenforced": EnrollmentTroubleshootingEventFailureCategory_EnrollmentRestrictionsEnforced,
		"featurenotsupported":            EnrollmentTroubleshootingEventFailureCategory_FeatureNotSupported,
		"inmaintenance":                  EnrollmentTroubleshootingEventFailureCategory_InMaintenance,
		"unknown":                        EnrollmentTroubleshootingEventFailureCategory_Unknown,
		"userabandonment":                EnrollmentTroubleshootingEventFailureCategory_UserAbandonment,
		"uservalidation":                 EnrollmentTroubleshootingEventFailureCategory_UserValidation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTroubleshootingEventFailureCategory(input)
	return &out, nil
}
