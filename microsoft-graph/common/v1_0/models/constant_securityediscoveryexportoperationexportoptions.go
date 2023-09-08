package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationExportOptions string

const (
	SecurityEdiscoveryExportOperationExportOptionsoriginalFiles  SecurityEdiscoveryExportOperationExportOptions = "OriginalFiles"
	SecurityEdiscoveryExportOperationExportOptionspdfReplacement SecurityEdiscoveryExportOperationExportOptions = "PdfReplacement"
	SecurityEdiscoveryExportOperationExportOptionstags           SecurityEdiscoveryExportOperationExportOptions = "Tags"
	SecurityEdiscoveryExportOperationExportOptionstext           SecurityEdiscoveryExportOperationExportOptions = "Text"
)

func PossibleValuesForSecurityEdiscoveryExportOperationExportOptions() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationExportOptionsoriginalFiles),
		string(SecurityEdiscoveryExportOperationExportOptionspdfReplacement),
		string(SecurityEdiscoveryExportOperationExportOptionstags),
		string(SecurityEdiscoveryExportOperationExportOptionstext),
	}
}

func parseSecurityEdiscoveryExportOperationExportOptions(input string) (*SecurityEdiscoveryExportOperationExportOptions, error) {
	vals := map[string]SecurityEdiscoveryExportOperationExportOptions{
		"originalfiles":  SecurityEdiscoveryExportOperationExportOptionsoriginalFiles,
		"pdfreplacement": SecurityEdiscoveryExportOperationExportOptionspdfReplacement,
		"tags":           SecurityEdiscoveryExportOperationExportOptionstags,
		"text":           SecurityEdiscoveryExportOperationExportOptionstext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationExportOptions(input)
	return &out, nil
}
