package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileCertificateStore string

const (
	IosScepCertificateProfileCertificateStore_Machine IosScepCertificateProfileCertificateStore = "machine"
	IosScepCertificateProfileCertificateStore_User    IosScepCertificateProfileCertificateStore = "user"
)

func PossibleValuesForIosScepCertificateProfileCertificateStore() []string {
	return []string{
		string(IosScepCertificateProfileCertificateStore_Machine),
		string(IosScepCertificateProfileCertificateStore_User),
	}
}

func (s *IosScepCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileCertificateStore(input string) (*IosScepCertificateProfileCertificateStore, error) {
	vals := map[string]IosScepCertificateProfileCertificateStore{
		"machine": IosScepCertificateProfileCertificateStore_Machine,
		"user":    IosScepCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileCertificateStore(input)
	return &out, nil
}
