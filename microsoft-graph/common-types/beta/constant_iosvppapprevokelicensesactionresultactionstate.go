package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResultActionState string

const (
	IosVppAppRevokeLicensesActionResultActionState_Active       IosVppAppRevokeLicensesActionResultActionState = "active"
	IosVppAppRevokeLicensesActionResultActionState_Canceled     IosVppAppRevokeLicensesActionResultActionState = "canceled"
	IosVppAppRevokeLicensesActionResultActionState_Done         IosVppAppRevokeLicensesActionResultActionState = "done"
	IosVppAppRevokeLicensesActionResultActionState_Failed       IosVppAppRevokeLicensesActionResultActionState = "failed"
	IosVppAppRevokeLicensesActionResultActionState_None         IosVppAppRevokeLicensesActionResultActionState = "none"
	IosVppAppRevokeLicensesActionResultActionState_NotSupported IosVppAppRevokeLicensesActionResultActionState = "notSupported"
	IosVppAppRevokeLicensesActionResultActionState_Pending      IosVppAppRevokeLicensesActionResultActionState = "pending"
)

func PossibleValuesForIosVppAppRevokeLicensesActionResultActionState() []string {
	return []string{
		string(IosVppAppRevokeLicensesActionResultActionState_Active),
		string(IosVppAppRevokeLicensesActionResultActionState_Canceled),
		string(IosVppAppRevokeLicensesActionResultActionState_Done),
		string(IosVppAppRevokeLicensesActionResultActionState_Failed),
		string(IosVppAppRevokeLicensesActionResultActionState_None),
		string(IosVppAppRevokeLicensesActionResultActionState_NotSupported),
		string(IosVppAppRevokeLicensesActionResultActionState_Pending),
	}
}

func (s *IosVppAppRevokeLicensesActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVppAppRevokeLicensesActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVppAppRevokeLicensesActionResultActionState(input string) (*IosVppAppRevokeLicensesActionResultActionState, error) {
	vals := map[string]IosVppAppRevokeLicensesActionResultActionState{
		"active":       IosVppAppRevokeLicensesActionResultActionState_Active,
		"canceled":     IosVppAppRevokeLicensesActionResultActionState_Canceled,
		"done":         IosVppAppRevokeLicensesActionResultActionState_Done,
		"failed":       IosVppAppRevokeLicensesActionResultActionState_Failed,
		"none":         IosVppAppRevokeLicensesActionResultActionState_None,
		"notsupported": IosVppAppRevokeLicensesActionResultActionState_NotSupported,
		"pending":      IosVppAppRevokeLicensesActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppRevokeLicensesActionResultActionState(input)
	return &out, nil
}
