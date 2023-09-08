package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAllowFileResponseActionIdentifier string

const (
	SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA1   SecurityAllowFileResponseActionIdentifier = "InitiatingProcessSHA1"
	SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA256 SecurityAllowFileResponseActionIdentifier = "InitiatingProcessSHA256"
	SecurityAllowFileResponseActionIdentifiersha1                    SecurityAllowFileResponseActionIdentifier = "Sha1"
	SecurityAllowFileResponseActionIdentifiersha256                  SecurityAllowFileResponseActionIdentifier = "Sha256"
)

func PossibleValuesForSecurityAllowFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA1),
		string(SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA256),
		string(SecurityAllowFileResponseActionIdentifiersha1),
		string(SecurityAllowFileResponseActionIdentifiersha256),
	}
}

func parseSecurityAllowFileResponseActionIdentifier(input string) (*SecurityAllowFileResponseActionIdentifier, error) {
	vals := map[string]SecurityAllowFileResponseActionIdentifier{
		"initiatingprocesssha1":   SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA1,
		"initiatingprocesssha256": SecurityAllowFileResponseActionIdentifierinitiatingProcessSHA256,
		"sha1":                    SecurityAllowFileResponseActionIdentifiersha1,
		"sha256":                  SecurityAllowFileResponseActionIdentifiersha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAllowFileResponseActionIdentifier(input)
	return &out, nil
}
