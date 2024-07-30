package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateDeviceStatusStatus string

const (
	IosUpdateDeviceStatusStatus_Compliant     IosUpdateDeviceStatusStatus = "compliant"
	IosUpdateDeviceStatusStatus_Conflict      IosUpdateDeviceStatusStatus = "conflict"
	IosUpdateDeviceStatusStatus_Error         IosUpdateDeviceStatusStatus = "error"
	IosUpdateDeviceStatusStatus_NonCompliant  IosUpdateDeviceStatusStatus = "nonCompliant"
	IosUpdateDeviceStatusStatus_NotApplicable IosUpdateDeviceStatusStatus = "notApplicable"
	IosUpdateDeviceStatusStatus_NotAssigned   IosUpdateDeviceStatusStatus = "notAssigned"
	IosUpdateDeviceStatusStatus_Remediated    IosUpdateDeviceStatusStatus = "remediated"
	IosUpdateDeviceStatusStatus_Unknown       IosUpdateDeviceStatusStatus = "unknown"
)

func PossibleValuesForIosUpdateDeviceStatusStatus() []string {
	return []string{
		string(IosUpdateDeviceStatusStatus_Compliant),
		string(IosUpdateDeviceStatusStatus_Conflict),
		string(IosUpdateDeviceStatusStatus_Error),
		string(IosUpdateDeviceStatusStatus_NonCompliant),
		string(IosUpdateDeviceStatusStatus_NotApplicable),
		string(IosUpdateDeviceStatusStatus_NotAssigned),
		string(IosUpdateDeviceStatusStatus_Remediated),
		string(IosUpdateDeviceStatusStatus_Unknown),
	}
}

func (s *IosUpdateDeviceStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosUpdateDeviceStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosUpdateDeviceStatusStatus(input string) (*IosUpdateDeviceStatusStatus, error) {
	vals := map[string]IosUpdateDeviceStatusStatus{
		"compliant":     IosUpdateDeviceStatusStatus_Compliant,
		"conflict":      IosUpdateDeviceStatusStatus_Conflict,
		"error":         IosUpdateDeviceStatusStatus_Error,
		"noncompliant":  IosUpdateDeviceStatusStatus_NonCompliant,
		"notapplicable": IosUpdateDeviceStatusStatus_NotApplicable,
		"notassigned":   IosUpdateDeviceStatusStatus_NotAssigned,
		"remediated":    IosUpdateDeviceStatusStatus_Remediated,
		"unknown":       IosUpdateDeviceStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateDeviceStatusStatus(input)
	return &out, nil
}
