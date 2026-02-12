package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBlobFSDatasetTypeProperties struct {
	Compression DatasetCompression   `json:"compression"`
	FileName    *interface{}         `json:"fileName,omitempty"`
	FolderPath  *interface{}         `json:"folderPath,omitempty"`
	Format      DatasetStorageFormat `json:"format"`
}

var _ json.Unmarshaler = &AzureBlobFSDatasetTypeProperties{}

func (s *AzureBlobFSDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FileName   *interface{} `json:"fileName,omitempty"`
		FolderPath *interface{} `json:"folderPath,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FileName = decoded.FileName
	s.FolderPath = decoded.FolderPath

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureBlobFSDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compression"]; ok {
		impl, err := UnmarshalDatasetCompressionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Compression' for 'AzureBlobFSDatasetTypeProperties': %+v", err)
		}
		s.Compression = impl
	}

	if v, ok := temp["format"]; ok {
		impl, err := UnmarshalDatasetStorageFormatImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Format' for 'AzureBlobFSDatasetTypeProperties': %+v", err)
		}
		s.Format = impl
	}

	return nil
}
