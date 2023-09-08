package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileCertificateStore string

const (
	Windows10PkcsCertificateProfileCertificateStoremachine Windows10PkcsCertificateProfileCertificateStore = "Machine"
	Windows10PkcsCertificateProfileCertificateStoreuser    Windows10PkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForWindows10PkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows10PkcsCertificateProfileCertificateStoremachine),
		string(Windows10PkcsCertificateProfileCertificateStoreuser),
	}
}

func parseWindows10PkcsCertificateProfileCertificateStore(input string) (*Windows10PkcsCertificateProfileCertificateStore, error) {
	vals := map[string]Windows10PkcsCertificateProfileCertificateStore{
		"machine": Windows10PkcsCertificateProfileCertificateStoremachine,
		"user":    Windows10PkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
