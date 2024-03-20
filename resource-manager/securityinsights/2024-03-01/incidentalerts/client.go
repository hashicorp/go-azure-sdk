package incidentalerts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentAlertsClient struct {
	Client *resourcemanager.Client
}

func NewIncidentAlertsClientWithBaseURI(sdkApi sdkEnv.Api) (*IncidentAlertsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "incidentalerts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IncidentAlertsClient: %+v", err)
	}

	return &IncidentAlertsClient{
		Client: client,
	}, nil
}
