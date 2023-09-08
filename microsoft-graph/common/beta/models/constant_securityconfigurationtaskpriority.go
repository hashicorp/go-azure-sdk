package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskPriority string

const (
	SecurityConfigurationTaskPriorityhigh SecurityConfigurationTaskPriority = "High"
	SecurityConfigurationTaskPrioritylow  SecurityConfigurationTaskPriority = "Low"
	SecurityConfigurationTaskPrioritynone SecurityConfigurationTaskPriority = "None"
)

func PossibleValuesForSecurityConfigurationTaskPriority() []string {
	return []string{
		string(SecurityConfigurationTaskPriorityhigh),
		string(SecurityConfigurationTaskPrioritylow),
		string(SecurityConfigurationTaskPrioritynone),
	}
}

func parseSecurityConfigurationTaskPriority(input string) (*SecurityConfigurationTaskPriority, error) {
	vals := map[string]SecurityConfigurationTaskPriority{
		"high": SecurityConfigurationTaskPriorityhigh,
		"low":  SecurityConfigurationTaskPrioritylow,
		"none": SecurityConfigurationTaskPrioritynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskPriority(input)
	return &out, nil
}
