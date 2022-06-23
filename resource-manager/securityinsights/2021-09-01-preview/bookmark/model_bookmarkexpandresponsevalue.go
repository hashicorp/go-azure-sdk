package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkExpandResponseValue struct {
	Edges    *[]ConnectedEntity `json:"edges,omitempty"`
	Entities *[]Entity          `json:"entities,omitempty"`
}
