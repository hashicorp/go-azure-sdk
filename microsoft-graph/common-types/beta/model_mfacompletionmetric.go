package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MfaCompletionMetric struct {
	AppId         *string `json:"appId,omitempty"`
	AttemptsCount *int64  `json:"attemptsCount,omitempty"`
	FactDate      *string `json:"factDate,omitempty"`
	Id            *string `json:"id,omitempty"`
	MfaMethod     *string `json:"mfaMethod,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	Os            *string `json:"os,omitempty"`
	SuccessCount  *int64  `json:"successCount,omitempty"`
}
