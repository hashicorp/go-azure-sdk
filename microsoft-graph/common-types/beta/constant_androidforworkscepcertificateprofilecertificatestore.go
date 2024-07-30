package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileCertificateStore string

const (
	AndroidForWorkScepCertificateProfileCertificateStore_Machine AndroidForWorkScepCertificateProfileCertificateStore = "machine"
	AndroidForWorkScepCertificateProfileCertificateStore_User    AndroidForWorkScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileCertificateStore_Machine),
		string(AndroidForWorkScepCertificateProfileCertificateStore_User),
	}
}

func (s *AndroidForWorkScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileCertificateStore(input string) (*AndroidForWorkScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileCertificateStore{
		"machine": AndroidForWorkScepCertificateProfileCertificateStore_Machine,
		"user":    AndroidForWorkScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileCertificateStore(input)
	return &out, nil
}
