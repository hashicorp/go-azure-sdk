package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessContainerAccessContainerType string

const (
	DelegatedAdminAccessContainerAccessContainerType_SecurityGroup DelegatedAdminAccessContainerAccessContainerType = "securityGroup"
)

func PossibleValuesForDelegatedAdminAccessContainerAccessContainerType() []string {
	return []string{
		string(DelegatedAdminAccessContainerAccessContainerType_SecurityGroup),
	}
}

func (s *DelegatedAdminAccessContainerAccessContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminAccessContainerAccessContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminAccessContainerAccessContainerType(input string) (*DelegatedAdminAccessContainerAccessContainerType, error) {
	vals := map[string]DelegatedAdminAccessContainerAccessContainerType{
		"securitygroup": DelegatedAdminAccessContainerAccessContainerType_SecurityGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminAccessContainerAccessContainerType(input)
	return &out, nil
}
