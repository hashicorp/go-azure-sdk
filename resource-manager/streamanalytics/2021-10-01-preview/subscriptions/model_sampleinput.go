package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SampleInput struct {
	CompatibilityLevel *string `json:"compatibilityLevel,omitempty"`
	DataLocale         *string `json:"dataLocale,omitempty"`
	EventsUri          *string `json:"eventsUri,omitempty"`
	Input              *Input  `json:"input,omitempty"`
}
