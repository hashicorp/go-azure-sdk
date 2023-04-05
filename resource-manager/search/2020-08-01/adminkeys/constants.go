package adminkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminKeyKind string

const (
	AdminKeyKindPrimary   AdminKeyKind = "primary"
	AdminKeyKindSecondary AdminKeyKind = "secondary"
)

func PossibleValuesForAdminKeyKind() []string {
	return []string{
		string(AdminKeyKindPrimary),
		string(AdminKeyKindSecondary),
	}
}
