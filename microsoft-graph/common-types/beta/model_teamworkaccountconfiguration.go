package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAccountConfiguration struct {
	ODataType                           *string                                      `json:"@odata.type,omitempty"`
	OnPremisesCalendarSyncConfiguration *TeamworkOnPremisesCalendarSyncConfiguration `json:"onPremisesCalendarSyncConfiguration,omitempty"`
	SupportedClient                     *TeamworkAccountConfigurationSupportedClient `json:"supportedClient,omitempty"`
}
