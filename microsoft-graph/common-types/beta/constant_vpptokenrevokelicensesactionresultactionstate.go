package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenRevokeLicensesActionResultActionState string

const (
	VppTokenRevokeLicensesActionResultActionState_Active       VppTokenRevokeLicensesActionResultActionState = "active"
	VppTokenRevokeLicensesActionResultActionState_Canceled     VppTokenRevokeLicensesActionResultActionState = "canceled"
	VppTokenRevokeLicensesActionResultActionState_Done         VppTokenRevokeLicensesActionResultActionState = "done"
	VppTokenRevokeLicensesActionResultActionState_Failed       VppTokenRevokeLicensesActionResultActionState = "failed"
	VppTokenRevokeLicensesActionResultActionState_None         VppTokenRevokeLicensesActionResultActionState = "none"
	VppTokenRevokeLicensesActionResultActionState_NotSupported VppTokenRevokeLicensesActionResultActionState = "notSupported"
	VppTokenRevokeLicensesActionResultActionState_Pending      VppTokenRevokeLicensesActionResultActionState = "pending"
)

func PossibleValuesForVppTokenRevokeLicensesActionResultActionState() []string {
	return []string{
		string(VppTokenRevokeLicensesActionResultActionState_Active),
		string(VppTokenRevokeLicensesActionResultActionState_Canceled),
		string(VppTokenRevokeLicensesActionResultActionState_Done),
		string(VppTokenRevokeLicensesActionResultActionState_Failed),
		string(VppTokenRevokeLicensesActionResultActionState_None),
		string(VppTokenRevokeLicensesActionResultActionState_NotSupported),
		string(VppTokenRevokeLicensesActionResultActionState_Pending),
	}
}

func (s *VppTokenRevokeLicensesActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenRevokeLicensesActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenRevokeLicensesActionResultActionState(input string) (*VppTokenRevokeLicensesActionResultActionState, error) {
	vals := map[string]VppTokenRevokeLicensesActionResultActionState{
		"active":       VppTokenRevokeLicensesActionResultActionState_Active,
		"canceled":     VppTokenRevokeLicensesActionResultActionState_Canceled,
		"done":         VppTokenRevokeLicensesActionResultActionState_Done,
		"failed":       VppTokenRevokeLicensesActionResultActionState_Failed,
		"none":         VppTokenRevokeLicensesActionResultActionState_None,
		"notsupported": VppTokenRevokeLicensesActionResultActionState_NotSupported,
		"pending":      VppTokenRevokeLicensesActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenRevokeLicensesActionResultActionState(input)
	return &out, nil
}
