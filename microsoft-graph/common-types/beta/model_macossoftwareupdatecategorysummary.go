package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateCategorySummary struct {
	DeviceId              *string                                           `json:"deviceId,omitempty"`
	DisplayName           *string                                           `json:"displayName,omitempty"`
	FailedUpdateCount     *int64                                            `json:"failedUpdateCount,omitempty"`
	Id                    *string                                           `json:"id,omitempty"`
	LastUpdatedDateTime   *string                                           `json:"lastUpdatedDateTime,omitempty"`
	ODataType             *string                                           `json:"@odata.type,omitempty"`
	SuccessfulUpdateCount *int64                                            `json:"successfulUpdateCount,omitempty"`
	TotalUpdateCount      *int64                                            `json:"totalUpdateCount,omitempty"`
	UpdateCategory        *MacOSSoftwareUpdateCategorySummaryUpdateCategory `json:"updateCategory,omitempty"`
	UpdateStateSummaries  *[]MacOSSoftwareUpdateStateSummary                `json:"updateStateSummaries,omitempty"`
	UserId                *string                                           `json:"userId,omitempty"`
}
