package mhsmprivatelinkresources

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmSkuFamily string

const (
	ManagedHsmSkuFamilyB ManagedHsmSkuFamily = "B"
)

func PossibleValuesForManagedHsmSkuFamily() []string {
	return []string{
		string(ManagedHsmSkuFamilyB),
	}
}

func parseManagedHsmSkuFamily(input string) (*ManagedHsmSkuFamily, error) {
	vals := map[string]ManagedHsmSkuFamily{
		"b": ManagedHsmSkuFamilyB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuFamily(input)
	return &out, nil
}

type ManagedHsmSkuName string

const (
	ManagedHsmSkuNameCustomBThreeTwo ManagedHsmSkuName = "Custom_B32"
	ManagedHsmSkuNameStandardBOne    ManagedHsmSkuName = "Standard_B1"
)

func PossibleValuesForManagedHsmSkuName() []string {
	return []string{
		string(ManagedHsmSkuNameCustomBThreeTwo),
		string(ManagedHsmSkuNameStandardBOne),
	}
}

func parseManagedHsmSkuName(input string) (*ManagedHsmSkuName, error) {
	vals := map[string]ManagedHsmSkuName{
		"custom_b32":  ManagedHsmSkuNameCustomBThreeTwo,
		"standard_b1": ManagedHsmSkuNameStandardBOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuName(input)
	return &out, nil
}
