package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationMessage struct {
	Message *string          `json:"message,omitempty"`
	State   *ValidationState `json:"state,omitempty"`
}
