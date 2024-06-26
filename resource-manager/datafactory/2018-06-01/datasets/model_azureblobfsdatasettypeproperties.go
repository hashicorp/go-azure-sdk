package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBlobFSDatasetTypeProperties struct {
	Compression *DatasetCompression  `json:"compression,omitempty"`
	FileName    *string              `json:"fileName,omitempty"`
	FolderPath  *string              `json:"folderPath,omitempty"`
	Format      DatasetStorageFormat `json:"format"`
}

var _ json.Unmarshaler = &AzureBlobFSDatasetTypeProperties{}

func (s *AzureBlobFSDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AzureBlobFSDatasetTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AzureBlobFSDatasetTypeProperties: %+v", err)
	}

	s.Compression = decoded.Compression
	s.FileName = decoded.FileName
	s.FolderPath = decoded.FolderPath

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureBlobFSDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["format"]; ok {
		impl, err := unmarshalDatasetStorageFormatImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Format' for 'AzureBlobFSDatasetTypeProperties': %+v", err)
		}
		s.Format = impl
	}
	return nil
}
