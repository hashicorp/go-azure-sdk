package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftChangeRequestAssignedTo string

const (
	OpenShiftChangeRequestAssignedTo_Manager   OpenShiftChangeRequestAssignedTo = "manager"
	OpenShiftChangeRequestAssignedTo_Recipient OpenShiftChangeRequestAssignedTo = "recipient"
	OpenShiftChangeRequestAssignedTo_Sender    OpenShiftChangeRequestAssignedTo = "sender"
	OpenShiftChangeRequestAssignedTo_System    OpenShiftChangeRequestAssignedTo = "system"
)

func PossibleValuesForOpenShiftChangeRequestAssignedTo() []string {
	return []string{
		string(OpenShiftChangeRequestAssignedTo_Manager),
		string(OpenShiftChangeRequestAssignedTo_Recipient),
		string(OpenShiftChangeRequestAssignedTo_Sender),
		string(OpenShiftChangeRequestAssignedTo_System),
	}
}

func (s *OpenShiftChangeRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftChangeRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftChangeRequestAssignedTo(input string) (*OpenShiftChangeRequestAssignedTo, error) {
	vals := map[string]OpenShiftChangeRequestAssignedTo{
		"manager":   OpenShiftChangeRequestAssignedTo_Manager,
		"recipient": OpenShiftChangeRequestAssignedTo_Recipient,
		"sender":    OpenShiftChangeRequestAssignedTo_Sender,
		"system":    OpenShiftChangeRequestAssignedTo_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestAssignedTo(input)
	return &out, nil
}
