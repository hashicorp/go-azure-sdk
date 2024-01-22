package departments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepartmentListResult struct {
	NextLink *string       `json:"nextLink,omitempty"`
	Value    *[]Department `json:"value,omitempty"`
}
