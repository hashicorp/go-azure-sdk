package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileCertificateStore string

const (
	IosPkcsCertificateProfileCertificateStoremachine IosPkcsCertificateProfileCertificateStore = "Machine"
	IosPkcsCertificateProfileCertificateStoreuser    IosPkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForIosPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(IosPkcsCertificateProfileCertificateStoremachine),
		string(IosPkcsCertificateProfileCertificateStoreuser),
	}
}

func parseIosPkcsCertificateProfileCertificateStore(input string) (*IosPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]IosPkcsCertificateProfileCertificateStore{
		"machine": IosPkcsCertificateProfileCertificateStoremachine,
		"user":    IosPkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
