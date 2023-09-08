package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTroubleshootingEventFailureCategory string

const (
	EnrollmentTroubleshootingEventFailureCategoryaccountValidation              EnrollmentTroubleshootingEventFailureCategory = "AccountValidation"
	EnrollmentTroubleshootingEventFailureCategoryauthentication                 EnrollmentTroubleshootingEventFailureCategory = "Authentication"
	EnrollmentTroubleshootingEventFailureCategoryauthorization                  EnrollmentTroubleshootingEventFailureCategory = "Authorization"
	EnrollmentTroubleshootingEventFailureCategorybadRequest                     EnrollmentTroubleshootingEventFailureCategory = "BadRequest"
	EnrollmentTroubleshootingEventFailureCategoryclientDisconnected             EnrollmentTroubleshootingEventFailureCategory = "ClientDisconnected"
	EnrollmentTroubleshootingEventFailureCategorydeviceNotSupported             EnrollmentTroubleshootingEventFailureCategory = "DeviceNotSupported"
	EnrollmentTroubleshootingEventFailureCategoryenrollmentRestrictionsEnforced EnrollmentTroubleshootingEventFailureCategory = "EnrollmentRestrictionsEnforced"
	EnrollmentTroubleshootingEventFailureCategoryfeatureNotSupported            EnrollmentTroubleshootingEventFailureCategory = "FeatureNotSupported"
	EnrollmentTroubleshootingEventFailureCategoryinMaintenance                  EnrollmentTroubleshootingEventFailureCategory = "InMaintenance"
	EnrollmentTroubleshootingEventFailureCategoryunknown                        EnrollmentTroubleshootingEventFailureCategory = "Unknown"
	EnrollmentTroubleshootingEventFailureCategoryuserAbandonment                EnrollmentTroubleshootingEventFailureCategory = "UserAbandonment"
	EnrollmentTroubleshootingEventFailureCategoryuserValidation                 EnrollmentTroubleshootingEventFailureCategory = "UserValidation"
)

func PossibleValuesForEnrollmentTroubleshootingEventFailureCategory() []string {
	return []string{
		string(EnrollmentTroubleshootingEventFailureCategoryaccountValidation),
		string(EnrollmentTroubleshootingEventFailureCategoryauthentication),
		string(EnrollmentTroubleshootingEventFailureCategoryauthorization),
		string(EnrollmentTroubleshootingEventFailureCategorybadRequest),
		string(EnrollmentTroubleshootingEventFailureCategoryclientDisconnected),
		string(EnrollmentTroubleshootingEventFailureCategorydeviceNotSupported),
		string(EnrollmentTroubleshootingEventFailureCategoryenrollmentRestrictionsEnforced),
		string(EnrollmentTroubleshootingEventFailureCategoryfeatureNotSupported),
		string(EnrollmentTroubleshootingEventFailureCategoryinMaintenance),
		string(EnrollmentTroubleshootingEventFailureCategoryunknown),
		string(EnrollmentTroubleshootingEventFailureCategoryuserAbandonment),
		string(EnrollmentTroubleshootingEventFailureCategoryuserValidation),
	}
}

func parseEnrollmentTroubleshootingEventFailureCategory(input string) (*EnrollmentTroubleshootingEventFailureCategory, error) {
	vals := map[string]EnrollmentTroubleshootingEventFailureCategory{
		"accountvalidation":              EnrollmentTroubleshootingEventFailureCategoryaccountValidation,
		"authentication":                 EnrollmentTroubleshootingEventFailureCategoryauthentication,
		"authorization":                  EnrollmentTroubleshootingEventFailureCategoryauthorization,
		"badrequest":                     EnrollmentTroubleshootingEventFailureCategorybadRequest,
		"clientdisconnected":             EnrollmentTroubleshootingEventFailureCategoryclientDisconnected,
		"devicenotsupported":             EnrollmentTroubleshootingEventFailureCategorydeviceNotSupported,
		"enrollmentrestrictionsenforced": EnrollmentTroubleshootingEventFailureCategoryenrollmentRestrictionsEnforced,
		"featurenotsupported":            EnrollmentTroubleshootingEventFailureCategoryfeatureNotSupported,
		"inmaintenance":                  EnrollmentTroubleshootingEventFailureCategoryinMaintenance,
		"unknown":                        EnrollmentTroubleshootingEventFailureCategoryunknown,
		"userabandonment":                EnrollmentTroubleshootingEventFailureCategoryuserAbandonment,
		"uservalidation":                 EnrollmentTroubleshootingEventFailureCategoryuserValidation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTroubleshootingEventFailureCategory(input)
	return &out, nil
}
