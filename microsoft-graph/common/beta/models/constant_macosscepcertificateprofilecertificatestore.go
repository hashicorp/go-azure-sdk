package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileCertificateStore string

const (
	MacOSScepCertificateProfileCertificateStoremachine MacOSScepCertificateProfileCertificateStore = "Machine"
	MacOSScepCertificateProfileCertificateStoreuser    MacOSScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForMacOSScepCertificateProfileCertificateStore() []string {
	return []string{
		string(MacOSScepCertificateProfileCertificateStoremachine),
		string(MacOSScepCertificateProfileCertificateStoreuser),
	}
}

func parseMacOSScepCertificateProfileCertificateStore(input string) (*MacOSScepCertificateProfileCertificateStore, error) {
	vals := map[string]MacOSScepCertificateProfileCertificateStore{
		"machine": MacOSScepCertificateProfileCertificateStoremachine,
		"user":    MacOSScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileCertificateStore(input)
	return &out, nil
}
