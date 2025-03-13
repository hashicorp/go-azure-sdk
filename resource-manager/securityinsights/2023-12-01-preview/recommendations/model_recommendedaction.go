package recommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedAction struct {
	LinkText string    `json:"linkText"`
	LinkURL  string    `json:"linkUrl"`
	State    *Priority `json:"state,omitempty"`
}
