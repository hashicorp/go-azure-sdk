package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalRecommendationConfiguration = LogAnalyticsQueryRecommendationConfiguration{}

type LogAnalyticsQueryRecommendationConfiguration struct {
	QueryText                           string    `json:"queryText"`
	RequiredDiagnosticSettingCategories *[]string `json:"requiredDiagnosticSettingCategories,omitempty"`
	RequiredTables                      *[]string `json:"requiredTables,omitempty"`
	TimeGrain                           *string   `json:"timeGrain,omitempty"`
	ValueColumnName                     *string   `json:"valueColumnName,omitempty"`

	// Fields inherited from SignalRecommendationConfiguration

	SignalKind SignalRecommendationKind `json:"signalKind"`
}

func (s LogAnalyticsQueryRecommendationConfiguration) SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl {
	return BaseSignalRecommendationConfigurationImpl{
		SignalKind: s.SignalKind,
	}
}

var _ json.Marshaler = LogAnalyticsQueryRecommendationConfiguration{}

func (s LogAnalyticsQueryRecommendationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper LogAnalyticsQueryRecommendationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LogAnalyticsQueryRecommendationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LogAnalyticsQueryRecommendationConfiguration: %+v", err)
	}

	decoded["signalKind"] = "LogAnalyticsQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LogAnalyticsQueryRecommendationConfiguration: %+v", err)
	}

	return encoded, nil
}
