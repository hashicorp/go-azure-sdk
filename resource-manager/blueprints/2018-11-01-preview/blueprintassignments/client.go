package blueprintassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlueprintAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewBlueprintAssignmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*BlueprintAssignmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "blueprintassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BlueprintAssignmentsClient: %+v", err)
	}

	return &BlueprintAssignmentsClient{
		Client: client,
	}, nil
}
