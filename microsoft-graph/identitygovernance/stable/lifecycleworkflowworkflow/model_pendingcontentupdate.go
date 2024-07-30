package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingContentUpdate struct {
	ODataType      *string `json:"@odata.type,omitempty"`
	QueuedDateTime *string `json:"queuedDateTime,omitempty"`
}
