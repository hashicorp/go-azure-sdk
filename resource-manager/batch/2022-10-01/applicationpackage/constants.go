package applicationpackage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageState string

const (
	PackageStateActive  PackageState = "Active"
	PackageStatePending PackageState = "Pending"
)

func PossibleValuesForPackageState() []string {
	return []string{
		string(PackageStateActive),
		string(PackageStatePending),
	}
}
