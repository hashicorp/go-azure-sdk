package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiInstanceSettings struct {
	CommonResourceFiles     *[]ResourceFile `json:"commonResourceFiles,omitempty"`
	CoordinationCommandLine string          `json:"coordinationCommandLine"`
	NumberOfInstances       *int64          `json:"numberOfInstances,omitempty"`
}
