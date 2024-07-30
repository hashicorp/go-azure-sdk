package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificateStore string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificateStore_Machine AospDeviceOwnerPkcsCertificateProfileCertificateStore = "machine"
	AospDeviceOwnerPkcsCertificateProfileCertificateStore_User    AospDeviceOwnerPkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificateStore_Machine),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateStore_User),
	}
}

func (s *AospDeviceOwnerPkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerPkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificateStore(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificateStore{
		"machine": AospDeviceOwnerPkcsCertificateProfileCertificateStore_Machine,
		"user":    AospDeviceOwnerPkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
