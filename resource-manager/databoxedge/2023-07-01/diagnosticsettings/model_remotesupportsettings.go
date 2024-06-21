package diagnosticsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteSupportSettings struct {
	AccessLevel              *AccessLevel           `json:"accessLevel,omitempty"`
	ExpirationTimeStampInUTC *string                `json:"expirationTimeStampInUTC,omitempty"`
	RemoteApplicationType    *RemoteApplicationType `json:"remoteApplicationType,omitempty"`
}
