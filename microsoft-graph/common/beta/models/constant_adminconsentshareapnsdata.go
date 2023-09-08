package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentShareAPNSData string

const (
	AdminConsentShareAPNSDatagranted       AdminConsentShareAPNSData = "Granted"
	AdminConsentShareAPNSDatanotConfigured AdminConsentShareAPNSData = "NotConfigured"
	AdminConsentShareAPNSDatanotGranted    AdminConsentShareAPNSData = "NotGranted"
)

func PossibleValuesForAdminConsentShareAPNSData() []string {
	return []string{
		string(AdminConsentShareAPNSDatagranted),
		string(AdminConsentShareAPNSDatanotConfigured),
		string(AdminConsentShareAPNSDatanotGranted),
	}
}

func parseAdminConsentShareAPNSData(input string) (*AdminConsentShareAPNSData, error) {
	vals := map[string]AdminConsentShareAPNSData{
		"granted":       AdminConsentShareAPNSDatagranted,
		"notconfigured": AdminConsentShareAPNSDatanotConfigured,
		"notgranted":    AdminConsentShareAPNSDatanotGranted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdminConsentShareAPNSData(input)
	return &out, nil
}
