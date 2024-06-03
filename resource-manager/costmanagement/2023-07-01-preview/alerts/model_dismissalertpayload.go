package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DismissAlertPayload struct {
	Properties *AlertProperties `json:"properties,omitempty"`
}
