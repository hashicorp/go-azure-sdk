package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalAddressType string

const (
	PhysicalAddressTypebusiness PhysicalAddressType = "Business"
	PhysicalAddressTypehome     PhysicalAddressType = "Home"
	PhysicalAddressTypeother    PhysicalAddressType = "Other"
	PhysicalAddressTypeunknown  PhysicalAddressType = "Unknown"
)

func PossibleValuesForPhysicalAddressType() []string {
	return []string{
		string(PhysicalAddressTypebusiness),
		string(PhysicalAddressTypehome),
		string(PhysicalAddressTypeother),
		string(PhysicalAddressTypeunknown),
	}
}

func parsePhysicalAddressType(input string) (*PhysicalAddressType, error) {
	vals := map[string]PhysicalAddressType{
		"business": PhysicalAddressTypebusiness,
		"home":     PhysicalAddressTypehome,
		"other":    PhysicalAddressTypeother,
		"unknown":  PhysicalAddressTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhysicalAddressType(input)
	return &out, nil
}
