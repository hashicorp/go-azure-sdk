package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GCPRequestProperties struct {
	ProjectId         string   `json:"projectId"`
	SubscriptionNames []string `json:"subscriptionNames"`
}
