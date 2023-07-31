package policydescription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDescriptionClient struct {
	Client *resourcemanager.Client
}

func NewPolicyDescriptionClientWithBaseURI(api environments.Api) (*PolicyDescriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "policydescription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyDescriptionClient: %+v", err)
	}

	return &PolicyDescriptionClient{
		Client: client,
	}, nil
}
