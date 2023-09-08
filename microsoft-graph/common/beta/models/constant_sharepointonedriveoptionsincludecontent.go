package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointOneDriveOptionsIncludeContent string

const (
	SharePointOneDriveOptionsIncludeContentprivateContent SharePointOneDriveOptionsIncludeContent = "PrivateContent"
	SharePointOneDriveOptionsIncludeContentsharedContent  SharePointOneDriveOptionsIncludeContent = "SharedContent"
)

func PossibleValuesForSharePointOneDriveOptionsIncludeContent() []string {
	return []string{
		string(SharePointOneDriveOptionsIncludeContentprivateContent),
		string(SharePointOneDriveOptionsIncludeContentsharedContent),
	}
}

func parseSharePointOneDriveOptionsIncludeContent(input string) (*SharePointOneDriveOptionsIncludeContent, error) {
	vals := map[string]SharePointOneDriveOptionsIncludeContent{
		"privatecontent": SharePointOneDriveOptionsIncludeContentprivateContent,
		"sharedcontent":  SharePointOneDriveOptionsIncludeContentsharedContent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharePointOneDriveOptionsIncludeContent(input)
	return &out, nil
}
