package securitymlanalyticssettings

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMLAnalyticsSetting interface {
	SecurityMLAnalyticsSetting() BaseSecurityMLAnalyticsSettingImpl
}

var _ SecurityMLAnalyticsSetting = BaseSecurityMLAnalyticsSettingImpl{}

type BaseSecurityMLAnalyticsSettingImpl struct {
	Etag       *string                         `json:"etag,omitempty"`
	Id         *string                         `json:"id,omitempty"`
	Kind       SecurityMLAnalyticsSettingsKind `json:"kind"`
	Name       *string                         `json:"name,omitempty"`
	SystemData *systemdata.SystemData          `json:"systemData,omitempty"`
	Type       *string                         `json:"type,omitempty"`
}

func (s BaseSecurityMLAnalyticsSettingImpl) SecurityMLAnalyticsSetting() BaseSecurityMLAnalyticsSettingImpl {
	return s
}

var _ SecurityMLAnalyticsSetting = RawSecurityMLAnalyticsSettingImpl{}

// RawSecurityMLAnalyticsSettingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityMLAnalyticsSettingImpl struct {
	securityMLAnalyticsSetting BaseSecurityMLAnalyticsSettingImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawSecurityMLAnalyticsSettingImpl) SecurityMLAnalyticsSetting() BaseSecurityMLAnalyticsSettingImpl {
	return s.securityMLAnalyticsSetting
}

func UnmarshalSecurityMLAnalyticsSettingImplementation(input []byte) (SecurityMLAnalyticsSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMLAnalyticsSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Anomaly") {
		var out AnomalySecurityMLAnalyticsSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AnomalySecurityMLAnalyticsSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityMLAnalyticsSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityMLAnalyticsSettingImpl: %+v", err)
	}

	return RawSecurityMLAnalyticsSettingImpl{
		securityMLAnalyticsSetting: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
