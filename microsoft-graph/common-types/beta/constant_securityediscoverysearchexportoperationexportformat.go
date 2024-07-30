package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationExportFormat string

const (
	SecurityEdiscoverySearchExportOperationExportFormat_Eml SecurityEdiscoverySearchExportOperationExportFormat = "eml"
	SecurityEdiscoverySearchExportOperationExportFormat_Msg SecurityEdiscoverySearchExportOperationExportFormat = "msg"
	SecurityEdiscoverySearchExportOperationExportFormat_Pst SecurityEdiscoverySearchExportOperationExportFormat = "pst"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationExportFormat() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationExportFormat_Eml),
		string(SecurityEdiscoverySearchExportOperationExportFormat_Msg),
		string(SecurityEdiscoverySearchExportOperationExportFormat_Pst),
	}
}

func (s *SecurityEdiscoverySearchExportOperationExportFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationExportFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationExportFormat(input string) (*SecurityEdiscoverySearchExportOperationExportFormat, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationExportFormat{
		"eml": SecurityEdiscoverySearchExportOperationExportFormat_Eml,
		"msg": SecurityEdiscoverySearchExportOperationExportFormat_Msg,
		"pst": SecurityEdiscoverySearchExportOperationExportFormat_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationExportFormat(input)
	return &out, nil
}
