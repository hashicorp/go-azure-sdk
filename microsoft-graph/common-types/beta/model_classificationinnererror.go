package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationInnerError struct {
	ActivityId      *string `json:"activityId,omitempty"`
	ClientRequestId *string `json:"clientRequestId,omitempty"`
	Code            *string `json:"code,omitempty"`
	ErrorDateTime   *string `json:"errorDateTime,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
