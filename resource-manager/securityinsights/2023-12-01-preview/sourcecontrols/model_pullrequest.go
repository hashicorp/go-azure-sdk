package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PullRequest struct {
	State *State  `json:"state,omitempty"`
	Url   *string `json:"url,omitempty"`
}
