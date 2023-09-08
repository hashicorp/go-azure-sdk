package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81TrustedRootCertificateDestinationStore string

const (
	Windows81TrustedRootCertificateDestinationStorecomputerCertStoreIntermediate Windows81TrustedRootCertificateDestinationStore = "ComputerCertStoreIntermediate"
	Windows81TrustedRootCertificateDestinationStorecomputerCertStoreRoot         Windows81TrustedRootCertificateDestinationStore = "ComputerCertStoreRoot"
	Windows81TrustedRootCertificateDestinationStoreuserCertStoreIntermediate     Windows81TrustedRootCertificateDestinationStore = "UserCertStoreIntermediate"
)

func PossibleValuesForWindows81TrustedRootCertificateDestinationStore() []string {
	return []string{
		string(Windows81TrustedRootCertificateDestinationStorecomputerCertStoreIntermediate),
		string(Windows81TrustedRootCertificateDestinationStorecomputerCertStoreRoot),
		string(Windows81TrustedRootCertificateDestinationStoreuserCertStoreIntermediate),
	}
}

func parseWindows81TrustedRootCertificateDestinationStore(input string) (*Windows81TrustedRootCertificateDestinationStore, error) {
	vals := map[string]Windows81TrustedRootCertificateDestinationStore{
		"computercertstoreintermediate": Windows81TrustedRootCertificateDestinationStorecomputerCertStoreIntermediate,
		"computercertstoreroot":         Windows81TrustedRootCertificateDestinationStorecomputerCertStoreRoot,
		"usercertstoreintermediate":     Windows81TrustedRootCertificateDestinationStoreuserCertStoreIntermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81TrustedRootCertificateDestinationStore(input)
	return &out, nil
}
