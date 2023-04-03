package actiongroupsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebhookReceiver struct {
	Name       string `json:"name"`
	ServiceUri string `json:"serviceUri"`
}
