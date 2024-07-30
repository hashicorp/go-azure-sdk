package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftWorkCloudPCAccessState string

const (
	ShiftWorkCloudPCAccessState_Activating          ShiftWorkCloudPCAccessState = "activating"
	ShiftWorkCloudPCAccessState_ActivationFailed    ShiftWorkCloudPCAccessState = "activationFailed"
	ShiftWorkCloudPCAccessState_Active              ShiftWorkCloudPCAccessState = "active"
	ShiftWorkCloudPCAccessState_NoLicensesAvailable ShiftWorkCloudPCAccessState = "noLicensesAvailable"
	ShiftWorkCloudPCAccessState_StandbyMode         ShiftWorkCloudPCAccessState = "standbyMode"
	ShiftWorkCloudPCAccessState_Unassigned          ShiftWorkCloudPCAccessState = "unassigned"
)

func PossibleValuesForShiftWorkCloudPCAccessState() []string {
	return []string{
		string(ShiftWorkCloudPCAccessState_Activating),
		string(ShiftWorkCloudPCAccessState_ActivationFailed),
		string(ShiftWorkCloudPCAccessState_Active),
		string(ShiftWorkCloudPCAccessState_NoLicensesAvailable),
		string(ShiftWorkCloudPCAccessState_StandbyMode),
		string(ShiftWorkCloudPCAccessState_Unassigned),
	}
}

func (s *ShiftWorkCloudPCAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShiftWorkCloudPCAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShiftWorkCloudPCAccessState(input string) (*ShiftWorkCloudPCAccessState, error) {
	vals := map[string]ShiftWorkCloudPCAccessState{
		"activating":          ShiftWorkCloudPCAccessState_Activating,
		"activationfailed":    ShiftWorkCloudPCAccessState_ActivationFailed,
		"active":              ShiftWorkCloudPCAccessState_Active,
		"nolicensesavailable": ShiftWorkCloudPCAccessState_NoLicensesAvailable,
		"standbymode":         ShiftWorkCloudPCAccessState_StandbyMode,
		"unassigned":          ShiftWorkCloudPCAccessState_Unassigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftWorkCloudPCAccessState(input)
	return &out, nil
}
