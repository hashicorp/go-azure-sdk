package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecentNotebookSourceService string

const (
	RecentNotebookSourceService_OnPremOneDriveForBusiness RecentNotebookSourceService = "OnPremOneDriveForBusiness"
	RecentNotebookSourceService_OneDrive                  RecentNotebookSourceService = "OneDrive"
	RecentNotebookSourceService_OneDriveForBusiness       RecentNotebookSourceService = "OneDriveForBusiness"
	RecentNotebookSourceService_Unknown                   RecentNotebookSourceService = "Unknown"
)

func PossibleValuesForRecentNotebookSourceService() []string {
	return []string{
		string(RecentNotebookSourceService_OnPremOneDriveForBusiness),
		string(RecentNotebookSourceService_OneDrive),
		string(RecentNotebookSourceService_OneDriveForBusiness),
		string(RecentNotebookSourceService_Unknown),
	}
}

func (s *RecentNotebookSourceService) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecentNotebookSourceService(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecentNotebookSourceService(input string) (*RecentNotebookSourceService, error) {
	vals := map[string]RecentNotebookSourceService{
		"onpremonedriveforbusiness": RecentNotebookSourceService_OnPremOneDriveForBusiness,
		"onedrive":                  RecentNotebookSourceService_OneDrive,
		"onedriveforbusiness":       RecentNotebookSourceService_OneDriveForBusiness,
		"unknown":                   RecentNotebookSourceService_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecentNotebookSourceService(input)
	return &out, nil
}
