package revisions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValueListResult struct {
	Items    *[]KeyValue `json:"items,omitempty"`
	NextLink *string     `json:"@nextLink,omitempty"`
}
