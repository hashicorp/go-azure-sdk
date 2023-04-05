package environmentversion

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
