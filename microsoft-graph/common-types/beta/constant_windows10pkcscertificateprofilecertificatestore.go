package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileCertificateStore string

const (
	Windows10PkcsCertificateProfileCertificateStore_Machine Windows10PkcsCertificateProfileCertificateStore = "machine"
	Windows10PkcsCertificateProfileCertificateStore_User    Windows10PkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForWindows10PkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows10PkcsCertificateProfileCertificateStore_Machine),
		string(Windows10PkcsCertificateProfileCertificateStore_User),
	}
}

func (s *Windows10PkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PkcsCertificateProfileCertificateStore(input string) (*Windows10PkcsCertificateProfileCertificateStore, error) {
	vals := map[string]Windows10PkcsCertificateProfileCertificateStore{
		"machine": Windows10PkcsCertificateProfileCertificateStore_Machine,
		"user":    Windows10PkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
