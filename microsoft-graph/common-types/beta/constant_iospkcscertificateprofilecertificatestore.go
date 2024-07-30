package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileCertificateStore string

const (
	IosPkcsCertificateProfileCertificateStore_Machine IosPkcsCertificateProfileCertificateStore = "machine"
	IosPkcsCertificateProfileCertificateStore_User    IosPkcsCertificateProfileCertificateStore = "user"
)

func PossibleValuesForIosPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(IosPkcsCertificateProfileCertificateStore_Machine),
		string(IosPkcsCertificateProfileCertificateStore_User),
	}
}

func (s *IosPkcsCertificateProfileCertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosPkcsCertificateProfileCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosPkcsCertificateProfileCertificateStore(input string) (*IosPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]IosPkcsCertificateProfileCertificateStore{
		"machine": IosPkcsCertificateProfileCertificateStore_Machine,
		"user":    IosPkcsCertificateProfileCertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
