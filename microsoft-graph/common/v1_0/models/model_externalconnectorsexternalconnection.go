package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnection struct {
	ActivitySettings *ExternalConnectorsActivitySettings        `json:"activitySettings,omitempty"`
	Configuration    *ExternalConnectorsConfiguration           `json:"configuration,omitempty"`
	Description      *string                                    `json:"description,omitempty"`
	Groups           *[]ExternalConnectorsExternalGroup         `json:"groups,omitempty"`
	Id               *string                                    `json:"id,omitempty"`
	Items            *[]ExternalConnectorsExternalItem          `json:"items,omitempty"`
	Name             *string                                    `json:"name,omitempty"`
	ODataType        *string                                    `json:"@odata.type,omitempty"`
	Operations       *[]ExternalConnectorsConnectionOperation   `json:"operations,omitempty"`
	Schema           *ExternalConnectorsSchema                  `json:"schema,omitempty"`
	SearchSettings   *ExternalConnectorsSearchSettings          `json:"searchSettings,omitempty"`
	State            *ExternalConnectorsExternalConnectionState `json:"state,omitempty"`
}
