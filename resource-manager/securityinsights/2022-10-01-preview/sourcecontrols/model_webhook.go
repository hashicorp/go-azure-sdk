package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Webhook struct {
	RotateWebhookSecret     *bool   `json:"rotateWebhookSecret,omitempty"`
	WebhookId               *string `json:"webhookId,omitempty"`
	WebhookSecretUpdateTime *string `json:"webhookSecretUpdateTime,omitempty"`
	WebhookURL              *string `json:"webhookUrl,omitempty"`
}
