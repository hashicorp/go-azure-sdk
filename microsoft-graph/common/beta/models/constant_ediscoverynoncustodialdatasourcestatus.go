package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryNoncustodialDataSourceStatus string

const (
	EdiscoveryNoncustodialDataSourceStatusActive   EdiscoveryNoncustodialDataSourceStatus = "Active"
	EdiscoveryNoncustodialDataSourceStatusReleased EdiscoveryNoncustodialDataSourceStatus = "Released"
)

func PossibleValuesForEdiscoveryNoncustodialDataSourceStatus() []string {
	return []string{
		string(EdiscoveryNoncustodialDataSourceStatusActive),
		string(EdiscoveryNoncustodialDataSourceStatusReleased),
	}
}

func parseEdiscoveryNoncustodialDataSourceStatus(input string) (*EdiscoveryNoncustodialDataSourceStatus, error) {
	vals := map[string]EdiscoveryNoncustodialDataSourceStatus{
		"active":   EdiscoveryNoncustodialDataSourceStatusActive,
		"released": EdiscoveryNoncustodialDataSourceStatusReleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryNoncustodialDataSourceStatus(input)
	return &out, nil
}
