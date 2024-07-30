package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectSubjectLifecycle string

const (
	AccessPackageSubjectSubjectLifecycle_Governed    AccessPackageSubjectSubjectLifecycle = "governed"
	AccessPackageSubjectSubjectLifecycle_NotDefined  AccessPackageSubjectSubjectLifecycle = "notDefined"
	AccessPackageSubjectSubjectLifecycle_NotGoverned AccessPackageSubjectSubjectLifecycle = "notGoverned"
)

func PossibleValuesForAccessPackageSubjectSubjectLifecycle() []string {
	return []string{
		string(AccessPackageSubjectSubjectLifecycle_Governed),
		string(AccessPackageSubjectSubjectLifecycle_NotDefined),
		string(AccessPackageSubjectSubjectLifecycle_NotGoverned),
	}
}

func (s *AccessPackageSubjectSubjectLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageSubjectSubjectLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageSubjectSubjectLifecycle(input string) (*AccessPackageSubjectSubjectLifecycle, error) {
	vals := map[string]AccessPackageSubjectSubjectLifecycle{
		"governed":    AccessPackageSubjectSubjectLifecycle_Governed,
		"notdefined":  AccessPackageSubjectSubjectLifecycle_NotDefined,
		"notgoverned": AccessPackageSubjectSubjectLifecycle_NotGoverned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectSubjectLifecycle(input)
	return &out, nil
}
