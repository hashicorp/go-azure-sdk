package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteCollection struct {
	DataLocationCode *string `json:"dataLocationCode,omitempty"`
	Hostname         *string `json:"hostname,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Root             *Root   `json:"root,omitempty"`
}
