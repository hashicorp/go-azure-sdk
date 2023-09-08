package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskCategory string

const (
	SecurityConfigurationTaskCategoryadvancedThreatProtection SecurityConfigurationTaskCategory = "AdvancedThreatProtection"
	SecurityConfigurationTaskCategoryunknown                  SecurityConfigurationTaskCategory = "Unknown"
)

func PossibleValuesForSecurityConfigurationTaskCategory() []string {
	return []string{
		string(SecurityConfigurationTaskCategoryadvancedThreatProtection),
		string(SecurityConfigurationTaskCategoryunknown),
	}
}

func parseSecurityConfigurationTaskCategory(input string) (*SecurityConfigurationTaskCategory, error) {
	vals := map[string]SecurityConfigurationTaskCategory{
		"advancedthreatprotection": SecurityConfigurationTaskCategoryadvancedThreatProtection,
		"unknown":                  SecurityConfigurationTaskCategoryunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskCategory(input)
	return &out, nil
}
