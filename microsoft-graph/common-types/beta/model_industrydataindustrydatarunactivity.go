package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunActivity struct {
	Activity      *IndustryDataIndustryDataActivity          `json:"activity,omitempty"`
	BlockingError *PublicError                               `json:"blockingError,omitempty"`
	DisplayName   *string                                    `json:"displayName,omitempty"`
	Id            *string                                    `json:"id,omitempty"`
	ODataType     *string                                    `json:"@odata.type,omitempty"`
	Status        *IndustryDataIndustryDataRunActivityStatus `json:"status,omitempty"`
}
