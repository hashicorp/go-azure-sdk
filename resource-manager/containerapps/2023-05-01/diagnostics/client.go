package diagnostics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsClient struct {
	Client *resourcemanager.Client
}

func NewDiagnosticsClientWithBaseURI(sdkApi sdkEnv.Api) (*DiagnosticsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "diagnostics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiagnosticsClient: %+v", err)
	}

	return &DiagnosticsClient{
		Client: client,
	}, nil
}
