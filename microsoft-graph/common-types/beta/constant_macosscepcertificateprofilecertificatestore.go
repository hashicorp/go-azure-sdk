package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileCertificateStore string

const (
	MacOSScepCertificateProfileCertificateStore_Machine MacOSScepCertificateProfileCertificateStore = "machine"
	MacOSScepCertificateProfileCertificateStore_User    MacOSScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForMacOSScepCertificateProfileCertificateStore() []string {
	return []string{
		string(MacOSScepCertificateProfileCertificateStore_Machine),
		string(MacOSScepCertificateProfileCertificateStore_User),
	}
}

func (s *MacOSScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileCertificateStore(input string) (*MacOSScepCertificateProfileCertificateStore, error) {
	vals := map[string]MacOSScepCertificateProfileCertificateStore{
		"machine": MacOSScepCertificateProfileCertificateStore_Machine,
		"user":    MacOSScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileCertificateStore(input)
	return &out, nil
}
