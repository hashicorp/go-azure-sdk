package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppRevokeLicensesActionResultActionState string

const (
	MacOsVppAppRevokeLicensesActionResultActionState_Active       MacOsVppAppRevokeLicensesActionResultActionState = "active"
	MacOsVppAppRevokeLicensesActionResultActionState_Canceled     MacOsVppAppRevokeLicensesActionResultActionState = "canceled"
	MacOsVppAppRevokeLicensesActionResultActionState_Done         MacOsVppAppRevokeLicensesActionResultActionState = "done"
	MacOsVppAppRevokeLicensesActionResultActionState_Failed       MacOsVppAppRevokeLicensesActionResultActionState = "failed"
	MacOsVppAppRevokeLicensesActionResultActionState_None         MacOsVppAppRevokeLicensesActionResultActionState = "none"
	MacOsVppAppRevokeLicensesActionResultActionState_NotSupported MacOsVppAppRevokeLicensesActionResultActionState = "notSupported"
	MacOsVppAppRevokeLicensesActionResultActionState_Pending      MacOsVppAppRevokeLicensesActionResultActionState = "pending"
)

func PossibleValuesForMacOsVppAppRevokeLicensesActionResultActionState() []string {
	return []string{
		string(MacOsVppAppRevokeLicensesActionResultActionState_Active),
		string(MacOsVppAppRevokeLicensesActionResultActionState_Canceled),
		string(MacOsVppAppRevokeLicensesActionResultActionState_Done),
		string(MacOsVppAppRevokeLicensesActionResultActionState_Failed),
		string(MacOsVppAppRevokeLicensesActionResultActionState_None),
		string(MacOsVppAppRevokeLicensesActionResultActionState_NotSupported),
		string(MacOsVppAppRevokeLicensesActionResultActionState_Pending),
	}
}

func (s *MacOsVppAppRevokeLicensesActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOsVppAppRevokeLicensesActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOsVppAppRevokeLicensesActionResultActionState(input string) (*MacOsVppAppRevokeLicensesActionResultActionState, error) {
	vals := map[string]MacOsVppAppRevokeLicensesActionResultActionState{
		"active":       MacOsVppAppRevokeLicensesActionResultActionState_Active,
		"canceled":     MacOsVppAppRevokeLicensesActionResultActionState_Canceled,
		"done":         MacOsVppAppRevokeLicensesActionResultActionState_Done,
		"failed":       MacOsVppAppRevokeLicensesActionResultActionState_Failed,
		"none":         MacOsVppAppRevokeLicensesActionResultActionState_None,
		"notsupported": MacOsVppAppRevokeLicensesActionResultActionState_NotSupported,
		"pending":      MacOsVppAppRevokeLicensesActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppRevokeLicensesActionResultActionState(input)
	return &out, nil
}
