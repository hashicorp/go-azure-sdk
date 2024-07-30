package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileHashHashType string

const (
	FileHashHashType_AuthenticodeHash256 FileHashHashType = "authenticodeHash256"
	FileHashHashType_Ctph                FileHashHashType = "ctph"
	FileHashHashType_LsHash              FileHashHashType = "lsHash"
	FileHashHashType_Md5                 FileHashHashType = "md5"
	FileHashHashType_Sha1                FileHashHashType = "sha1"
	FileHashHashType_Sha256              FileHashHashType = "sha256"
	FileHashHashType_Unknown             FileHashHashType = "unknown"
)

func PossibleValuesForFileHashHashType() []string {
	return []string{
		string(FileHashHashType_AuthenticodeHash256),
		string(FileHashHashType_Ctph),
		string(FileHashHashType_LsHash),
		string(FileHashHashType_Md5),
		string(FileHashHashType_Sha1),
		string(FileHashHashType_Sha256),
		string(FileHashHashType_Unknown),
	}
}

func (s *FileHashHashType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileHashHashType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileHashHashType(input string) (*FileHashHashType, error) {
	vals := map[string]FileHashHashType{
		"authenticodehash256": FileHashHashType_AuthenticodeHash256,
		"ctph":                FileHashHashType_Ctph,
		"lshash":              FileHashHashType_LsHash,
		"md5":                 FileHashHashType_Md5,
		"sha1":                FileHashHashType_Sha1,
		"sha256":              FileHashHashType_Sha256,
		"unknown":             FileHashHashType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileHashHashType(input)
	return &out, nil
}
