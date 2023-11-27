package servertrustgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerTrustGroup struct {
	Id         *string                     `json:"id,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *ServerTrustGroupProperties `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
