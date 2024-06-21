package integrationaccountsessions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountSessionProperties struct {
	ChangedTime *string      `json:"changedTime,omitempty"`
	Content     *interface{} `json:"content,omitempty"`
	CreatedTime *string      `json:"createdTime,omitempty"`
}
