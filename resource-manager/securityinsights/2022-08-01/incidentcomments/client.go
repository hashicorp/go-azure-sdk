package incidentcomments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentCommentsClient struct {
	Client *resourcemanager.Client
}

func NewIncidentCommentsClientWithBaseURI(sdkApi sdkEnv.Api) (*IncidentCommentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "incidentcomments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncidentCommentsClient: %+v", err)
	}

	return &IncidentCommentsClient{
		Client: client,
	}, nil
}
