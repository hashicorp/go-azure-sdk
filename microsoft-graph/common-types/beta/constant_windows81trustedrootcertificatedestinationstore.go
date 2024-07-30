package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81TrustedRootCertificateDestinationStore string

const (
	Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate Windows81TrustedRootCertificateDestinationStore = "computerCertStoreIntermediate"
	Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreRoot         Windows81TrustedRootCertificateDestinationStore = "computerCertStoreRoot"
	Windows81TrustedRootCertificateDestinationStore_UserCertStoreIntermediate     Windows81TrustedRootCertificateDestinationStore = "userCertStoreIntermediate"
)

func PossibleValuesForWindows81TrustedRootCertificateDestinationStore() []string {
	return []string{
		string(Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate),
		string(Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreRoot),
		string(Windows81TrustedRootCertificateDestinationStore_UserCertStoreIntermediate),
	}
}

func (s *Windows81TrustedRootCertificateDestinationStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81TrustedRootCertificateDestinationStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81TrustedRootCertificateDestinationStore(input string) (*Windows81TrustedRootCertificateDestinationStore, error) {
	vals := map[string]Windows81TrustedRootCertificateDestinationStore{
		"computercertstoreintermediate": Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreIntermediate,
		"computercertstoreroot":         Windows81TrustedRootCertificateDestinationStore_ComputerCertStoreRoot,
		"usercertstoreintermediate":     Windows81TrustedRootCertificateDestinationStore_UserCertStoreIntermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81TrustedRootCertificateDestinationStore(input)
	return &out, nil
}
