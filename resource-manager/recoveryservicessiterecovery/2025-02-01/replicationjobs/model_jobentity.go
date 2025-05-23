package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobEntity struct {
	JobFriendlyName    *string `json:"jobFriendlyName,omitempty"`
	JobId              *string `json:"jobId,omitempty"`
	JobScenarioName    *string `json:"jobScenarioName,omitempty"`
	TargetInstanceType *string `json:"targetInstanceType,omitempty"`
	TargetObjectId     *string `json:"targetObjectId,omitempty"`
	TargetObjectName   *string `json:"targetObjectName,omitempty"`
}
