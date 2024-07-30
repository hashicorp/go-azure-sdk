package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateStore string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_Machine AndroidDeviceOwnerPkcsCertificateProfileCertificateStore = "machine"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_User    AndroidDeviceOwnerPkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_Machine),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_User),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateStore(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateStore{
		"machine": AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_Machine,
		"user":    AndroidDeviceOwnerPkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
