package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode string

const (
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Deleted      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "deleted"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NoError      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "noError"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NotFound     MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "notFound"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Unauthorized MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyPolicySetItemErrorCode() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Deleted),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NoError),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NotFound),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmWindowsInformationProtectionPolicyPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmWindowsInformationProtectionPolicyPolicySetItemErrorCode(input string) (*MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode{
		"deleted":      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Deleted,
		"noerror":      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NoError,
		"notfound":     MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_NotFound,
		"unauthorized": MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode(input)
	return &out, nil
}
