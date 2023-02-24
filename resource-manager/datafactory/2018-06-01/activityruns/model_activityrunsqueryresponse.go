package activityruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityRunsQueryResponse struct {
	ContinuationToken *string       `json:"continuationToken,omitempty"`
	Value             []ActivityRun `json:"value"`
}
