package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueueSettings struct {
	JobTier *JobTier `json:"jobTier,omitempty"`
}
