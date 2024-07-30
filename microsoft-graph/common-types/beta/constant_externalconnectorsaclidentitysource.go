package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAclIdentitySource string

const (
	ExternalConnectorsAclIdentitySource_AzureActiveDirectory ExternalConnectorsAclIdentitySource = "azureActiveDirectory"
	ExternalConnectorsAclIdentitySource_External             ExternalConnectorsAclIdentitySource = "external"
)

func PossibleValuesForExternalConnectorsAclIdentitySource() []string {
	return []string{
		string(ExternalConnectorsAclIdentitySource_AzureActiveDirectory),
		string(ExternalConnectorsAclIdentitySource_External),
	}
}

func (s *ExternalConnectorsAclIdentitySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsAclIdentitySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsAclIdentitySource(input string) (*ExternalConnectorsAclIdentitySource, error) {
	vals := map[string]ExternalConnectorsAclIdentitySource{
		"azureactivedirectory": ExternalConnectorsAclIdentitySource_AzureActiveDirectory,
		"external":             ExternalConnectorsAclIdentitySource_External,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAclIdentitySource(input)
	return &out, nil
}
