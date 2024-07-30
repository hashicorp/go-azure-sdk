package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetItemStatus string

const (
	PolicySetItemStatus_Error          PolicySetItemStatus = "error"
	PolicySetItemStatus_NotAssigned    PolicySetItemStatus = "notAssigned"
	PolicySetItemStatus_PartialSuccess PolicySetItemStatus = "partialSuccess"
	PolicySetItemStatus_Success        PolicySetItemStatus = "success"
	PolicySetItemStatus_Unknown        PolicySetItemStatus = "unknown"
	PolicySetItemStatus_Validating     PolicySetItemStatus = "validating"
)

func PossibleValuesForPolicySetItemStatus() []string {
	return []string{
		string(PolicySetItemStatus_Error),
		string(PolicySetItemStatus_NotAssigned),
		string(PolicySetItemStatus_PartialSuccess),
		string(PolicySetItemStatus_Success),
		string(PolicySetItemStatus_Unknown),
		string(PolicySetItemStatus_Validating),
	}
}

func (s *PolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicySetItemStatus(input string) (*PolicySetItemStatus, error) {
	vals := map[string]PolicySetItemStatus{
		"error":          PolicySetItemStatus_Error,
		"notassigned":    PolicySetItemStatus_NotAssigned,
		"partialsuccess": PolicySetItemStatus_PartialSuccess,
		"success":        PolicySetItemStatus_Success,
		"unknown":        PolicySetItemStatus_Unknown,
		"validating":     PolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetItemStatus(input)
	return &out, nil
}
