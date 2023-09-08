package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnection struct {
	Configuration *Configuration           `json:"configuration,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	Groups        *[]ExternalGroup         `json:"groups,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	Items         *[]ExternalItem          `json:"items,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	ODataType     *string                  `json:"@odata.type,omitempty"`
	Operations    *[]ConnectionOperation   `json:"operations,omitempty"`
	Schema        *Schema                  `json:"schema,omitempty"`
	State         *ExternalConnectionState `json:"state,omitempty"`
}
