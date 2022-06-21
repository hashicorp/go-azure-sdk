package environmentversion

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentType string

const (
	EnvironmentTypeCurated     EnvironmentType = "Curated"
	EnvironmentTypeUserCreated EnvironmentType = "UserCreated"
)

func PossibleValuesForEnvironmentType() []string {
	return []string{
		string(EnvironmentTypeCurated),
		string(EnvironmentTypeUserCreated),
	}
}

func parseEnvironmentType(input string) (*EnvironmentType, error) {
	vals := map[string]EnvironmentType{
		"curated":     EnvironmentTypeCurated,
		"usercreated": EnvironmentTypeUserCreated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnvironmentType(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   OperatingSystemType = "Linux"
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
)

func PossibleValuesForOperatingSystemType() []string {
	return []string{
		string(OperatingSystemTypeLinux),
		string(OperatingSystemTypeWindows),
	}
}

func parseOperatingSystemType(input string) (*OperatingSystemType, error) {
	vals := map[string]OperatingSystemType{
		"linux":   OperatingSystemTypeLinux,
		"windows": OperatingSystemTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemType(input)
	return &out, nil
}
