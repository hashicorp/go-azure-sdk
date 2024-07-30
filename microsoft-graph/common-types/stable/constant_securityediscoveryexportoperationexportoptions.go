package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationExportOptions string

const (
	SecurityEdiscoveryExportOperationExportOptions_OriginalFiles  SecurityEdiscoveryExportOperationExportOptions = "originalFiles"
	SecurityEdiscoveryExportOperationExportOptions_PdfReplacement SecurityEdiscoveryExportOperationExportOptions = "pdfReplacement"
	SecurityEdiscoveryExportOperationExportOptions_Tags           SecurityEdiscoveryExportOperationExportOptions = "tags"
	SecurityEdiscoveryExportOperationExportOptions_Text           SecurityEdiscoveryExportOperationExportOptions = "text"
)

func PossibleValuesForSecurityEdiscoveryExportOperationExportOptions() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationExportOptions_OriginalFiles),
		string(SecurityEdiscoveryExportOperationExportOptions_PdfReplacement),
		string(SecurityEdiscoveryExportOperationExportOptions_Tags),
		string(SecurityEdiscoveryExportOperationExportOptions_Text),
	}
}

func (s *SecurityEdiscoveryExportOperationExportOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryExportOperationExportOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryExportOperationExportOptions(input string) (*SecurityEdiscoveryExportOperationExportOptions, error) {
	vals := map[string]SecurityEdiscoveryExportOperationExportOptions{
		"originalfiles":  SecurityEdiscoveryExportOperationExportOptions_OriginalFiles,
		"pdfreplacement": SecurityEdiscoveryExportOperationExportOptions_PdfReplacement,
		"tags":           SecurityEdiscoveryExportOperationExportOptions_Tags,
		"text":           SecurityEdiscoveryExportOperationExportOptions_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationExportOptions(input)
	return &out, nil
}
