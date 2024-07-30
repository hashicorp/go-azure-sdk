package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressBookAccountTargetContentType string

const (
	AddressBookAccountTargetContentType_AddressBook AddressBookAccountTargetContentType = "addressBook"
	AddressBookAccountTargetContentType_IncludeAll  AddressBookAccountTargetContentType = "includeAll"
	AddressBookAccountTargetContentType_Unknown     AddressBookAccountTargetContentType = "unknown"
)

func PossibleValuesForAddressBookAccountTargetContentType() []string {
	return []string{
		string(AddressBookAccountTargetContentType_AddressBook),
		string(AddressBookAccountTargetContentType_IncludeAll),
		string(AddressBookAccountTargetContentType_Unknown),
	}
}

func (s *AddressBookAccountTargetContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddressBookAccountTargetContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddressBookAccountTargetContentType(input string) (*AddressBookAccountTargetContentType, error) {
	vals := map[string]AddressBookAccountTargetContentType{
		"addressbook": AddressBookAccountTargetContentType_AddressBook,
		"includeall":  AddressBookAccountTargetContentType_IncludeAll,
		"unknown":     AddressBookAccountTargetContentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddressBookAccountTargetContentType(input)
	return &out, nil
}
