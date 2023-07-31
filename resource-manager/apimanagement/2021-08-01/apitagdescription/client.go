package apitagdescription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiTagDescriptionClient struct {
	Client *resourcemanager.Client
}

func NewApiTagDescriptionClientWithBaseURI(api environments.Api) (*ApiTagDescriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apitagdescription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiTagDescriptionClient: %+v", err)
	}

	return &ApiTagDescriptionClient{
		Client: client,
	}, nil
}
