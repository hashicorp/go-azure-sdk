package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAclIdentitySource string

const (
	ExternalConnectorsAclIdentitySourceazureActiveDirectory ExternalConnectorsAclIdentitySource = "AzureActiveDirectory"
	ExternalConnectorsAclIdentitySourceexternal             ExternalConnectorsAclIdentitySource = "External"
)

func PossibleValuesForExternalConnectorsAclIdentitySource() []string {
	return []string{
		string(ExternalConnectorsAclIdentitySourceazureActiveDirectory),
		string(ExternalConnectorsAclIdentitySourceexternal),
	}
}

func parseExternalConnectorsAclIdentitySource(input string) (*ExternalConnectorsAclIdentitySource, error) {
	vals := map[string]ExternalConnectorsAclIdentitySource{
		"azureactivedirectory": ExternalConnectorsAclIdentitySourceazureActiveDirectory,
		"external":             ExternalConnectorsAclIdentitySourceexternal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAclIdentitySource(input)
	return &out, nil
}
