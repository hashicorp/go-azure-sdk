package standardassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewStandardAssignmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*StandardAssignmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "standardassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StandardAssignmentsClient: %+v", err)
	}

	return &StandardAssignmentsClient{
		Client: client,
	}, nil
}
