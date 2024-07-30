package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatternedRecurrence struct {
	ODataType *string            `json:"@odata.type,omitempty"`
	Pattern   *RecurrencePattern `json:"pattern,omitempty"`
	Range     *RecurrenceRange   `json:"range,omitempty"`
}
