package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeAppleVppLicensesActionResultActionState string

const (
	RevokeAppleVppLicensesActionResultActionState_Active       RevokeAppleVppLicensesActionResultActionState = "active"
	RevokeAppleVppLicensesActionResultActionState_Canceled     RevokeAppleVppLicensesActionResultActionState = "canceled"
	RevokeAppleVppLicensesActionResultActionState_Done         RevokeAppleVppLicensesActionResultActionState = "done"
	RevokeAppleVppLicensesActionResultActionState_Failed       RevokeAppleVppLicensesActionResultActionState = "failed"
	RevokeAppleVppLicensesActionResultActionState_None         RevokeAppleVppLicensesActionResultActionState = "none"
	RevokeAppleVppLicensesActionResultActionState_NotSupported RevokeAppleVppLicensesActionResultActionState = "notSupported"
	RevokeAppleVppLicensesActionResultActionState_Pending      RevokeAppleVppLicensesActionResultActionState = "pending"
)

func PossibleValuesForRevokeAppleVppLicensesActionResultActionState() []string {
	return []string{
		string(RevokeAppleVppLicensesActionResultActionState_Active),
		string(RevokeAppleVppLicensesActionResultActionState_Canceled),
		string(RevokeAppleVppLicensesActionResultActionState_Done),
		string(RevokeAppleVppLicensesActionResultActionState_Failed),
		string(RevokeAppleVppLicensesActionResultActionState_None),
		string(RevokeAppleVppLicensesActionResultActionState_NotSupported),
		string(RevokeAppleVppLicensesActionResultActionState_Pending),
	}
}

func (s *RevokeAppleVppLicensesActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRevokeAppleVppLicensesActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRevokeAppleVppLicensesActionResultActionState(input string) (*RevokeAppleVppLicensesActionResultActionState, error) {
	vals := map[string]RevokeAppleVppLicensesActionResultActionState{
		"active":       RevokeAppleVppLicensesActionResultActionState_Active,
		"canceled":     RevokeAppleVppLicensesActionResultActionState_Canceled,
		"done":         RevokeAppleVppLicensesActionResultActionState_Done,
		"failed":       RevokeAppleVppLicensesActionResultActionState_Failed,
		"none":         RevokeAppleVppLicensesActionResultActionState_None,
		"notsupported": RevokeAppleVppLicensesActionResultActionState_NotSupported,
		"pending":      RevokeAppleVppLicensesActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RevokeAppleVppLicensesActionResultActionState(input)
	return &out, nil
}
