package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftChangeRequestState string

const (
	OpenShiftChangeRequestState_Approved OpenShiftChangeRequestState = "approved"
	OpenShiftChangeRequestState_Declined OpenShiftChangeRequestState = "declined"
	OpenShiftChangeRequestState_Pending  OpenShiftChangeRequestState = "pending"
)

func PossibleValuesForOpenShiftChangeRequestState() []string {
	return []string{
		string(OpenShiftChangeRequestState_Approved),
		string(OpenShiftChangeRequestState_Declined),
		string(OpenShiftChangeRequestState_Pending),
	}
}

func (s *OpenShiftChangeRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftChangeRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftChangeRequestState(input string) (*OpenShiftChangeRequestState, error) {
	vals := map[string]OpenShiftChangeRequestState{
		"approved": OpenShiftChangeRequestState_Approved,
		"declined": OpenShiftChangeRequestState_Declined,
		"pending":  OpenShiftChangeRequestState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestState(input)
	return &out, nil
}
