package incrementalrestorepoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GrantAccessData struct {
	Access                   AccessLevel `json:"access"`
	DurationInSeconds        int64       `json:"durationInSeconds"`
	GetSecureVMGuestStateSAS *bool       `json:"getSecureVMGuestStateSAS,omitempty"`
}
