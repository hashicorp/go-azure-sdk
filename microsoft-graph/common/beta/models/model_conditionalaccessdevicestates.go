package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessDeviceStates struct {
	ExcludeStates *[]string `json:"excludeStates,omitempty"`
	IncludeStates *[]string `json:"includeStates,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
}
