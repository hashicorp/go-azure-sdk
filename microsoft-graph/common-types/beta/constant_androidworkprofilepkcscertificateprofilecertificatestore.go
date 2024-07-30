package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileCertificateStore string

const (
	AndroidWorkProfilePkcsCertificateProfileCertificateStore_Machine AndroidWorkProfilePkcsCertificateProfileCertificateStore = "machine"
	AndroidWorkProfilePkcsCertificateProfileCertificateStore_User    AndroidWorkProfilePkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileCertificateStore_Machine),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateStore_User),
	}
}

func (s *AndroidWorkProfilePkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfilePkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfilePkcsCertificateProfileCertificateStore(input string) (*AndroidWorkProfilePkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileCertificateStore{
		"machine": AndroidWorkProfilePkcsCertificateProfileCertificateStore_Machine,
		"user":    AndroidWorkProfilePkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
