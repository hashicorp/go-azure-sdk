package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskStatus string

const (
	SecurityConfigurationTaskStatusactive    SecurityConfigurationTaskStatus = "Active"
	SecurityConfigurationTaskStatuscompleted SecurityConfigurationTaskStatus = "Completed"
	SecurityConfigurationTaskStatuspending   SecurityConfigurationTaskStatus = "Pending"
	SecurityConfigurationTaskStatusrejected  SecurityConfigurationTaskStatus = "Rejected"
	SecurityConfigurationTaskStatusunknown   SecurityConfigurationTaskStatus = "Unknown"
)

func PossibleValuesForSecurityConfigurationTaskStatus() []string {
	return []string{
		string(SecurityConfigurationTaskStatusactive),
		string(SecurityConfigurationTaskStatuscompleted),
		string(SecurityConfigurationTaskStatuspending),
		string(SecurityConfigurationTaskStatusrejected),
		string(SecurityConfigurationTaskStatusunknown),
	}
}

func parseSecurityConfigurationTaskStatus(input string) (*SecurityConfigurationTaskStatus, error) {
	vals := map[string]SecurityConfigurationTaskStatus{
		"active":    SecurityConfigurationTaskStatusactive,
		"completed": SecurityConfigurationTaskStatuscompleted,
		"pending":   SecurityConfigurationTaskStatuspending,
		"rejected":  SecurityConfigurationTaskStatusrejected,
		"unknown":   SecurityConfigurationTaskStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskStatus(input)
	return &out, nil
}
