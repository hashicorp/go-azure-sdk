package replicationprotectionclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentScenarioDetails struct {
	JobId        *string `json:"jobId,omitempty"`
	ScenarioName *string `json:"scenarioName,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
}
