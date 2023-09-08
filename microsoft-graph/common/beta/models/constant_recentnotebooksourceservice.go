package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecentNotebookSourceService string

const (
	RecentNotebookSourceServiceOnPremOneDriveForBusiness RecentNotebookSourceService = "OnPremOneDriveForBusiness"
	RecentNotebookSourceServiceOneDrive                  RecentNotebookSourceService = "OneDrive"
	RecentNotebookSourceServiceOneDriveForBusiness       RecentNotebookSourceService = "OneDriveForBusiness"
	RecentNotebookSourceServiceUnknown                   RecentNotebookSourceService = "Unknown"
)

func PossibleValuesForRecentNotebookSourceService() []string {
	return []string{
		string(RecentNotebookSourceServiceOnPremOneDriveForBusiness),
		string(RecentNotebookSourceServiceOneDrive),
		string(RecentNotebookSourceServiceOneDriveForBusiness),
		string(RecentNotebookSourceServiceUnknown),
	}
}

func parseRecentNotebookSourceService(input string) (*RecentNotebookSourceService, error) {
	vals := map[string]RecentNotebookSourceService{
		"onpremonedriveforbusiness": RecentNotebookSourceServiceOnPremOneDriveForBusiness,
		"onedrive":                  RecentNotebookSourceServiceOneDrive,
		"onedriveforbusiness":       RecentNotebookSourceServiceOneDriveForBusiness,
		"unknown":                   RecentNotebookSourceServiceUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecentNotebookSourceService(input)
	return &out, nil
}
