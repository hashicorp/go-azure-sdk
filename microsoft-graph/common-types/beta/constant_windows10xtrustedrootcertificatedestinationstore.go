package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XTrustedRootCertificateDestinationStore string

const (
	Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate Windows10XTrustedRootCertificateDestinationStore = "computerCertStoreIntermediate"
	Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreRoot         Windows10XTrustedRootCertificateDestinationStore = "computerCertStoreRoot"
	Windows10XTrustedRootCertificateDestinationStore_UserCertStoreIntermediate     Windows10XTrustedRootCertificateDestinationStore = "userCertStoreIntermediate"
)

func PossibleValuesForWindows10XTrustedRootCertificateDestinationStore() []string {
	return []string{
		string(Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate),
		string(Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreRoot),
		string(Windows10XTrustedRootCertificateDestinationStore_UserCertStoreIntermediate),
	}
}

func (s *Windows10XTrustedRootCertificateDestinationStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XTrustedRootCertificateDestinationStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XTrustedRootCertificateDestinationStore(input string) (*Windows10XTrustedRootCertificateDestinationStore, error) {
	vals := map[string]Windows10XTrustedRootCertificateDestinationStore{
		"computercertstoreintermediate": Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate,
		"computercertstoreroot":         Windows10XTrustedRootCertificateDestinationStore_ComputerCertStoreRoot,
		"usercertstoreintermediate":     Windows10XTrustedRootCertificateDestinationStore_UserCertStoreIntermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XTrustedRootCertificateDestinationStore(input)
	return &out, nil
}
