package daprcomponentresiliencypolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentResiliencyPolicyTimeoutPolicyConfiguration struct {
	ResponseTimeoutInSeconds *int64 `json:"responseTimeoutInSeconds,omitempty"`
}
