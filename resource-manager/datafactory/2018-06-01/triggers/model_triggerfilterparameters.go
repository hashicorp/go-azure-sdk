package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerFilterParameters struct {
	ContinuationToken *string `json:"continuationToken,omitempty"`
	ParentTriggerName *string `json:"parentTriggerName,omitempty"`
}
