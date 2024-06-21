package datatypes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerSaS struct {
	ExpiryTimeStamp string `json:"expiryTimeStamp"`
	IPAddress       string `json:"ipAddress"`
	StartTimeStamp  string `json:"startTimeStamp"`
}
