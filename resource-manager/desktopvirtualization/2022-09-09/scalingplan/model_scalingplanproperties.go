package scalingplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingPlanProperties struct {
	Description        *string                     `json:"description,omitempty"`
	ExclusionTag       *string                     `json:"exclusionTag,omitempty"`
	FriendlyName       *string                     `json:"friendlyName,omitempty"`
	HostPoolReferences *[]ScalingHostPoolReference `json:"hostPoolReferences,omitempty"`
	HostPoolType       *ScalingHostPoolType        `json:"hostPoolType,omitempty"`
	ObjectId           *string                     `json:"objectId,omitempty"`
	Schedules          *[]ScalingSchedule          `json:"schedules,omitempty"`
	TimeZone           string                      `json:"timeZone"`
}
