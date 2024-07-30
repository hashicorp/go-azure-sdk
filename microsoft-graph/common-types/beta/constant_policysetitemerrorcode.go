package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetItemErrorCode string

const (
	PolicySetItemErrorCode_Deleted      PolicySetItemErrorCode = "deleted"
	PolicySetItemErrorCode_NoError      PolicySetItemErrorCode = "noError"
	PolicySetItemErrorCode_NotFound     PolicySetItemErrorCode = "notFound"
	PolicySetItemErrorCode_Unauthorized PolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForPolicySetItemErrorCode() []string {
	return []string{
		string(PolicySetItemErrorCode_Deleted),
		string(PolicySetItemErrorCode_NoError),
		string(PolicySetItemErrorCode_NotFound),
		string(PolicySetItemErrorCode_Unauthorized),
	}
}

func (s *PolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicySetItemErrorCode(input string) (*PolicySetItemErrorCode, error) {
	vals := map[string]PolicySetItemErrorCode{
		"deleted":      PolicySetItemErrorCode_Deleted,
		"noerror":      PolicySetItemErrorCode_NoError,
		"notfound":     PolicySetItemErrorCode_NotFound,
		"unauthorized": PolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetItemErrorCode(input)
	return &out, nil
}
