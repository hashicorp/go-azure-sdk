package integrationruntimeobjectmetadata

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
