package datasets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompressionCodec string

const (
	CompressionCodecBzipTwo    CompressionCodec = "bzip2"
	CompressionCodecDeflate    CompressionCodec = "deflate"
	CompressionCodecGzip       CompressionCodec = "gzip"
	CompressionCodecLzFour     CompressionCodec = "lz4"
	CompressionCodecSnappy     CompressionCodec = "snappy"
	CompressionCodecTar        CompressionCodec = "tar"
	CompressionCodecTarGZip    CompressionCodec = "tarGZip"
	CompressionCodecZipDeflate CompressionCodec = "zipDeflate"
)

func PossibleValuesForCompressionCodec() []string {
	return []string{
		string(CompressionCodecBzipTwo),
		string(CompressionCodecDeflate),
		string(CompressionCodecGzip),
		string(CompressionCodecLzFour),
		string(CompressionCodecSnappy),
		string(CompressionCodecTar),
		string(CompressionCodecTarGZip),
		string(CompressionCodecZipDeflate),
	}
}

func (s *CompressionCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompressionCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompressionCodec(input string) (*CompressionCodec, error) {
	vals := map[string]CompressionCodec{
		"bzip2":      CompressionCodecBzipTwo,
		"deflate":    CompressionCodecDeflate,
		"gzip":       CompressionCodecGzip,
		"lz4":        CompressionCodecLzFour,
		"snappy":     CompressionCodecSnappy,
		"tar":        CompressionCodecTar,
		"targzip":    CompressionCodecTarGZip,
		"zipdeflate": CompressionCodecZipDeflate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompressionCodec(input)
	return &out, nil
}

type JsonFormatFilePattern string

const (
	JsonFormatFilePatternArrayOfObjects JsonFormatFilePattern = "arrayOfObjects"
	JsonFormatFilePatternSetOfObjects   JsonFormatFilePattern = "setOfObjects"
)

func PossibleValuesForJsonFormatFilePattern() []string {
	return []string{
		string(JsonFormatFilePatternArrayOfObjects),
		string(JsonFormatFilePatternSetOfObjects),
	}
}

func (s *JsonFormatFilePattern) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJsonFormatFilePattern(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJsonFormatFilePattern(input string) (*JsonFormatFilePattern, error) {
	vals := map[string]JsonFormatFilePattern{
		"arrayofobjects": JsonFormatFilePatternArrayOfObjects,
		"setofobjects":   JsonFormatFilePatternSetOfObjects,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JsonFormatFilePattern(input)
	return &out, nil
}

type OrcCompressionCodec string

const (
	OrcCompressionCodecLzo    OrcCompressionCodec = "lzo"
	OrcCompressionCodecNone   OrcCompressionCodec = "none"
	OrcCompressionCodecSnappy OrcCompressionCodec = "snappy"
	OrcCompressionCodecZlib   OrcCompressionCodec = "zlib"
)

func PossibleValuesForOrcCompressionCodec() []string {
	return []string{
		string(OrcCompressionCodecLzo),
		string(OrcCompressionCodecNone),
		string(OrcCompressionCodecSnappy),
		string(OrcCompressionCodecZlib),
	}
}

func (s *OrcCompressionCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrcCompressionCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrcCompressionCodec(input string) (*OrcCompressionCodec, error) {
	vals := map[string]OrcCompressionCodec{
		"lzo":    OrcCompressionCodecLzo,
		"none":   OrcCompressionCodecNone,
		"snappy": OrcCompressionCodecSnappy,
		"zlib":   OrcCompressionCodecZlib,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrcCompressionCodec(input)
	return &out, nil
}

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeObject),
		string(ParameterTypeSecureString),
		string(ParameterTypeString),
	}
}

func (s *ParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":        ParameterTypeArray,
		"bool":         ParameterTypeBool,
		"float":        ParameterTypeFloat,
		"int":          ParameterTypeInt,
		"object":       ParameterTypeObject,
		"securestring": ParameterTypeSecureString,
		"string":       ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}

type Type string

const (
	TypeLinkedServiceReference Type = "LinkedServiceReference"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLinkedServiceReference),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"linkedservicereference": TypeLinkedServiceReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
