package department

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepartmentClient struct {
	Client *resourcemanager.Client
}

func NewDepartmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DepartmentClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "department", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepartmentClient: %+v", err)
	}

	return &DepartmentClient{
		Client: client,
	}, nil
}
