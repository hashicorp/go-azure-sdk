package resourcegroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportTemplateOutputFormat string

const (
	ExportTemplateOutputFormatBicep ExportTemplateOutputFormat = "Bicep"
	ExportTemplateOutputFormatJson  ExportTemplateOutputFormat = "Json"
)

func PossibleValuesForExportTemplateOutputFormat() []string {
	return []string{
		string(ExportTemplateOutputFormatBicep),
		string(ExportTemplateOutputFormatJson),
	}
}

func (s *ExportTemplateOutputFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExportTemplateOutputFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExportTemplateOutputFormat(input string) (*ExportTemplateOutputFormat, error) {
	vals := map[string]ExportTemplateOutputFormat{
		"bicep": ExportTemplateOutputFormatBicep,
		"json":  ExportTemplateOutputFormatJson,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportTemplateOutputFormat(input)
	return &out, nil
}
