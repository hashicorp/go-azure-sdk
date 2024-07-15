package daprsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprSubscriptionRoutes struct {
	Default *string                      `json:"default,omitempty"`
	Rules   *[]DaprSubscriptionRouteRule `json:"rules,omitempty"`
}
