package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationEncryptionProtocol string

const (
	WorkforceIntegrationEncryptionProtocolsharedSecret WorkforceIntegrationEncryptionProtocol = "SharedSecret"
)

func PossibleValuesForWorkforceIntegrationEncryptionProtocol() []string {
	return []string{
		string(WorkforceIntegrationEncryptionProtocolsharedSecret),
	}
}

func parseWorkforceIntegrationEncryptionProtocol(input string) (*WorkforceIntegrationEncryptionProtocol, error) {
	vals := map[string]WorkforceIntegrationEncryptionProtocol{
		"sharedsecret": WorkforceIntegrationEncryptionProtocolsharedSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationEncryptionProtocol(input)
	return &out, nil
}
