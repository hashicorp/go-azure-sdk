package webhooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventContent struct {
	Action    *string  `json:"action,omitempty"`
	Actor     *Actor   `json:"actor,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Request   *Request `json:"request,omitempty"`
	Source    *Source  `json:"source,omitempty"`
	Target    *Target  `json:"target,omitempty"`
	Timestamp *string  `json:"timestamp,omitempty"`
}
