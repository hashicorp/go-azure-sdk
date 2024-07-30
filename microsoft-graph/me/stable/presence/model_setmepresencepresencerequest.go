package presence

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetMePresencePresenceRequest struct {
	Activity           *string `json:"activity,omitempty"`
	Availability       *string `json:"availability,omitempty"`
	ExpirationDuration *string `json:"expirationDuration,omitempty"`
	SessionId          *string `json:"sessionId,omitempty"`
}
