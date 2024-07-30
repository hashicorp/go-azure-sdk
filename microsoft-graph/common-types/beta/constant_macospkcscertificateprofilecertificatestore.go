package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileCertificateStore string

const (
	MacOSPkcsCertificateProfileCertificateStore_Machine MacOSPkcsCertificateProfileCertificateStore = "machine"
	MacOSPkcsCertificateProfileCertificateStore_User    MacOSPkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForMacOSPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(MacOSPkcsCertificateProfileCertificateStore_Machine),
		string(MacOSPkcsCertificateProfileCertificateStore_User),
	}
}

func (s *MacOSPkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPkcsCertificateProfileCertificateStore(input string) (*MacOSPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]MacOSPkcsCertificateProfileCertificateStore{
		"machine": MacOSPkcsCertificateProfileCertificateStore_Machine,
		"user":    MacOSPkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
