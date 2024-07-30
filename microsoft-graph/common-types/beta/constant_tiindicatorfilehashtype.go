package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorFileHashType string

const (
	TiIndicatorFileHashType_AuthenticodeHash256 TiIndicatorFileHashType = "authenticodeHash256"
	TiIndicatorFileHashType_Ctph                TiIndicatorFileHashType = "ctph"
	TiIndicatorFileHashType_LsHash              TiIndicatorFileHashType = "lsHash"
	TiIndicatorFileHashType_Md5                 TiIndicatorFileHashType = "md5"
	TiIndicatorFileHashType_Sha1                TiIndicatorFileHashType = "sha1"
	TiIndicatorFileHashType_Sha256              TiIndicatorFileHashType = "sha256"
	TiIndicatorFileHashType_Unknown             TiIndicatorFileHashType = "unknown"
)

func PossibleValuesForTiIndicatorFileHashType() []string {
	return []string{
		string(TiIndicatorFileHashType_AuthenticodeHash256),
		string(TiIndicatorFileHashType_Ctph),
		string(TiIndicatorFileHashType_LsHash),
		string(TiIndicatorFileHashType_Md5),
		string(TiIndicatorFileHashType_Sha1),
		string(TiIndicatorFileHashType_Sha256),
		string(TiIndicatorFileHashType_Unknown),
	}
}

func (s *TiIndicatorFileHashType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiIndicatorFileHashType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiIndicatorFileHashType(input string) (*TiIndicatorFileHashType, error) {
	vals := map[string]TiIndicatorFileHashType{
		"authenticodehash256": TiIndicatorFileHashType_AuthenticodeHash256,
		"ctph":                TiIndicatorFileHashType_Ctph,
		"lshash":              TiIndicatorFileHashType_LsHash,
		"md5":                 TiIndicatorFileHashType_Md5,
		"sha1":                TiIndicatorFileHashType_Sha1,
		"sha256":              TiIndicatorFileHashType_Sha256,
		"unknown":             TiIndicatorFileHashType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorFileHashType(input)
	return &out, nil
}
