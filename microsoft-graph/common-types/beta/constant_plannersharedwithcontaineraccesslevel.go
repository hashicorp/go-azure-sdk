package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainerAccessLevel string

const (
	PlannerSharedWithContainerAccessLevel_FullAccess      PlannerSharedWithContainerAccessLevel = "fullAccess"
	PlannerSharedWithContainerAccessLevel_ReadAccess      PlannerSharedWithContainerAccessLevel = "readAccess"
	PlannerSharedWithContainerAccessLevel_ReadWriteAccess PlannerSharedWithContainerAccessLevel = "readWriteAccess"
)

func PossibleValuesForPlannerSharedWithContainerAccessLevel() []string {
	return []string{
		string(PlannerSharedWithContainerAccessLevel_FullAccess),
		string(PlannerSharedWithContainerAccessLevel_ReadAccess),
		string(PlannerSharedWithContainerAccessLevel_ReadWriteAccess),
	}
}

func (s *PlannerSharedWithContainerAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerSharedWithContainerAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerSharedWithContainerAccessLevel(input string) (*PlannerSharedWithContainerAccessLevel, error) {
	vals := map[string]PlannerSharedWithContainerAccessLevel{
		"fullaccess":      PlannerSharedWithContainerAccessLevel_FullAccess,
		"readaccess":      PlannerSharedWithContainerAccessLevel_ReadAccess,
		"readwriteaccess": PlannerSharedWithContainerAccessLevel_ReadWriteAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerAccessLevel(input)
	return &out, nil
}
