package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TileConfiguration interface {
	TileConfiguration() BaseTileConfigurationImpl
}

var _ TileConfiguration = BaseTileConfigurationImpl{}

type BaseTileConfigurationImpl struct {
	Type string `json:"type"`
}

func (s BaseTileConfigurationImpl) TileConfiguration() BaseTileConfigurationImpl {
	return s
}

var _ TileConfiguration = RawTileConfigurationImpl{}

// RawTileConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTileConfigurationImpl struct {
	tileConfiguration BaseTileConfigurationImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawTileConfigurationImpl) TileConfiguration() BaseTileConfigurationImpl {
	return s.tileConfiguration
}

func UnmarshalTileConfigurationImplementation(input []byte) (TileConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TileConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "barChart") {
		var out BarChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BarChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "command") {
		var out CommandConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "command") {
		var out CommandTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "dataExplorer") {
		var out DataExplorerTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataExplorerTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "deviceCount") {
		var out DeviceCountTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceCountTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "eventChart") {
		var out EventChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "eventHistoryChart") {
		var out EventHistoryChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventHistoryChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "externalContent") {
		var out ExternalContentTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalContentTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "heatMapChart") {
		var out HeatMapConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HeatMapConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "image") {
		var out ImageTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "kpi") {
		var out KpiTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KpiTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "label") {
		var out LabelTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LabelTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "lineChart") {
		var out LineChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LineChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "lkv") {
		var out LkvTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LkvTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "mapProperty") {
		var out MapPropertyConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MapPropertyConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "mapTelemetry") {
		var out MapTelemetryConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MapTelemetryConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "markdown") {
		var out MarkdownTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MarkdownTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "pieChart") {
		var out PieChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PieChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "property") {
		var out PropertyTileConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PropertyTileConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "stateChart") {
		var out StateChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StateChartConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "stateHistoryChart") {
		var out StateHistoryChartConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StateHistoryChartConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseTileConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTileConfigurationImpl: %+v", err)
	}

	return RawTileConfigurationImpl{
		tileConfiguration: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
