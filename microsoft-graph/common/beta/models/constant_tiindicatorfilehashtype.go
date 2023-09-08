package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorFileHashType string

const (
	TiIndicatorFileHashTypeauthenticodeHash256 TiIndicatorFileHashType = "AuthenticodeHash256"
	TiIndicatorFileHashTypectph                TiIndicatorFileHashType = "Ctph"
	TiIndicatorFileHashTypelsHash              TiIndicatorFileHashType = "LsHash"
	TiIndicatorFileHashTypemd5                 TiIndicatorFileHashType = "Md5"
	TiIndicatorFileHashTypesha1                TiIndicatorFileHashType = "Sha1"
	TiIndicatorFileHashTypesha256              TiIndicatorFileHashType = "Sha256"
	TiIndicatorFileHashTypeunknown             TiIndicatorFileHashType = "Unknown"
)

func PossibleValuesForTiIndicatorFileHashType() []string {
	return []string{
		string(TiIndicatorFileHashTypeauthenticodeHash256),
		string(TiIndicatorFileHashTypectph),
		string(TiIndicatorFileHashTypelsHash),
		string(TiIndicatorFileHashTypemd5),
		string(TiIndicatorFileHashTypesha1),
		string(TiIndicatorFileHashTypesha256),
		string(TiIndicatorFileHashTypeunknown),
	}
}

func parseTiIndicatorFileHashType(input string) (*TiIndicatorFileHashType, error) {
	vals := map[string]TiIndicatorFileHashType{
		"authenticodehash256": TiIndicatorFileHashTypeauthenticodeHash256,
		"ctph":                TiIndicatorFileHashTypectph,
		"lshash":              TiIndicatorFileHashTypelsHash,
		"md5":                 TiIndicatorFileHashTypemd5,
		"sha1":                TiIndicatorFileHashTypesha1,
		"sha256":              TiIndicatorFileHashTypesha256,
		"unknown":             TiIndicatorFileHashTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorFileHashType(input)
	return &out, nil
}
