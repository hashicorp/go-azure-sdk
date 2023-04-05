package webhooks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type WebhookAction string

const (
	WebhookActionChartDelete WebhookAction = "chart_delete"
	WebhookActionChartPush   WebhookAction = "chart_push"
	WebhookActionDelete      WebhookAction = "delete"
	WebhookActionPush        WebhookAction = "push"
	WebhookActionQuarantine  WebhookAction = "quarantine"
)

func PossibleValuesForWebhookAction() []string {
	return []string{
		string(WebhookActionChartDelete),
		string(WebhookActionChartPush),
		string(WebhookActionDelete),
		string(WebhookActionPush),
		string(WebhookActionQuarantine),
	}
}

func (s *WebhookAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForWebhookAction() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = WebhookAction(decoded)
	return nil
}

type WebhookStatus string

const (
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusEnabled  WebhookStatus = "enabled"
)

func PossibleValuesForWebhookStatus() []string {
	return []string{
		string(WebhookStatusDisabled),
		string(WebhookStatusEnabled),
	}
}

func (s *WebhookStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForWebhookStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = WebhookStatus(decoded)
	return nil
}
