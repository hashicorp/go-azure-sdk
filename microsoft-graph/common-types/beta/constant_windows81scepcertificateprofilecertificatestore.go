package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileCertificateStore string

const (
	Windows81SCEPCertificateProfileCertificateStore_Machine Windows81SCEPCertificateProfileCertificateStore = "machine"
	Windows81SCEPCertificateProfileCertificateStore_User    Windows81SCEPCertificateProfileCertificateStore = "user"
)

func PossibleValuesForWindows81SCEPCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows81SCEPCertificateProfileCertificateStore_Machine),
		string(Windows81SCEPCertificateProfileCertificateStore_User),
	}
}

func (s *Windows81SCEPCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileCertificateStore(input string) (*Windows81SCEPCertificateProfileCertificateStore, error) {
	vals := map[string]Windows81SCEPCertificateProfileCertificateStore{
		"machine": Windows81SCEPCertificateProfileCertificateStore_Machine,
		"user":    Windows81SCEPCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileCertificateStore(input)
	return &out, nil
}
