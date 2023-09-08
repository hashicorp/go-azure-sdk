package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XTrustedRootCertificateDestinationStore string

const (
	Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreIntermediate Windows10XTrustedRootCertificateDestinationStore = "ComputerCertStoreIntermediate"
	Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreRoot         Windows10XTrustedRootCertificateDestinationStore = "ComputerCertStoreRoot"
	Windows10XTrustedRootCertificateDestinationStoreuserCertStoreIntermediate     Windows10XTrustedRootCertificateDestinationStore = "UserCertStoreIntermediate"
)

func PossibleValuesForWindows10XTrustedRootCertificateDestinationStore() []string {
	return []string{
		string(Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreIntermediate),
		string(Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreRoot),
		string(Windows10XTrustedRootCertificateDestinationStoreuserCertStoreIntermediate),
	}
}

func parseWindows10XTrustedRootCertificateDestinationStore(input string) (*Windows10XTrustedRootCertificateDestinationStore, error) {
	vals := map[string]Windows10XTrustedRootCertificateDestinationStore{
		"computercertstoreintermediate": Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreIntermediate,
		"computercertstoreroot":         Windows10XTrustedRootCertificateDestinationStorecomputerCertStoreRoot,
		"usercertstoreintermediate":     Windows10XTrustedRootCertificateDestinationStoreuserCertStoreIntermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XTrustedRootCertificateDestinationStore(input)
	return &out, nil
}
