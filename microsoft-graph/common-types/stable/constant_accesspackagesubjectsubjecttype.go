package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectSubjectType string

const (
	AccessPackageSubjectSubjectType_NotSpecified     AccessPackageSubjectSubjectType = "notSpecified"
	AccessPackageSubjectSubjectType_ServicePrincipal AccessPackageSubjectSubjectType = "servicePrincipal"
	AccessPackageSubjectSubjectType_User             AccessPackageSubjectSubjectType = "user"
)

func PossibleValuesForAccessPackageSubjectSubjectType() []string {
	return []string{
		string(AccessPackageSubjectSubjectType_NotSpecified),
		string(AccessPackageSubjectSubjectType_ServicePrincipal),
		string(AccessPackageSubjectSubjectType_User),
	}
}

func (s *AccessPackageSubjectSubjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageSubjectSubjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageSubjectSubjectType(input string) (*AccessPackageSubjectSubjectType, error) {
	vals := map[string]AccessPackageSubjectSubjectType{
		"notspecified":     AccessPackageSubjectSubjectType_NotSpecified,
		"serviceprincipal": AccessPackageSubjectSubjectType_ServicePrincipal,
		"user":             AccessPackageSubjectSubjectType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectSubjectType(input)
	return &out, nil
}
