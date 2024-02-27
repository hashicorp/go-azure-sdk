package sourcecontrols

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Webhook struct {
	RotateWebhookSecret     *bool   `json:"rotateWebhookSecret,omitempty"`
	WebhookId               *string `json:"webhookId,omitempty"`
	WebhookSecretUpdateTime *string `json:"webhookSecretUpdateTime,omitempty"`
	WebhookUrl              *string `json:"webhookUrl,omitempty"`
}

func (o *Webhook) GetWebhookSecretUpdateTimeAsTime() (*time.Time, error) {
	if o.WebhookSecretUpdateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.WebhookSecretUpdateTime, "2006-01-02T15:04:05Z07:00")
}
