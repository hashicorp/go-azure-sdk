package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServicePort struct {
	AppProtocol *string                                `json:"appProtocol,omitempty"`
	Name        *string                                `json:"name,omitempty"`
	NodePort    *int64                                 `json:"nodePort,omitempty"`
	ODataType   *string                                `json:"@odata.type,omitempty"`
	Port        *int64                                 `json:"port,omitempty"`
	Protocol    *SecurityKubernetesServicePortProtocol `json:"protocol,omitempty"`
	TargetPort  *string                                `json:"targetPort,omitempty"`
}
