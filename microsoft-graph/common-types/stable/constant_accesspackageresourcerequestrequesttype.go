package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequestRequestType string

const (
	AccessPackageResourceRequestRequestType_AdminAdd     AccessPackageResourceRequestRequestType = "adminAdd"
	AccessPackageResourceRequestRequestType_AdminRemove  AccessPackageResourceRequestRequestType = "adminRemove"
	AccessPackageResourceRequestRequestType_AdminUpdate  AccessPackageResourceRequestRequestType = "adminUpdate"
	AccessPackageResourceRequestRequestType_NotSpecified AccessPackageResourceRequestRequestType = "notSpecified"
	AccessPackageResourceRequestRequestType_OnBehalfAdd  AccessPackageResourceRequestRequestType = "onBehalfAdd"
	AccessPackageResourceRequestRequestType_SystemAdd    AccessPackageResourceRequestRequestType = "systemAdd"
	AccessPackageResourceRequestRequestType_SystemRemove AccessPackageResourceRequestRequestType = "systemRemove"
	AccessPackageResourceRequestRequestType_SystemUpdate AccessPackageResourceRequestRequestType = "systemUpdate"
	AccessPackageResourceRequestRequestType_UserAdd      AccessPackageResourceRequestRequestType = "userAdd"
	AccessPackageResourceRequestRequestType_UserRemove   AccessPackageResourceRequestRequestType = "userRemove"
	AccessPackageResourceRequestRequestType_UserUpdate   AccessPackageResourceRequestRequestType = "userUpdate"
)

func PossibleValuesForAccessPackageResourceRequestRequestType() []string {
	return []string{
		string(AccessPackageResourceRequestRequestType_AdminAdd),
		string(AccessPackageResourceRequestRequestType_AdminRemove),
		string(AccessPackageResourceRequestRequestType_AdminUpdate),
		string(AccessPackageResourceRequestRequestType_NotSpecified),
		string(AccessPackageResourceRequestRequestType_OnBehalfAdd),
		string(AccessPackageResourceRequestRequestType_SystemAdd),
		string(AccessPackageResourceRequestRequestType_SystemRemove),
		string(AccessPackageResourceRequestRequestType_SystemUpdate),
		string(AccessPackageResourceRequestRequestType_UserAdd),
		string(AccessPackageResourceRequestRequestType_UserRemove),
		string(AccessPackageResourceRequestRequestType_UserUpdate),
	}
}

func (s *AccessPackageResourceRequestRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageResourceRequestRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageResourceRequestRequestType(input string) (*AccessPackageResourceRequestRequestType, error) {
	vals := map[string]AccessPackageResourceRequestRequestType{
		"adminadd":     AccessPackageResourceRequestRequestType_AdminAdd,
		"adminremove":  AccessPackageResourceRequestRequestType_AdminRemove,
		"adminupdate":  AccessPackageResourceRequestRequestType_AdminUpdate,
		"notspecified": AccessPackageResourceRequestRequestType_NotSpecified,
		"onbehalfadd":  AccessPackageResourceRequestRequestType_OnBehalfAdd,
		"systemadd":    AccessPackageResourceRequestRequestType_SystemAdd,
		"systemremove": AccessPackageResourceRequestRequestType_SystemRemove,
		"systemupdate": AccessPackageResourceRequestRequestType_SystemUpdate,
		"useradd":      AccessPackageResourceRequestRequestType_UserAdd,
		"userremove":   AccessPackageResourceRequestRequestType_UserRemove,
		"userupdate":   AccessPackageResourceRequestRequestType_UserUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageResourceRequestRequestType(input)
	return &out, nil
}
