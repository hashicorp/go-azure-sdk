package accesspolicyassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewAccessPolicyAssignmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessPolicyAssignmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "accesspolicyassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessPolicyAssignmentsClient: %+v", err)
	}

	return &AccessPolicyAssignmentsClient{
		Client: client,
	}, nil
}
