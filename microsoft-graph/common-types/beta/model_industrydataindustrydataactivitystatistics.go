package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityStatistics struct {
	ActivityId  *string                                           `json:"activityId,omitempty"`
	DisplayName *string                                           `json:"displayName,omitempty"`
	ODataType   *string                                           `json:"@odata.type,omitempty"`
	Status      *IndustryDataIndustryDataActivityStatisticsStatus `json:"status,omitempty"`
}
