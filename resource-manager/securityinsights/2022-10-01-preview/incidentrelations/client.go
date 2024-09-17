package incidentrelations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentRelationsClient struct {
	Client *resourcemanager.Client
}

func NewIncidentRelationsClientWithBaseURI(sdkApi sdkEnv.Api) (*IncidentRelationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "incidentrelations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncidentRelationsClient: %+v", err)
	}

	return &IncidentRelationsClient{
		Client: client,
	}, nil
}
