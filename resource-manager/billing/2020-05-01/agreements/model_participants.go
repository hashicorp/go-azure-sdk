package agreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Participants struct {
	Email      *string `json:"email,omitempty"`
	Status     *string `json:"status,omitempty"`
	StatusDate *string `json:"statusDate,omitempty"`
}
