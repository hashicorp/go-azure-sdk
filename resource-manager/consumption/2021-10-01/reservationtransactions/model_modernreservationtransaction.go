package reservationtransactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernReservationTransaction struct {
	Id         *string                                `json:"id,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Properties ModernReservationTransactionProperties `json:"properties"`
	Tags       *[]string                              `json:"tags,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}
