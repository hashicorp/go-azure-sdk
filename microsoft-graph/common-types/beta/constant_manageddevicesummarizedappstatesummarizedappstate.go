package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceSummarizedAppStateSummarizedAppState string

const (
	ManagedDeviceSummarizedAppStateSummarizedAppState_Fail          ManagedDeviceSummarizedAppStateSummarizedAppState = "fail"
	ManagedDeviceSummarizedAppStateSummarizedAppState_NotApplicable ManagedDeviceSummarizedAppStateSummarizedAppState = "notApplicable"
	ManagedDeviceSummarizedAppStateSummarizedAppState_Pending       ManagedDeviceSummarizedAppStateSummarizedAppState = "pending"
	ManagedDeviceSummarizedAppStateSummarizedAppState_ScriptError   ManagedDeviceSummarizedAppStateSummarizedAppState = "scriptError"
	ManagedDeviceSummarizedAppStateSummarizedAppState_Success       ManagedDeviceSummarizedAppStateSummarizedAppState = "success"
	ManagedDeviceSummarizedAppStateSummarizedAppState_Unknown       ManagedDeviceSummarizedAppStateSummarizedAppState = "unknown"
)

func PossibleValuesForManagedDeviceSummarizedAppStateSummarizedAppState() []string {
	return []string{
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_Fail),
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_NotApplicable),
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_Pending),
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_ScriptError),
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_Success),
		string(ManagedDeviceSummarizedAppStateSummarizedAppState_Unknown),
	}
}

func (s *ManagedDeviceSummarizedAppStateSummarizedAppState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceSummarizedAppStateSummarizedAppState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceSummarizedAppStateSummarizedAppState(input string) (*ManagedDeviceSummarizedAppStateSummarizedAppState, error) {
	vals := map[string]ManagedDeviceSummarizedAppStateSummarizedAppState{
		"fail":          ManagedDeviceSummarizedAppStateSummarizedAppState_Fail,
		"notapplicable": ManagedDeviceSummarizedAppStateSummarizedAppState_NotApplicable,
		"pending":       ManagedDeviceSummarizedAppStateSummarizedAppState_Pending,
		"scripterror":   ManagedDeviceSummarizedAppStateSummarizedAppState_ScriptError,
		"success":       ManagedDeviceSummarizedAppStateSummarizedAppState_Success,
		"unknown":       ManagedDeviceSummarizedAppStateSummarizedAppState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceSummarizedAppStateSummarizedAppState(input)
	return &out, nil
}
