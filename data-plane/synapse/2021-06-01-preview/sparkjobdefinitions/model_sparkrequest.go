package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkRequest struct {
	Archives       *[]string          `json:"archives,omitempty"`
	Args           *[]string          `json:"args,omitempty"`
	ClassName      *string            `json:"className,omitempty"`
	Conf           *map[string]string `json:"conf,omitempty"`
	DriverCores    *int64             `json:"driverCores,omitempty"`
	DriverMemory   *string            `json:"driverMemory,omitempty"`
	ExecutorCores  *int64             `json:"executorCores,omitempty"`
	ExecutorMemory *string            `json:"executorMemory,omitempty"`
	File           *string            `json:"file,omitempty"`
	Files          *[]string          `json:"files,omitempty"`
	Jars           *[]string          `json:"jars,omitempty"`
	Name           *string            `json:"name,omitempty"`
	NumExecutors   *int64             `json:"numExecutors,omitempty"`
	PyFiles        *[]string          `json:"pyFiles,omitempty"`
}
