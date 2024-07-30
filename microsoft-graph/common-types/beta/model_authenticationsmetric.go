package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationsMetric struct {
	Appid         *string `json:"appid,omitempty"`
	AttemptsCount *int64  `json:"attemptsCount,omitempty"`
	Country       *string `json:"country,omitempty"`
	FactDate      *string `json:"factDate,omitempty"`
	Id            *string `json:"id,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	Os            *string `json:"os,omitempty"`
	SuccessCount  *int64  `json:"successCount,omitempty"`
}
