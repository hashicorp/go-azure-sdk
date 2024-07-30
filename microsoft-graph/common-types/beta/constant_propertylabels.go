package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyLabels string

const (
	PropertyLabels_Authors              PropertyLabels = "authors"
	PropertyLabels_CreatedBy            PropertyLabels = "createdBy"
	PropertyLabels_CreatedDateTime      PropertyLabels = "createdDateTime"
	PropertyLabels_FileExtension        PropertyLabels = "fileExtension"
	PropertyLabels_FileName             PropertyLabels = "fileName"
	PropertyLabels_LastModifiedBy       PropertyLabels = "lastModifiedBy"
	PropertyLabels_LastModifiedDateTime PropertyLabels = "lastModifiedDateTime"
	PropertyLabels_Title                PropertyLabels = "title"
	PropertyLabels_Url                  PropertyLabels = "url"
)

func PossibleValuesForPropertyLabels() []string {
	return []string{
		string(PropertyLabels_Authors),
		string(PropertyLabels_CreatedBy),
		string(PropertyLabels_CreatedDateTime),
		string(PropertyLabels_FileExtension),
		string(PropertyLabels_FileName),
		string(PropertyLabels_LastModifiedBy),
		string(PropertyLabels_LastModifiedDateTime),
		string(PropertyLabels_Title),
		string(PropertyLabels_Url),
	}
}

func (s *PropertyLabels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePropertyLabels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePropertyLabels(input string) (*PropertyLabels, error) {
	vals := map[string]PropertyLabels{
		"authors":              PropertyLabels_Authors,
		"createdby":            PropertyLabels_CreatedBy,
		"createddatetime":      PropertyLabels_CreatedDateTime,
		"fileextension":        PropertyLabels_FileExtension,
		"filename":             PropertyLabels_FileName,
		"lastmodifiedby":       PropertyLabels_LastModifiedBy,
		"lastmodifieddatetime": PropertyLabels_LastModifiedDateTime,
		"title":                PropertyLabels_Title,
		"url":                  PropertyLabels_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertyLabels(input)
	return &out, nil
}
