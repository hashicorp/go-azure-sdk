package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationExportOptions string

const (
	EdiscoveryCaseExportOperationExportOptionsfileInfo       EdiscoveryCaseExportOperationExportOptions = "FileInfo"
	EdiscoveryCaseExportOperationExportOptionsoriginalFiles  EdiscoveryCaseExportOperationExportOptions = "OriginalFiles"
	EdiscoveryCaseExportOperationExportOptionspdfReplacement EdiscoveryCaseExportOperationExportOptions = "PdfReplacement"
	EdiscoveryCaseExportOperationExportOptionstags           EdiscoveryCaseExportOperationExportOptions = "Tags"
	EdiscoveryCaseExportOperationExportOptionstext           EdiscoveryCaseExportOperationExportOptions = "Text"
)

func PossibleValuesForEdiscoveryCaseExportOperationExportOptions() []string {
	return []string{
		string(EdiscoveryCaseExportOperationExportOptionsfileInfo),
		string(EdiscoveryCaseExportOperationExportOptionsoriginalFiles),
		string(EdiscoveryCaseExportOperationExportOptionspdfReplacement),
		string(EdiscoveryCaseExportOperationExportOptionstags),
		string(EdiscoveryCaseExportOperationExportOptionstext),
	}
}

func parseEdiscoveryCaseExportOperationExportOptions(input string) (*EdiscoveryCaseExportOperationExportOptions, error) {
	vals := map[string]EdiscoveryCaseExportOperationExportOptions{
		"fileinfo":       EdiscoveryCaseExportOperationExportOptionsfileInfo,
		"originalfiles":  EdiscoveryCaseExportOperationExportOptionsoriginalFiles,
		"pdfreplacement": EdiscoveryCaseExportOperationExportOptionspdfReplacement,
		"tags":           EdiscoveryCaseExportOperationExportOptionstags,
		"text":           EdiscoveryCaseExportOperationExportOptionstext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationExportOptions(input)
	return &out, nil
}
