package notebooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookSessionProperties struct {
	DriverCores    int64  `json:"driverCores"`
	DriverMemory   string `json:"driverMemory"`
	ExecutorCores  int64  `json:"executorCores"`
	ExecutorMemory string `json:"executorMemory"`
	NumExecutors   int64  `json:"numExecutors"`
}
