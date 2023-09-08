package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileCertificateStore string

const (
	Windows10XSCEPCertificateProfileCertificateStoremachine Windows10XSCEPCertificateProfileCertificateStore = "Machine"
	Windows10XSCEPCertificateProfileCertificateStoreuser    Windows10XSCEPCertificateProfileCertificateStore = "User"
)

func PossibleValuesForWindows10XSCEPCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileCertificateStoremachine),
		string(Windows10XSCEPCertificateProfileCertificateStoreuser),
	}
}

func parseWindows10XSCEPCertificateProfileCertificateStore(input string) (*Windows10XSCEPCertificateProfileCertificateStore, error) {
	vals := map[string]Windows10XSCEPCertificateProfileCertificateStore{
		"machine": Windows10XSCEPCertificateProfileCertificateStoremachine,
		"user":    Windows10XSCEPCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileCertificateStore(input)
	return &out, nil
}
