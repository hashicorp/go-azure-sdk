package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileHashHashType string

const (
	FileHashHashTypeauthenticodeHash256 FileHashHashType = "AuthenticodeHash256"
	FileHashHashTypectph                FileHashHashType = "Ctph"
	FileHashHashTypelsHash              FileHashHashType = "LsHash"
	FileHashHashTypemd5                 FileHashHashType = "Md5"
	FileHashHashTypesha1                FileHashHashType = "Sha1"
	FileHashHashTypesha256              FileHashHashType = "Sha256"
	FileHashHashTypeunknown             FileHashHashType = "Unknown"
)

func PossibleValuesForFileHashHashType() []string {
	return []string{
		string(FileHashHashTypeauthenticodeHash256),
		string(FileHashHashTypectph),
		string(FileHashHashTypelsHash),
		string(FileHashHashTypemd5),
		string(FileHashHashTypesha1),
		string(FileHashHashTypesha256),
		string(FileHashHashTypeunknown),
	}
}

func parseFileHashHashType(input string) (*FileHashHashType, error) {
	vals := map[string]FileHashHashType{
		"authenticodehash256": FileHashHashTypeauthenticodeHash256,
		"ctph":                FileHashHashTypectph,
		"lshash":              FileHashHashTypelsHash,
		"md5":                 FileHashHashTypemd5,
		"sha1":                FileHashHashTypesha1,
		"sha256":              FileHashHashTypesha256,
		"unknown":             FileHashHashTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileHashHashType(input)
	return &out, nil
}
