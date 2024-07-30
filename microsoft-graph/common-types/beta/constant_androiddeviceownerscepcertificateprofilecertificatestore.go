package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateStore string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateStore_Machine AndroidDeviceOwnerScepCertificateProfileCertificateStore = "machine"
	AndroidDeviceOwnerScepCertificateProfileCertificateStore_User    AndroidDeviceOwnerScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateStore_Machine),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateStore_User),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateStore(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateStore{
		"machine": AndroidDeviceOwnerScepCertificateProfileCertificateStore_Machine,
		"user":    AndroidDeviceOwnerScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateStore(input)
	return &out, nil
}
