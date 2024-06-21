package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointProvisioningStepStatus struct {
	AdditionalInformation *map[string]string `json:"additionalInformation,omitempty"`
	EndTime               *string            `json:"endTime,omitempty"`
	ErrorCode             *int64             `json:"errorCode,omitempty"`
	MinutesLeft           *int64             `json:"minutesLeft,omitempty"`
	Name                  *string            `json:"name,omitempty"`
	ProgressPercentage    *int64             `json:"progressPercentage,omitempty"`
	StartTime             *string            `json:"startTime,omitempty"`
	Status                *string            `json:"status,omitempty"`
}
