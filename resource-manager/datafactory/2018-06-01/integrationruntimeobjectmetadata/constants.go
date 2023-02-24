package integrationruntimeobjectmetadata

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SsisObjectMetadataType string

const (
	SsisObjectMetadataTypeEnvironment SsisObjectMetadataType = "Environment"
	SsisObjectMetadataTypeFolder      SsisObjectMetadataType = "Folder"
	SsisObjectMetadataTypePackage     SsisObjectMetadataType = "Package"
	SsisObjectMetadataTypeProject     SsisObjectMetadataType = "Project"
)

func PossibleValuesForSsisObjectMetadataType() []string {
	return []string{
		string(SsisObjectMetadataTypeEnvironment),
		string(SsisObjectMetadataTypeFolder),
		string(SsisObjectMetadataTypePackage),
		string(SsisObjectMetadataTypeProject),
	}
}

func parseSsisObjectMetadataType(input string) (*SsisObjectMetadataType, error) {
	vals := map[string]SsisObjectMetadataType{
		"environment": SsisObjectMetadataTypeEnvironment,
		"folder":      SsisObjectMetadataTypeFolder,
		"package":     SsisObjectMetadataTypePackage,
		"project":     SsisObjectMetadataTypeProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SsisObjectMetadataType(input)
	return &out, nil
}
