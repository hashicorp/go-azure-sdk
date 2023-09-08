package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiInformationPackageType string

const (
	Win32LobAppMsiInformationPackageTypedualPurpose Win32LobAppMsiInformationPackageType = "DualPurpose"
	Win32LobAppMsiInformationPackageTypeperMachine  Win32LobAppMsiInformationPackageType = "PerMachine"
	Win32LobAppMsiInformationPackageTypeperUser     Win32LobAppMsiInformationPackageType = "PerUser"
)

func PossibleValuesForWin32LobAppMsiInformationPackageType() []string {
	return []string{
		string(Win32LobAppMsiInformationPackageTypedualPurpose),
		string(Win32LobAppMsiInformationPackageTypeperMachine),
		string(Win32LobAppMsiInformationPackageTypeperUser),
	}
}

func parseWin32LobAppMsiInformationPackageType(input string) (*Win32LobAppMsiInformationPackageType, error) {
	vals := map[string]Win32LobAppMsiInformationPackageType{
		"dualpurpose": Win32LobAppMsiInformationPackageTypedualPurpose,
		"permachine":  Win32LobAppMsiInformationPackageTypeperMachine,
		"peruser":     Win32LobAppMsiInformationPackageTypeperUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppMsiInformationPackageType(input)
	return &out, nil
}
