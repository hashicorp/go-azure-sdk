package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemSourceApplication string

const (
	DriveItemSourceApplicationoffice     DriveItemSourceApplication = "Office"
	DriveItemSourceApplicationoneDrive   DriveItemSourceApplication = "OneDrive"
	DriveItemSourceApplicationpowerPoint DriveItemSourceApplication = "PowerPoint"
	DriveItemSourceApplicationsharePoint DriveItemSourceApplication = "SharePoint"
	DriveItemSourceApplicationstream     DriveItemSourceApplication = "Stream"
	DriveItemSourceApplicationteams      DriveItemSourceApplication = "Teams"
	DriveItemSourceApplicationyammer     DriveItemSourceApplication = "Yammer"
)

func PossibleValuesForDriveItemSourceApplication() []string {
	return []string{
		string(DriveItemSourceApplicationoffice),
		string(DriveItemSourceApplicationoneDrive),
		string(DriveItemSourceApplicationpowerPoint),
		string(DriveItemSourceApplicationsharePoint),
		string(DriveItemSourceApplicationstream),
		string(DriveItemSourceApplicationteams),
		string(DriveItemSourceApplicationyammer),
	}
}

func parseDriveItemSourceApplication(input string) (*DriveItemSourceApplication, error) {
	vals := map[string]DriveItemSourceApplication{
		"office":     DriveItemSourceApplicationoffice,
		"onedrive":   DriveItemSourceApplicationoneDrive,
		"powerpoint": DriveItemSourceApplicationpowerPoint,
		"sharepoint": DriveItemSourceApplicationsharePoint,
		"stream":     DriveItemSourceApplicationstream,
		"teams":      DriveItemSourceApplicationteams,
		"yammer":     DriveItemSourceApplicationyammer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriveItemSourceApplication(input)
	return &out, nil
}
