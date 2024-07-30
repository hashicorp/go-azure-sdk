package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetErrorCode string

const (
	PolicySetErrorCode_Deleted      PolicySetErrorCode = "deleted"
	PolicySetErrorCode_NoError      PolicySetErrorCode = "noError"
	PolicySetErrorCode_NotFound     PolicySetErrorCode = "notFound"
	PolicySetErrorCode_Unauthorized PolicySetErrorCode = "unauthorized"
)

func PossibleValuesForPolicySetErrorCode() []string {
	return []string{
		string(PolicySetErrorCode_Deleted),
		string(PolicySetErrorCode_NoError),
		string(PolicySetErrorCode_NotFound),
		string(PolicySetErrorCode_Unauthorized),
	}
}

func (s *PolicySetErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicySetErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicySetErrorCode(input string) (*PolicySetErrorCode, error) {
	vals := map[string]PolicySetErrorCode{
		"deleted":      PolicySetErrorCode_Deleted,
		"noerror":      PolicySetErrorCode_NoError,
		"notfound":     PolicySetErrorCode_NotFound,
		"unauthorized": PolicySetErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetErrorCode(input)
	return &out, nil
}
