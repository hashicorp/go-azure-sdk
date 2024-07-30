package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationExportOptions string

const (
	EdiscoveryCaseExportOperationExportOptions_FileInfo       EdiscoveryCaseExportOperationExportOptions = "fileInfo"
	EdiscoveryCaseExportOperationExportOptions_OriginalFiles  EdiscoveryCaseExportOperationExportOptions = "originalFiles"
	EdiscoveryCaseExportOperationExportOptions_PdfReplacement EdiscoveryCaseExportOperationExportOptions = "pdfReplacement"
	EdiscoveryCaseExportOperationExportOptions_Tags           EdiscoveryCaseExportOperationExportOptions = "tags"
	EdiscoveryCaseExportOperationExportOptions_Text           EdiscoveryCaseExportOperationExportOptions = "text"
)

func PossibleValuesForEdiscoveryCaseExportOperationExportOptions() []string {
	return []string{
		string(EdiscoveryCaseExportOperationExportOptions_FileInfo),
		string(EdiscoveryCaseExportOperationExportOptions_OriginalFiles),
		string(EdiscoveryCaseExportOperationExportOptions_PdfReplacement),
		string(EdiscoveryCaseExportOperationExportOptions_Tags),
		string(EdiscoveryCaseExportOperationExportOptions_Text),
	}
}

func (s *EdiscoveryCaseExportOperationExportOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseExportOperationExportOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseExportOperationExportOptions(input string) (*EdiscoveryCaseExportOperationExportOptions, error) {
	vals := map[string]EdiscoveryCaseExportOperationExportOptions{
		"fileinfo":       EdiscoveryCaseExportOperationExportOptions_FileInfo,
		"originalfiles":  EdiscoveryCaseExportOperationExportOptions_OriginalFiles,
		"pdfreplacement": EdiscoveryCaseExportOperationExportOptions_PdfReplacement,
		"tags":           EdiscoveryCaseExportOperationExportOptions_Tags,
		"text":           EdiscoveryCaseExportOperationExportOptions_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationExportOptions(input)
	return &out, nil
}
