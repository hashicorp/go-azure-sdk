package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressBookAccountTargetContentType string

const (
	AddressBookAccountTargetContentTypeaddressBook AddressBookAccountTargetContentType = "AddressBook"
	AddressBookAccountTargetContentTypeincludeAll  AddressBookAccountTargetContentType = "IncludeAll"
	AddressBookAccountTargetContentTypeunknown     AddressBookAccountTargetContentType = "Unknown"
)

func PossibleValuesForAddressBookAccountTargetContentType() []string {
	return []string{
		string(AddressBookAccountTargetContentTypeaddressBook),
		string(AddressBookAccountTargetContentTypeincludeAll),
		string(AddressBookAccountTargetContentTypeunknown),
	}
}

func parseAddressBookAccountTargetContentType(input string) (*AddressBookAccountTargetContentType, error) {
	vals := map[string]AddressBookAccountTargetContentType{
		"addressbook": AddressBookAccountTargetContentTypeaddressBook,
		"includeall":  AddressBookAccountTargetContentTypeincludeAll,
		"unknown":     AddressBookAccountTargetContentTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddressBookAccountTargetContentType(input)
	return &out, nil
}
