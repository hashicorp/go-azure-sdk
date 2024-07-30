package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConvertIdResult struct {
	ErrorDetails *GenericError `json:"errorDetails,omitempty"`
	ODataType    *string       `json:"@odata.type,omitempty"`
	SourceId     *string       `json:"sourceId,omitempty"`
	TargetId     *string       `json:"targetId,omitempty"`
}
