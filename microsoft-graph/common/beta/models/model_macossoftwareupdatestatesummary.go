package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateStateSummary struct {
	DisplayName         *string                                        `json:"displayName,omitempty"`
	Id                  *string                                        `json:"id,omitempty"`
	LastUpdatedDateTime *string                                        `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                        `json:"@odata.type,omitempty"`
	ProductKey          *string                                        `json:"productKey,omitempty"`
	State               *MacOSSoftwareUpdateStateSummaryState          `json:"state,omitempty"`
	UpdateCategory      *MacOSSoftwareUpdateStateSummaryUpdateCategory `json:"updateCategory,omitempty"`
	UpdateVersion       *string                                        `json:"updateVersion,omitempty"`
}
