package datasetmapping

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetMapping interface {
	DataSetMapping() BaseDataSetMappingImpl
}

var _ DataSetMapping = BaseDataSetMappingImpl{}

type BaseDataSetMappingImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       DataSetMappingKind     `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseDataSetMappingImpl) DataSetMapping() BaseDataSetMappingImpl {
	return s
}

var _ DataSetMapping = RawDataSetMappingImpl{}

// RawDataSetMappingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataSetMappingImpl struct {
	dataSetMapping BaseDataSetMappingImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawDataSetMappingImpl) DataSetMapping() BaseDataSetMappingImpl {
	return s.dataSetMapping
}

func UnmarshalDataSetMappingImplementation(input []byte) (DataSetMapping, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataSetMapping into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AdlsGen2File") {
		var out ADLSGen2FileDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ADLSGen2FileDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AdlsGen2FileSystem") {
		var out ADLSGen2FileSystemDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ADLSGen2FileSystemDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AdlsGen2Folder") {
		var out ADLSGen2FolderDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ADLSGen2FolderDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Container") {
		var out BlobContainerDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobContainerDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Blob") {
		var out BlobDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "BlobFolder") {
		var out BlobFolderDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobFolderDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KustoCluster") {
		var out KustoClusterDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KustoClusterDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KustoDatabase") {
		var out KustoDatabaseDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KustoDatabaseDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KustoTable") {
		var out KustoTableDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KustoTableDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SqlDBTable") {
		var out SqlDBTableDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SqlDBTableDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SqlDWTable") {
		var out SqlDWTableDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SqlDWTableDataSetMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SynapseWorkspaceSqlPoolTable") {
		var out SynapseWorkspaceSqlPoolTableDataSetMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SynapseWorkspaceSqlPoolTableDataSetMapping: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataSetMappingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataSetMappingImpl: %+v", err)
	}

	return RawDataSetMappingImpl{
		dataSetMapping: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
