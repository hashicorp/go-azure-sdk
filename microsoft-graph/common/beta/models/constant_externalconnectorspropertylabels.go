package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyLabels string

const (
	ExternalConnectorsPropertyLabelsauthors              ExternalConnectorsPropertyLabels = "Authors"
	ExternalConnectorsPropertyLabelscontainerName        ExternalConnectorsPropertyLabels = "ContainerName"
	ExternalConnectorsPropertyLabelscontainerUrl         ExternalConnectorsPropertyLabels = "ContainerUrl"
	ExternalConnectorsPropertyLabelscreatedBy            ExternalConnectorsPropertyLabels = "CreatedBy"
	ExternalConnectorsPropertyLabelscreatedDateTime      ExternalConnectorsPropertyLabels = "CreatedDateTime"
	ExternalConnectorsPropertyLabelsfileExtension        ExternalConnectorsPropertyLabels = "FileExtension"
	ExternalConnectorsPropertyLabelsfileName             ExternalConnectorsPropertyLabels = "FileName"
	ExternalConnectorsPropertyLabelsiconUrl              ExternalConnectorsPropertyLabels = "IconUrl"
	ExternalConnectorsPropertyLabelslastModifiedBy       ExternalConnectorsPropertyLabels = "LastModifiedBy"
	ExternalConnectorsPropertyLabelslastModifiedDateTime ExternalConnectorsPropertyLabels = "LastModifiedDateTime"
	ExternalConnectorsPropertyLabelstitle                ExternalConnectorsPropertyLabels = "Title"
	ExternalConnectorsPropertyLabelsurl                  ExternalConnectorsPropertyLabels = "Url"
)

func PossibleValuesForExternalConnectorsPropertyLabels() []string {
	return []string{
		string(ExternalConnectorsPropertyLabelsauthors),
		string(ExternalConnectorsPropertyLabelscontainerName),
		string(ExternalConnectorsPropertyLabelscontainerUrl),
		string(ExternalConnectorsPropertyLabelscreatedBy),
		string(ExternalConnectorsPropertyLabelscreatedDateTime),
		string(ExternalConnectorsPropertyLabelsfileExtension),
		string(ExternalConnectorsPropertyLabelsfileName),
		string(ExternalConnectorsPropertyLabelsiconUrl),
		string(ExternalConnectorsPropertyLabelslastModifiedBy),
		string(ExternalConnectorsPropertyLabelslastModifiedDateTime),
		string(ExternalConnectorsPropertyLabelstitle),
		string(ExternalConnectorsPropertyLabelsurl),
	}
}

func parseExternalConnectorsPropertyLabels(input string) (*ExternalConnectorsPropertyLabels, error) {
	vals := map[string]ExternalConnectorsPropertyLabels{
		"authors":              ExternalConnectorsPropertyLabelsauthors,
		"containername":        ExternalConnectorsPropertyLabelscontainerName,
		"containerurl":         ExternalConnectorsPropertyLabelscontainerUrl,
		"createdby":            ExternalConnectorsPropertyLabelscreatedBy,
		"createddatetime":      ExternalConnectorsPropertyLabelscreatedDateTime,
		"fileextension":        ExternalConnectorsPropertyLabelsfileExtension,
		"filename":             ExternalConnectorsPropertyLabelsfileName,
		"iconurl":              ExternalConnectorsPropertyLabelsiconUrl,
		"lastmodifiedby":       ExternalConnectorsPropertyLabelslastModifiedBy,
		"lastmodifieddatetime": ExternalConnectorsPropertyLabelslastModifiedDateTime,
		"title":                ExternalConnectorsPropertyLabelstitle,
		"url":                  ExternalConnectorsPropertyLabelsurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyLabels(input)
	return &out, nil
}
