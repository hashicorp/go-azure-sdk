package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentPathMap struct {
	ContentType *ContentType `json:"contentType,omitempty"`
	Path        *string      `json:"path,omitempty"`
}
