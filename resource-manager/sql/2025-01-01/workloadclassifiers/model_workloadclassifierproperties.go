package workloadclassifiers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadClassifierProperties struct {
	Context    *string `json:"context,omitempty"`
	EndTime    *string `json:"endTime,omitempty"`
	Importance *string `json:"importance,omitempty"`
	Label      *string `json:"label,omitempty"`
	MemberName string  `json:"memberName"`
	StartTime  *string `json:"startTime,omitempty"`
}
