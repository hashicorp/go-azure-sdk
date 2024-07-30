package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10SecureAssessmentConfigurationConfigurationAccountType string

const (
	Windows10SecureAssessmentConfigurationConfigurationAccountType_AzureADAccount    Windows10SecureAssessmentConfigurationConfigurationAccountType = "azureADAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountType_DomainAccount     Windows10SecureAssessmentConfigurationConfigurationAccountType = "domainAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalAccount      Windows10SecureAssessmentConfigurationConfigurationAccountType = "localAccount"
	Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalGuestAccount Windows10SecureAssessmentConfigurationConfigurationAccountType = "localGuestAccount"
)

func PossibleValuesForWindows10SecureAssessmentConfigurationConfigurationAccountType() []string {
	return []string{
		string(Windows10SecureAssessmentConfigurationConfigurationAccountType_AzureADAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountType_DomainAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalAccount),
		string(Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalGuestAccount),
	}
}

func (s *Windows10SecureAssessmentConfigurationConfigurationAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10SecureAssessmentConfigurationConfigurationAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10SecureAssessmentConfigurationConfigurationAccountType(input string) (*Windows10SecureAssessmentConfigurationConfigurationAccountType, error) {
	vals := map[string]Windows10SecureAssessmentConfigurationConfigurationAccountType{
		"azureadaccount":    Windows10SecureAssessmentConfigurationConfigurationAccountType_AzureADAccount,
		"domainaccount":     Windows10SecureAssessmentConfigurationConfigurationAccountType_DomainAccount,
		"localaccount":      Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalAccount,
		"localguestaccount": Windows10SecureAssessmentConfigurationConfigurationAccountType_LocalGuestAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10SecureAssessmentConfigurationConfigurationAccountType(input)
	return &out, nil
}
