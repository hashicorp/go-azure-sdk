package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookLinks struct {
	ODataType        *string       `json:"@odata.type,omitempty"`
	OneNoteClientUrl *ExternalLink `json:"oneNoteClientUrl,omitempty"`
	OneNoteWebUrl    *ExternalLink `json:"oneNoteWebUrl,omitempty"`
}
