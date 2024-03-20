package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileShareDatasetTypeProperties struct {
	Compression           *DatasetCompression  `json:"compression,omitempty"`
	FileFilter            *interface{}         `json:"fileFilter,omitempty"`
	FileName              *interface{}         `json:"fileName,omitempty"`
	FolderPath            *interface{}         `json:"folderPath,omitempty"`
	Format                DatasetStorageFormat `json:"format"`
	ModifiedDatetimeEnd   *interface{}         `json:"modifiedDatetimeEnd,omitempty"`
	ModifiedDatetimeStart *interface{}         `json:"modifiedDatetimeStart,omitempty"`
}

var _ json.Unmarshaler = &FileShareDatasetTypeProperties{}

func (s *FileShareDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias FileShareDatasetTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into FileShareDatasetTypeProperties: %+v", err)
	}

	s.Compression = decoded.Compression
	s.FileFilter = decoded.FileFilter
	s.FileName = decoded.FileName
	s.FolderPath = decoded.FolderPath
	s.ModifiedDatetimeEnd = decoded.ModifiedDatetimeEnd
	s.ModifiedDatetimeStart = decoded.ModifiedDatetimeStart

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FileShareDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["format"]; ok {
		impl, err := unmarshalDatasetStorageFormatImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Format' for 'FileShareDatasetTypeProperties': %+v", err)
		}
		s.Format = impl
	}
	return nil
}
