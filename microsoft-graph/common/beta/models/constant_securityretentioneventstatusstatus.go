package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionEventStatusStatus string

const (
	SecurityRetentionEventStatusStatuserror        SecurityRetentionEventStatusStatus = "Error"
	SecurityRetentionEventStatusStatusnotAvaliable SecurityRetentionEventStatusStatus = "NotAvaliable"
	SecurityRetentionEventStatusStatuspending      SecurityRetentionEventStatusStatus = "Pending"
	SecurityRetentionEventStatusStatussuccess      SecurityRetentionEventStatusStatus = "Success"
)

func PossibleValuesForSecurityRetentionEventStatusStatus() []string {
	return []string{
		string(SecurityRetentionEventStatusStatuserror),
		string(SecurityRetentionEventStatusStatusnotAvaliable),
		string(SecurityRetentionEventStatusStatuspending),
		string(SecurityRetentionEventStatusStatussuccess),
	}
}

func parseSecurityRetentionEventStatusStatus(input string) (*SecurityRetentionEventStatusStatus, error) {
	vals := map[string]SecurityRetentionEventStatusStatus{
		"error":        SecurityRetentionEventStatusStatuserror,
		"notavaliable": SecurityRetentionEventStatusStatusnotAvaliable,
		"pending":      SecurityRetentionEventStatusStatuspending,
		"success":      SecurityRetentionEventStatusStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionEventStatusStatus(input)
	return &out, nil
}
