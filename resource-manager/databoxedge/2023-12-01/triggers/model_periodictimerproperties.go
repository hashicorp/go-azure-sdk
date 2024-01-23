package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeriodicTimerProperties struct {
	CustomContextTag *string                 `json:"customContextTag,omitempty"`
	SinkInfo         RoleSinkInfo            `json:"sinkInfo"`
	SourceInfo       PeriodicTimerSourceInfo `json:"sourceInfo"`
}
