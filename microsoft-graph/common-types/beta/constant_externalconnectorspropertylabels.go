package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyLabels string

const (
	ExternalConnectorsPropertyLabels_Authors              ExternalConnectorsPropertyLabels = "authors"
	ExternalConnectorsPropertyLabels_ContainerName        ExternalConnectorsPropertyLabels = "containerName"
	ExternalConnectorsPropertyLabels_ContainerUrl         ExternalConnectorsPropertyLabels = "containerUrl"
	ExternalConnectorsPropertyLabels_CreatedBy            ExternalConnectorsPropertyLabels = "createdBy"
	ExternalConnectorsPropertyLabels_CreatedDateTime      ExternalConnectorsPropertyLabels = "createdDateTime"
	ExternalConnectorsPropertyLabels_FileExtension        ExternalConnectorsPropertyLabels = "fileExtension"
	ExternalConnectorsPropertyLabels_FileName             ExternalConnectorsPropertyLabels = "fileName"
	ExternalConnectorsPropertyLabels_IconUrl              ExternalConnectorsPropertyLabels = "iconUrl"
	ExternalConnectorsPropertyLabels_LastModifiedBy       ExternalConnectorsPropertyLabels = "lastModifiedBy"
	ExternalConnectorsPropertyLabels_LastModifiedDateTime ExternalConnectorsPropertyLabels = "lastModifiedDateTime"
	ExternalConnectorsPropertyLabels_Title                ExternalConnectorsPropertyLabels = "title"
	ExternalConnectorsPropertyLabels_Url                  ExternalConnectorsPropertyLabels = "url"
)

func PossibleValuesForExternalConnectorsPropertyLabels() []string {
	return []string{
		string(ExternalConnectorsPropertyLabels_Authors),
		string(ExternalConnectorsPropertyLabels_ContainerName),
		string(ExternalConnectorsPropertyLabels_ContainerUrl),
		string(ExternalConnectorsPropertyLabels_CreatedBy),
		string(ExternalConnectorsPropertyLabels_CreatedDateTime),
		string(ExternalConnectorsPropertyLabels_FileExtension),
		string(ExternalConnectorsPropertyLabels_FileName),
		string(ExternalConnectorsPropertyLabels_IconUrl),
		string(ExternalConnectorsPropertyLabels_LastModifiedBy),
		string(ExternalConnectorsPropertyLabels_LastModifiedDateTime),
		string(ExternalConnectorsPropertyLabels_Title),
		string(ExternalConnectorsPropertyLabels_Url),
	}
}

func (s *ExternalConnectorsPropertyLabels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsPropertyLabels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsPropertyLabels(input string) (*ExternalConnectorsPropertyLabels, error) {
	vals := map[string]ExternalConnectorsPropertyLabels{
		"authors":              ExternalConnectorsPropertyLabels_Authors,
		"containername":        ExternalConnectorsPropertyLabels_ContainerName,
		"containerurl":         ExternalConnectorsPropertyLabels_ContainerUrl,
		"createdby":            ExternalConnectorsPropertyLabels_CreatedBy,
		"createddatetime":      ExternalConnectorsPropertyLabels_CreatedDateTime,
		"fileextension":        ExternalConnectorsPropertyLabels_FileExtension,
		"filename":             ExternalConnectorsPropertyLabels_FileName,
		"iconurl":              ExternalConnectorsPropertyLabels_IconUrl,
		"lastmodifiedby":       ExternalConnectorsPropertyLabels_LastModifiedBy,
		"lastmodifieddatetime": ExternalConnectorsPropertyLabels_LastModifiedDateTime,
		"title":                ExternalConnectorsPropertyLabels_Title,
		"url":                  ExternalConnectorsPropertyLabels_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyLabels(input)
	return &out, nil
}
