package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobProperties struct {
	Archives       *[]string    `json:"archives,omitempty"`
	Args           *[]string    `json:"args,omitempty"`
	ClassName      *string      `json:"className,omitempty"`
	Conf           *interface{} `json:"conf,omitempty"`
	DriverCores    int64        `json:"driverCores"`
	DriverMemory   string       `json:"driverMemory"`
	ExecutorCores  int64        `json:"executorCores"`
	ExecutorMemory string       `json:"executorMemory"`
	File           string       `json:"file"`
	Files          *[]string    `json:"files,omitempty"`
	Jars           *[]string    `json:"jars,omitempty"`
	Name           *string      `json:"name,omitempty"`
	NumExecutors   int64        `json:"numExecutors"`
}
