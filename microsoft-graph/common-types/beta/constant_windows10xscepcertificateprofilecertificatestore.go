package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileCertificateStore string

const (
	Windows10XSCEPCertificateProfileCertificateStore_Machine Windows10XSCEPCertificateProfileCertificateStore = "machine"
	Windows10XSCEPCertificateProfileCertificateStore_User    Windows10XSCEPCertificateProfileCertificateStore = "user"
)

func PossibleValuesForWindows10XSCEPCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileCertificateStore_Machine),
		string(Windows10XSCEPCertificateProfileCertificateStore_User),
	}
}

func (s *Windows10XSCEPCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileCertificateStore(input string) (*Windows10XSCEPCertificateProfileCertificateStore, error) {
	vals := map[string]Windows10XSCEPCertificateProfileCertificateStore{
		"machine": Windows10XSCEPCertificateProfileCertificateStore_Machine,
		"user":    Windows10XSCEPCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileCertificateStore(input)
	return &out, nil
}
