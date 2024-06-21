package sourcecontrolsyncjobstreams

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlSyncJobStreamProperties struct {
	SourceControlSyncJobStreamId *string     `json:"sourceControlSyncJobStreamId,omitempty"`
	StreamType                   *StreamType `json:"streamType,omitempty"`
	Summary                      *string     `json:"summary,omitempty"`
	Time                         *string     `json:"time,omitempty"`
}
