package presence

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetPresenceRequest struct {
	Activity           *string               `json:"activity,omitempty"`
	Availability       *string               `json:"availability,omitempty"`
	ExpirationDuration nullable.Type[string] `json:"expirationDuration,omitempty"`
	SessionId          nullable.Type[string] `json:"sessionId,omitempty"`
}
