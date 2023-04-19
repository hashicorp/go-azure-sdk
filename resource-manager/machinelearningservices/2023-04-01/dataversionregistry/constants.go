package dataversionregistry

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataType string

const (
	DataTypeMltable   DataType = "mltable"
	DataTypeUriFile   DataType = "uri_file"
	DataTypeUriFolder DataType = "uri_folder"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataTypeMltable),
		string(DataTypeUriFile),
		string(DataTypeUriFolder),
	}
}

func parseDataType(input string) (*DataType, error) {
	vals := map[string]DataType{
		"mltable":    DataTypeMltable,
		"uri_file":   DataTypeUriFile,
		"uri_folder": DataTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataType(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}

type PendingUploadCredentialType string

const (
	PendingUploadCredentialTypeSAS PendingUploadCredentialType = "SAS"
)

func PossibleValuesForPendingUploadCredentialType() []string {
	return []string{
		string(PendingUploadCredentialTypeSAS),
	}
}

func parsePendingUploadCredentialType(input string) (*PendingUploadCredentialType, error) {
	vals := map[string]PendingUploadCredentialType{
		"sas": PendingUploadCredentialTypeSAS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PendingUploadCredentialType(input)
	return &out, nil
}

type PendingUploadType string

const (
	PendingUploadTypeNone                   PendingUploadType = "None"
	PendingUploadTypeTemporaryBlobReference PendingUploadType = "TemporaryBlobReference"
)

func PossibleValuesForPendingUploadType() []string {
	return []string{
		string(PendingUploadTypeNone),
		string(PendingUploadTypeTemporaryBlobReference),
	}
}

func parsePendingUploadType(input string) (*PendingUploadType, error) {
	vals := map[string]PendingUploadType{
		"none":                   PendingUploadTypeNone,
		"temporaryblobreference": PendingUploadTypeTemporaryBlobReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PendingUploadType(input)
	return &out, nil
}
