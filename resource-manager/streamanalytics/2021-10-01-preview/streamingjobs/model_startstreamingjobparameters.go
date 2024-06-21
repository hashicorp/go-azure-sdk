package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartStreamingJobParameters struct {
	OutputStartMode *OutputStartMode `json:"outputStartMode,omitempty"`
	OutputStartTime *string          `json:"outputStartTime,omitempty"`
}
