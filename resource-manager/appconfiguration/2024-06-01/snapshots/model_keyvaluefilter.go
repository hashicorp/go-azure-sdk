package snapshots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValueFilter struct {
	Key   string  `json:"key"`
	Label *string `json:"label,omitempty"`
}
