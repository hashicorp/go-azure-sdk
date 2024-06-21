package localrulestacks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Changelog struct {
	Changes       []string `json:"changes"`
	LastCommitted *string  `json:"lastCommitted,omitempty"`
	LastModified  *string  `json:"lastModified,omitempty"`
}
