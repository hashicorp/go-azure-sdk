package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10SecureAssessmentConfigurationConfigurationAccountType string

const (
	Windows10SecureAssessmentConfigurationConfigurationAccountTypeazureADAccount    Windows10SecureAssessmentConfigurationConfigurationAccountType = "AzureADAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountTypedomainAccount     Windows10SecureAssessmentConfigurationConfigurationAccountType = "DomainAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalAccount      Windows10SecureAssessmentConfigurationConfigurationAccountType = "LocalAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalGuestAccount Windows10SecureAssessmentConfigurationConfigurationAccountType = "LocalGuestAccount"
)

func PossibleValuesForWindows10SecureAssessmentConfigurationConfigurationAccountType() []string {
	return []string{
		string(Windows10SecureAssessmentConfigurationConfigurationAccountTypeazureADAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountTypedomainAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalGuestAccount),
	}
}

func parseWindows10SecureAssessmentConfigurationConfigurationAccountType(input string) (*Windows10SecureAssessmentConfigurationConfigurationAccountType, error) {
	vals := map[string]Windows10SecureAssessmentConfigurationConfigurationAccountType{
		"azureadaccount":    Windows10SecureAssessmentConfigurationConfigurationAccountTypeazureADAccount,
		"domainaccount":     Windows10SecureAssessmentConfigurationConfigurationAccountTypedomainAccount,
		"localaccount":      Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalAccount,
		"localguestaccount": Windows10SecureAssessmentConfigurationConfigurationAccountTypelocalGuestAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10SecureAssessmentConfigurationConfigurationAccountType(input)
	return &out, nil
}
