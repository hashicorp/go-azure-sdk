package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm string

const (
	AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithmaes256 AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm = "Aes256"
)

func PossibleValuesForAppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm() []string {
	return []string{
		string(AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithmaes256),
	}
}

func parseAppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm(input string) (*AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm, error) {
	vals := map[string]AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm{
		"aes256": AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithmaes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm(input)
	return &out, nil
}
