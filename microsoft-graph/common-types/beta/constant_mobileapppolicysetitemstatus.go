package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemStatus string

const (
	MobileAppPolicySetItemStatus_Error          MobileAppPolicySetItemStatus = "error"
	MobileAppPolicySetItemStatus_NotAssigned    MobileAppPolicySetItemStatus = "notAssigned"
	MobileAppPolicySetItemStatus_PartialSuccess MobileAppPolicySetItemStatus = "partialSuccess"
	MobileAppPolicySetItemStatus_Success        MobileAppPolicySetItemStatus = "success"
	MobileAppPolicySetItemStatus_Unknown        MobileAppPolicySetItemStatus = "unknown"
	MobileAppPolicySetItemStatus_Validating     MobileAppPolicySetItemStatus = "validating"
)

func PossibleValuesForMobileAppPolicySetItemStatus() []string {
	return []string{
		string(MobileAppPolicySetItemStatus_Error),
		string(MobileAppPolicySetItemStatus_NotAssigned),
		string(MobileAppPolicySetItemStatus_PartialSuccess),
		string(MobileAppPolicySetItemStatus_Success),
		string(MobileAppPolicySetItemStatus_Unknown),
		string(MobileAppPolicySetItemStatus_Validating),
	}
}

func (s *MobileAppPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppPolicySetItemStatus(input string) (*MobileAppPolicySetItemStatus, error) {
	vals := map[string]MobileAppPolicySetItemStatus{
		"error":          MobileAppPolicySetItemStatus_Error,
		"notassigned":    MobileAppPolicySetItemStatus_NotAssigned,
		"partialsuccess": MobileAppPolicySetItemStatus_PartialSuccess,
		"success":        MobileAppPolicySetItemStatus_Success,
		"unknown":        MobileAppPolicySetItemStatus_Unknown,
		"validating":     MobileAppPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemStatus(input)
	return &out, nil
}
