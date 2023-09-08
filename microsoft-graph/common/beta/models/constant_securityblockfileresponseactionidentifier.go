package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlockFileResponseActionIdentifier string

const (
	SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA1   SecurityBlockFileResponseActionIdentifier = "InitiatingProcessSHA1"
	SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA256 SecurityBlockFileResponseActionIdentifier = "InitiatingProcessSHA256"
	SecurityBlockFileResponseActionIdentifiersha1                    SecurityBlockFileResponseActionIdentifier = "Sha1"
	SecurityBlockFileResponseActionIdentifiersha256                  SecurityBlockFileResponseActionIdentifier = "Sha256"
)

func PossibleValuesForSecurityBlockFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA1),
		string(SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA256),
		string(SecurityBlockFileResponseActionIdentifiersha1),
		string(SecurityBlockFileResponseActionIdentifiersha256),
	}
}

func parseSecurityBlockFileResponseActionIdentifier(input string) (*SecurityBlockFileResponseActionIdentifier, error) {
	vals := map[string]SecurityBlockFileResponseActionIdentifier{
		"initiatingprocesssha1":   SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA1,
		"initiatingprocesssha256": SecurityBlockFileResponseActionIdentifierinitiatingProcessSHA256,
		"sha1":                    SecurityBlockFileResponseActionIdentifiersha1,
		"sha256":                  SecurityBlockFileResponseActionIdentifiersha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlockFileResponseActionIdentifier(input)
	return &out, nil
}
