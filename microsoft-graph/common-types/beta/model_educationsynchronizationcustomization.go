package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationCustomization struct {
	AllowDisplayNameUpdate   *bool     `json:"allowDisplayNameUpdate,omitempty"`
	IsSyncDeferred           *bool     `json:"isSyncDeferred,omitempty"`
	ODataType                *string   `json:"@odata.type,omitempty"`
	OptionalPropertiesToSync *[]string `json:"optionalPropertiesToSync,omitempty"`
	SynchronizationStartDate *string   `json:"synchronizationStartDate,omitempty"`
}
