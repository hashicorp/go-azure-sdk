package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclIdentitySource string

const (
	AclIdentitySource_AzureActiveDirectory AclIdentitySource = "azureActiveDirectory"
	AclIdentitySource_External             AclIdentitySource = "external"
)

func PossibleValuesForAclIdentitySource() []string {
	return []string{
		string(AclIdentitySource_AzureActiveDirectory),
		string(AclIdentitySource_External),
	}
}

func (s *AclIdentitySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAclIdentitySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAclIdentitySource(input string) (*AclIdentitySource, error) {
	vals := map[string]AclIdentitySource{
		"azureactivedirectory": AclIdentitySource_AzureActiveDirectory,
		"external":             AclIdentitySource_External,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclIdentitySource(input)
	return &out, nil
}
