package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAclAccessType string

const (
	ExternalConnectorsAclAccessTypedeny  ExternalConnectorsAclAccessType = "Deny"
	ExternalConnectorsAclAccessTypegrant ExternalConnectorsAclAccessType = "Grant"
)

func PossibleValuesForExternalConnectorsAclAccessType() []string {
	return []string{
		string(ExternalConnectorsAclAccessTypedeny),
		string(ExternalConnectorsAclAccessTypegrant),
	}
}

func parseExternalConnectorsAclAccessType(input string) (*ExternalConnectorsAclAccessType, error) {
	vals := map[string]ExternalConnectorsAclAccessType{
		"deny":  ExternalConnectorsAclAccessTypedeny,
		"grant": ExternalConnectorsAclAccessTypegrant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAclAccessType(input)
	return &out, nil
}
