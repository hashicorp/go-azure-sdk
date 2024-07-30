package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationExportCriteria string

const (
	SecurityEdiscoverySearchExportOperationExportCriteria_PartiallyIndexed SecurityEdiscoverySearchExportOperationExportCriteria = "partiallyIndexed"
	SecurityEdiscoverySearchExportOperationExportCriteria_SearchHits       SecurityEdiscoverySearchExportOperationExportCriteria = "searchHits"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationExportCriteria() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationExportCriteria_PartiallyIndexed),
		string(SecurityEdiscoverySearchExportOperationExportCriteria_SearchHits),
	}
}

func (s *SecurityEdiscoverySearchExportOperationExportCriteria) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationExportCriteria(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationExportCriteria(input string) (*SecurityEdiscoverySearchExportOperationExportCriteria, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationExportCriteria{
		"partiallyindexed": SecurityEdiscoverySearchExportOperationExportCriteria_PartiallyIndexed,
		"searchhits":       SecurityEdiscoverySearchExportOperationExportCriteria_SearchHits,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationExportCriteria(input)
	return &out, nil
}
