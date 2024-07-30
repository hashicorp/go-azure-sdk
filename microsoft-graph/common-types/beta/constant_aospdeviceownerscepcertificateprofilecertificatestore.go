package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileCertificateStore string

const (
	AospDeviceOwnerScepCertificateProfileCertificateStore_Machine AospDeviceOwnerScepCertificateProfileCertificateStore = "machine"
	AospDeviceOwnerScepCertificateProfileCertificateStore_User    AospDeviceOwnerScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileCertificateStore_Machine),
		string(AospDeviceOwnerScepCertificateProfileCertificateStore_User),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileCertificateStore(input string) (*AospDeviceOwnerScepCertificateProfileCertificateStore, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileCertificateStore{
		"machine": AospDeviceOwnerScepCertificateProfileCertificateStore_Machine,
		"user":    AospDeviceOwnerScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileCertificateStore(input)
	return &out, nil
}
