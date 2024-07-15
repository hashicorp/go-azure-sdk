package daprsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprSubscriptionRouteRule struct {
	Match *string `json:"match,omitempty"`
	Path  *string `json:"path,omitempty"`
}
