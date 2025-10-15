package capturedlogs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapturedLogsClient struct {
	Client *resourcemanager.Client
}

func NewCapturedLogsClientWithBaseURI(sdkApi sdkEnv.Api) (*CapturedLogsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "capturedlogs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CapturedLogsClient: %+v", err)
	}

	return &CapturedLogsClient{
		Client: client,
	}, nil
}
