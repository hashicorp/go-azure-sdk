package artifacts

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileUploadOptions string

const (
	FileUploadOptionsNone                            FileUploadOptions = "None"
	FileUploadOptionsUploadFilesAndGenerateSasTokens FileUploadOptions = "UploadFilesAndGenerateSasTokens"
)

func PossibleValuesForFileUploadOptions() []string {
	return []string{
		string(FileUploadOptionsNone),
		string(FileUploadOptionsUploadFilesAndGenerateSasTokens),
	}
}

func parseFileUploadOptions(input string) (*FileUploadOptions, error) {
	vals := map[string]FileUploadOptions{
		"none":                            FileUploadOptionsNone,
		"uploadfilesandgeneratesastokens": FileUploadOptionsUploadFilesAndGenerateSasTokens,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileUploadOptions(input)
	return &out, nil
}
