package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryNoncustodialDataSourceStatus string

const (
	SecurityEdiscoveryNoncustodialDataSourceStatusactive   SecurityEdiscoveryNoncustodialDataSourceStatus = "Active"
	SecurityEdiscoveryNoncustodialDataSourceStatusreleased SecurityEdiscoveryNoncustodialDataSourceStatus = "Released"
)

func PossibleValuesForSecurityEdiscoveryNoncustodialDataSourceStatus() []string {
	return []string{
		string(SecurityEdiscoveryNoncustodialDataSourceStatusactive),
		string(SecurityEdiscoveryNoncustodialDataSourceStatusreleased),
	}
}

func parseSecurityEdiscoveryNoncustodialDataSourceStatus(input string) (*SecurityEdiscoveryNoncustodialDataSourceStatus, error) {
	vals := map[string]SecurityEdiscoveryNoncustodialDataSourceStatus{
		"active":   SecurityEdiscoveryNoncustodialDataSourceStatusactive,
		"released": SecurityEdiscoveryNoncustodialDataSourceStatusreleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryNoncustodialDataSourceStatus(input)
	return &out, nil
}
