package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyPolicySetItemStatus string

const (
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Error          MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "error"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_NotAssigned    MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "notAssigned"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_PartialSuccess MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "partialSuccess"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Success        MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "success"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Unknown        MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "unknown"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Validating     MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "validating"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyPolicySetItemStatus() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Error),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_NotAssigned),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_PartialSuccess),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Success),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Unknown),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Validating),
	}
}

func (s *MdmWindowsInformationProtectionPolicyPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmWindowsInformationProtectionPolicyPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmWindowsInformationProtectionPolicyPolicySetItemStatus(input string) (*MdmWindowsInformationProtectionPolicyPolicySetItemStatus, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyPolicySetItemStatus{
		"error":          MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Error,
		"notassigned":    MdmWindowsInformationProtectionPolicyPolicySetItemStatus_NotAssigned,
		"partialsuccess": MdmWindowsInformationProtectionPolicyPolicySetItemStatus_PartialSuccess,
		"success":        MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Success,
		"unknown":        MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Unknown,
		"validating":     MdmWindowsInformationProtectionPolicyPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyPolicySetItemStatus(input)
	return &out, nil
}
