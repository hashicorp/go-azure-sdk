package incidenttasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTasksClient struct {
	Client *resourcemanager.Client
}

func NewIncidentTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*IncidentTasksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "incidenttasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncidentTasksClient: %+v", err)
	}

	return &IncidentTasksClient{
		Client: client,
	}, nil
}
