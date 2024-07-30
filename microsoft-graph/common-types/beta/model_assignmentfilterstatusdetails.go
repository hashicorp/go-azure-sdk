package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterStatusDetails struct {
	DeviceProperties   *[]KeyValuePair                      `json:"deviceProperties,omitempty"`
	EvalutionSummaries *[]AssignmentFilterEvaluationSummary `json:"evalutionSummaries,omitempty"`
	ManagedDeviceId    *string                              `json:"managedDeviceId,omitempty"`
	ODataType          *string                              `json:"@odata.type,omitempty"`
	PayloadId          *string                              `json:"payloadId,omitempty"`
	UserId             *string                              `json:"userId,omitempty"`
}
