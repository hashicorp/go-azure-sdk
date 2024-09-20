package diagnosticspackages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsPackagesClient struct {
	Client *resourcemanager.Client
}

func NewDiagnosticsPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*DiagnosticsPackagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "diagnosticspackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiagnosticsPackagesClient: %+v", err)
	}

	return &DiagnosticsPackagesClient{
		Client: client,
	}, nil
}
