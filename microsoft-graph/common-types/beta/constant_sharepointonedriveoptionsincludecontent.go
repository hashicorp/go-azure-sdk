package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointOneDriveOptionsIncludeContent string

const (
	SharePointOneDriveOptionsIncludeContent_PrivateContent SharePointOneDriveOptionsIncludeContent = "privateContent"
	SharePointOneDriveOptionsIncludeContent_SharedContent  SharePointOneDriveOptionsIncludeContent = "sharedContent"
)

func PossibleValuesForSharePointOneDriveOptionsIncludeContent() []string {
	return []string{
		string(SharePointOneDriveOptionsIncludeContent_PrivateContent),
		string(SharePointOneDriveOptionsIncludeContent_SharedContent),
	}
}

func (s *SharePointOneDriveOptionsIncludeContent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharePointOneDriveOptionsIncludeContent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharePointOneDriveOptionsIncludeContent(input string) (*SharePointOneDriveOptionsIncludeContent, error) {
	vals := map[string]SharePointOneDriveOptionsIncludeContent{
		"privatecontent": SharePointOneDriveOptionsIncludeContent_PrivateContent,
		"sharedcontent":  SharePointOneDriveOptionsIncludeContent_SharedContent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharePointOneDriveOptionsIncludeContent(input)
	return &out, nil
}
