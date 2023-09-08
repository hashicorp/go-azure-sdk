package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileCertificateStore string

const (
	Windows81SCEPCertificateProfileCertificateStoremachine Windows81SCEPCertificateProfileCertificateStore = "Machine"
	Windows81SCEPCertificateProfileCertificateStoreuser    Windows81SCEPCertificateProfileCertificateStore = "User"
)

func PossibleValuesForWindows81SCEPCertificateProfileCertificateStore() []string {
	return []string{
		string(Windows81SCEPCertificateProfileCertificateStoremachine),
		string(Windows81SCEPCertificateProfileCertificateStoreuser),
	}
}

func parseWindows81SCEPCertificateProfileCertificateStore(input string) (*Windows81SCEPCertificateProfileCertificateStore, error) {
	vals := map[string]Windows81SCEPCertificateProfileCertificateStore{
		"machine": Windows81SCEPCertificateProfileCertificateStoremachine,
		"user":    Windows81SCEPCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileCertificateStore(input)
	return &out, nil
}
