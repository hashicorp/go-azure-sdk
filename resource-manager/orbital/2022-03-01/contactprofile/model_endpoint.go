package contactprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndPoint struct {
	EndPointName string   `json:"endPointName"`
	IpAddress    string   `json:"ipAddress"`
	Port         string   `json:"port"`
	Protocol     Protocol `json:"protocol"`
}
