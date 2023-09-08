package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVmMetadataCloudProvider string

const (
	SecurityVmMetadataCloudProviderazure   SecurityVmMetadataCloudProvider = "Azure"
	SecurityVmMetadataCloudProviderunknown SecurityVmMetadataCloudProvider = "Unknown"
)

func PossibleValuesForSecurityVmMetadataCloudProvider() []string {
	return []string{
		string(SecurityVmMetadataCloudProviderazure),
		string(SecurityVmMetadataCloudProviderunknown),
	}
}

func parseSecurityVmMetadataCloudProvider(input string) (*SecurityVmMetadataCloudProvider, error) {
	vals := map[string]SecurityVmMetadataCloudProvider{
		"azure":   SecurityVmMetadataCloudProviderazure,
		"unknown": SecurityVmMetadataCloudProviderunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityVmMetadataCloudProvider(input)
	return &out, nil
}
