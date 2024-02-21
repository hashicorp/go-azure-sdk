package incidentteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTeamClient struct {
	Client *resourcemanager.Client
}

func NewIncidentTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*IncidentTeamClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "incidentteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncidentTeamClient: %+v", err)
	}

	return &IncidentTeamClient{
		Client: client,
	}, nil
}
