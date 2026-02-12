package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmazonS3DatasetTypeProperties struct {
	BucketName            interface{}          `json:"bucketName"`
	Compression           DatasetCompression   `json:"compression"`
	Format                DatasetStorageFormat `json:"format"`
	Key                   *interface{}         `json:"key,omitempty"`
	ModifiedDatetimeEnd   *interface{}         `json:"modifiedDatetimeEnd,omitempty"`
	ModifiedDatetimeStart *interface{}         `json:"modifiedDatetimeStart,omitempty"`
	Prefix                *interface{}         `json:"prefix,omitempty"`
	Version               *interface{}         `json:"version,omitempty"`
}

var _ json.Unmarshaler = &AmazonS3DatasetTypeProperties{}

func (s *AmazonS3DatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BucketName            interface{}  `json:"bucketName"`
		Key                   *interface{} `json:"key,omitempty"`
		ModifiedDatetimeEnd   *interface{} `json:"modifiedDatetimeEnd,omitempty"`
		ModifiedDatetimeStart *interface{} `json:"modifiedDatetimeStart,omitempty"`
		Prefix                *interface{} `json:"prefix,omitempty"`
		Version               *interface{} `json:"version,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BucketName = decoded.BucketName
	s.Key = decoded.Key
	s.ModifiedDatetimeEnd = decoded.ModifiedDatetimeEnd
	s.ModifiedDatetimeStart = decoded.ModifiedDatetimeStart
	s.Prefix = decoded.Prefix
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AmazonS3DatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compression"]; ok {
		impl, err := UnmarshalDatasetCompressionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Compression' for 'AmazonS3DatasetTypeProperties': %+v", err)
		}
		s.Compression = impl
	}

	if v, ok := temp["format"]; ok {
		impl, err := UnmarshalDatasetStorageFormatImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Format' for 'AmazonS3DatasetTypeProperties': %+v", err)
		}
		s.Format = impl
	}

	return nil
}
