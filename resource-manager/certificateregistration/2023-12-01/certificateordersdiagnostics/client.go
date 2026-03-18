package certificateordersdiagnostics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateOrdersDiagnosticsClient struct {
	Client *resourcemanager.Client
}

func NewCertificateOrdersDiagnosticsClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateOrdersDiagnosticsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "certificateordersdiagnostics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateOrdersDiagnosticsClient: %+v", err)
	}

	return &CertificateOrdersDiagnosticsClient{
		Client: client,
	}, nil
}
