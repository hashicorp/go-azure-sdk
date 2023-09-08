package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileCertificateStore string

const (
	MacOSPkcsCertificateProfileCertificateStoremachine MacOSPkcsCertificateProfileCertificateStore = "Machine"
	MacOSPkcsCertificateProfileCertificateStoreuser    MacOSPkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForMacOSPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(MacOSPkcsCertificateProfileCertificateStoremachine),
		string(MacOSPkcsCertificateProfileCertificateStoreuser),
	}
}

func parseMacOSPkcsCertificateProfileCertificateStore(input string) (*MacOSPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]MacOSPkcsCertificateProfileCertificateStore{
		"machine": MacOSPkcsCertificateProfileCertificateStoremachine,
		"user":    MacOSPkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
