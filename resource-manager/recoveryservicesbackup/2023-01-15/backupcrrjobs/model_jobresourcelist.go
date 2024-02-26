package backupcrrjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobResourceList struct {
	NextLink *string        `json:"nextLink,omitempty"`
	Value    *[]JobResource `json:"value,omitempty"`
}
