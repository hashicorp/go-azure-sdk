package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestApiPollerDataConnectorProperties struct {
	AddOnAttributes         *map[string]string                `json:"addOnAttributes,omitempty"`
	Auth                    CcpAuthConfig                     `json:"auth"`
	ConnectorDefinitionName string                            `json:"connectorDefinitionName"`
	DataType                *string                           `json:"dataType,omitempty"`
	DcrConfig               *DCRConfiguration                 `json:"dcrConfig,omitempty"`
	IsActive                *bool                             `json:"isActive,omitempty"`
	Paging                  *RestApiPollerRequestPagingConfig `json:"paging,omitempty"`
	Request                 RestApiPollerRequestConfig        `json:"request"`
	Response                *CcpResponseConfig                `json:"response,omitempty"`
}

var _ json.Unmarshaler = &RestApiPollerDataConnectorProperties{}

func (s *RestApiPollerDataConnectorProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddOnAttributes         *map[string]string                `json:"addOnAttributes,omitempty"`
		ConnectorDefinitionName string                            `json:"connectorDefinitionName"`
		DataType                *string                           `json:"dataType,omitempty"`
		DcrConfig               *DCRConfiguration                 `json:"dcrConfig,omitempty"`
		IsActive                *bool                             `json:"isActive,omitempty"`
		Paging                  *RestApiPollerRequestPagingConfig `json:"paging,omitempty"`
		Request                 RestApiPollerRequestConfig        `json:"request"`
		Response                *CcpResponseConfig                `json:"response,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddOnAttributes = decoded.AddOnAttributes
	s.ConnectorDefinitionName = decoded.ConnectorDefinitionName
	s.DataType = decoded.DataType
	s.DcrConfig = decoded.DcrConfig
	s.IsActive = decoded.IsActive
	s.Paging = decoded.Paging
	s.Request = decoded.Request
	s.Response = decoded.Response

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RestApiPollerDataConnectorProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["auth"]; ok {
		impl, err := UnmarshalCcpAuthConfigImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Auth' for 'RestApiPollerDataConnectorProperties': %+v", err)
		}
		s.Auth = impl
	}

	return nil
}
