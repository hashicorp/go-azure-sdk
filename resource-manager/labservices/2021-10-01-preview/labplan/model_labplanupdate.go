package labplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabPlanUpdate struct {
	Properties *LabPlanUpdateProperties `json:"properties,omitempty"`
	Tags       *[]string                `json:"tags,omitempty"`
}
