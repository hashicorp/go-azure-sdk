package subscriptionquotaitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaItemProperties struct {
	Current *int64 `json:"current,omitempty"`
	Default *int64 `json:"default,omitempty"`
	Usage   *int64 `json:"usage,omitempty"`
}
