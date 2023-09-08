package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityStopAndQuarantineFileResponseActionIdentifier string

const (
	SecurityStopAndQuarantineFileResponseActionIdentifierdeviceId              SecurityStopAndQuarantineFileResponseActionIdentifier = "DeviceId"
	SecurityStopAndQuarantineFileResponseActionIdentifierinitiatingProcessSHA1 SecurityStopAndQuarantineFileResponseActionIdentifier = "InitiatingProcessSHA1"
	SecurityStopAndQuarantineFileResponseActionIdentifiersha1                  SecurityStopAndQuarantineFileResponseActionIdentifier = "Sha1"
)

func PossibleValuesForSecurityStopAndQuarantineFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityStopAndQuarantineFileResponseActionIdentifierdeviceId),
		string(SecurityStopAndQuarantineFileResponseActionIdentifierinitiatingProcessSHA1),
		string(SecurityStopAndQuarantineFileResponseActionIdentifiersha1),
	}
}

func parseSecurityStopAndQuarantineFileResponseActionIdentifier(input string) (*SecurityStopAndQuarantineFileResponseActionIdentifier, error) {
	vals := map[string]SecurityStopAndQuarantineFileResponseActionIdentifier{
		"deviceid":              SecurityStopAndQuarantineFileResponseActionIdentifierdeviceId,
		"initiatingprocesssha1": SecurityStopAndQuarantineFileResponseActionIdentifierinitiatingProcessSHA1,
		"sha1":                  SecurityStopAndQuarantineFileResponseActionIdentifiersha1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityStopAndQuarantineFileResponseActionIdentifier(input)
	return &out, nil
}
