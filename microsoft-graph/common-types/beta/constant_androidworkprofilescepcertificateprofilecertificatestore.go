package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileCertificateStore string

const (
	AndroidWorkProfileScepCertificateProfileCertificateStore_Machine AndroidWorkProfileScepCertificateProfileCertificateStore = "machine"
	AndroidWorkProfileScepCertificateProfileCertificateStore_User    AndroidWorkProfileScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileCertificateStore_Machine),
		string(AndroidWorkProfileScepCertificateProfileCertificateStore_User),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileCertificateStore(input string) (*AndroidWorkProfileScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileCertificateStore{
		"machine": AndroidWorkProfileScepCertificateProfileCertificateStore_Machine,
		"user":    AndroidWorkProfileScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileCertificateStore(input)
	return &out, nil
}
