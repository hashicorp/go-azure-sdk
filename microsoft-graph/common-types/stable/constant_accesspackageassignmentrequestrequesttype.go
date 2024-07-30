package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequestType string

const (
	AccessPackageAssignmentRequestRequestType_AdminAdd     AccessPackageAssignmentRequestRequestType = "adminAdd"
	AccessPackageAssignmentRequestRequestType_AdminRemove  AccessPackageAssignmentRequestRequestType = "adminRemove"
	AccessPackageAssignmentRequestRequestType_AdminUpdate  AccessPackageAssignmentRequestRequestType = "adminUpdate"
	AccessPackageAssignmentRequestRequestType_NotSpecified AccessPackageAssignmentRequestRequestType = "notSpecified"
	AccessPackageAssignmentRequestRequestType_OnBehalfAdd  AccessPackageAssignmentRequestRequestType = "onBehalfAdd"
	AccessPackageAssignmentRequestRequestType_SystemAdd    AccessPackageAssignmentRequestRequestType = "systemAdd"
	AccessPackageAssignmentRequestRequestType_SystemRemove AccessPackageAssignmentRequestRequestType = "systemRemove"
	AccessPackageAssignmentRequestRequestType_SystemUpdate AccessPackageAssignmentRequestRequestType = "systemUpdate"
	AccessPackageAssignmentRequestRequestType_UserAdd      AccessPackageAssignmentRequestRequestType = "userAdd"
	AccessPackageAssignmentRequestRequestType_UserRemove   AccessPackageAssignmentRequestRequestType = "userRemove"
	AccessPackageAssignmentRequestRequestType_UserUpdate   AccessPackageAssignmentRequestRequestType = "userUpdate"
)

func PossibleValuesForAccessPackageAssignmentRequestRequestType() []string {
	return []string{
		string(AccessPackageAssignmentRequestRequestType_AdminAdd),
		string(AccessPackageAssignmentRequestRequestType_AdminRemove),
		string(AccessPackageAssignmentRequestRequestType_AdminUpdate),
		string(AccessPackageAssignmentRequestRequestType_NotSpecified),
		string(AccessPackageAssignmentRequestRequestType_OnBehalfAdd),
		string(AccessPackageAssignmentRequestRequestType_SystemAdd),
		string(AccessPackageAssignmentRequestRequestType_SystemRemove),
		string(AccessPackageAssignmentRequestRequestType_SystemUpdate),
		string(AccessPackageAssignmentRequestRequestType_UserAdd),
		string(AccessPackageAssignmentRequestRequestType_UserRemove),
		string(AccessPackageAssignmentRequestRequestType_UserUpdate),
	}
}

func (s *AccessPackageAssignmentRequestRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentRequestRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentRequestRequestType(input string) (*AccessPackageAssignmentRequestRequestType, error) {
	vals := map[string]AccessPackageAssignmentRequestRequestType{
		"adminadd":     AccessPackageAssignmentRequestRequestType_AdminAdd,
		"adminremove":  AccessPackageAssignmentRequestRequestType_AdminRemove,
		"adminupdate":  AccessPackageAssignmentRequestRequestType_AdminUpdate,
		"notspecified": AccessPackageAssignmentRequestRequestType_NotSpecified,
		"onbehalfadd":  AccessPackageAssignmentRequestRequestType_OnBehalfAdd,
		"systemadd":    AccessPackageAssignmentRequestRequestType_SystemAdd,
		"systemremove": AccessPackageAssignmentRequestRequestType_SystemRemove,
		"systemupdate": AccessPackageAssignmentRequestRequestType_SystemUpdate,
		"useradd":      AccessPackageAssignmentRequestRequestType_UserAdd,
		"userremove":   AccessPackageAssignmentRequestRequestType_UserRemove,
		"userupdate":   AccessPackageAssignmentRequestRequestType_UserUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestRequestType(input)
	return &out, nil
}
