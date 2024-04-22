package certificateordersdiagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SampleUtterance struct {
	Links *[]string `json:"links,omitempty"`
	Qid   *string   `json:"qid,omitempty"`
	Text  *string   `json:"text,omitempty"`
}
