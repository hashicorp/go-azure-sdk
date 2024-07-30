package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm string

const (
	AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm_Aes256 AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm = "aes256"
)

func PossibleValuesForAppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm() []string {
	return []string{
		string(AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm_Aes256),
	}
}

func (s *AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm(input string) (*AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm, error) {
	vals := map[string]AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm{
		"aes256": AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm_Aes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm(input)
	return &out, nil
}
