package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemErrorCode string

const (
	MobileAppPolicySetItemErrorCode_Deleted      MobileAppPolicySetItemErrorCode = "deleted"
	MobileAppPolicySetItemErrorCode_NoError      MobileAppPolicySetItemErrorCode = "noError"
	MobileAppPolicySetItemErrorCode_NotFound     MobileAppPolicySetItemErrorCode = "notFound"
	MobileAppPolicySetItemErrorCode_Unauthorized MobileAppPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForMobileAppPolicySetItemErrorCode() []string {
	return []string{
		string(MobileAppPolicySetItemErrorCode_Deleted),
		string(MobileAppPolicySetItemErrorCode_NoError),
		string(MobileAppPolicySetItemErrorCode_NotFound),
		string(MobileAppPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *MobileAppPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppPolicySetItemErrorCode(input string) (*MobileAppPolicySetItemErrorCode, error) {
	vals := map[string]MobileAppPolicySetItemErrorCode{
		"deleted":      MobileAppPolicySetItemErrorCode_Deleted,
		"noerror":      MobileAppPolicySetItemErrorCode_NoError,
		"notfound":     MobileAppPolicySetItemErrorCode_NotFound,
		"unauthorized": MobileAppPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemErrorCode(input)
	return &out, nil
}
