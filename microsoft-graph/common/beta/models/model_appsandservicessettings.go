package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsAndServicesSettings struct {
	IsAppAndServicesTrialEnabled *bool   `json:"isAppAndServicesTrialEnabled,omitempty"`
	IsOfficeStoreEnabled         *bool   `json:"isOfficeStoreEnabled,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
}
