package resetqueryperformanceinsightdata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryPerformanceInsightResetDataResult struct {
	Message *string                                      `json:"message,omitempty"`
	Status  *QueryPerformanceInsightResetDataResultState `json:"status,omitempty"`
}
