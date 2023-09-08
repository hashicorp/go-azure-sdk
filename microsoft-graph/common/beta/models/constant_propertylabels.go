package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyLabels string

const (
	PropertyLabelsauthors              PropertyLabels = "Authors"
	PropertyLabelscreatedBy            PropertyLabels = "CreatedBy"
	PropertyLabelscreatedDateTime      PropertyLabels = "CreatedDateTime"
	PropertyLabelsfileExtension        PropertyLabels = "FileExtension"
	PropertyLabelsfileName             PropertyLabels = "FileName"
	PropertyLabelslastModifiedBy       PropertyLabels = "LastModifiedBy"
	PropertyLabelslastModifiedDateTime PropertyLabels = "LastModifiedDateTime"
	PropertyLabelstitle                PropertyLabels = "Title"
	PropertyLabelsurl                  PropertyLabels = "Url"
)

func PossibleValuesForPropertyLabels() []string {
	return []string{
		string(PropertyLabelsauthors),
		string(PropertyLabelscreatedBy),
		string(PropertyLabelscreatedDateTime),
		string(PropertyLabelsfileExtension),
		string(PropertyLabelsfileName),
		string(PropertyLabelslastModifiedBy),
		string(PropertyLabelslastModifiedDateTime),
		string(PropertyLabelstitle),
		string(PropertyLabelsurl),
	}
}

func parsePropertyLabels(input string) (*PropertyLabels, error) {
	vals := map[string]PropertyLabels{
		"authors":              PropertyLabelsauthors,
		"createdby":            PropertyLabelscreatedBy,
		"createddatetime":      PropertyLabelscreatedDateTime,
		"fileextension":        PropertyLabelsfileExtension,
		"filename":             PropertyLabelsfileName,
		"lastmodifiedby":       PropertyLabelslastModifiedBy,
		"lastmodifieddatetime": PropertyLabelslastModifiedDateTime,
		"title":                PropertyLabelstitle,
		"url":                  PropertyLabelsurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertyLabels(input)
	return &out, nil
}
