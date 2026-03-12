package platformworkloadidentityrolesets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformWorkloadIdentityRoleSetsClient struct {
	Client *resourcemanager.Client
}

func NewPlatformWorkloadIdentityRoleSetsClientWithBaseURI(sdkApi sdkEnv.Api) (*PlatformWorkloadIdentityRoleSetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "platformworkloadidentityrolesets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlatformWorkloadIdentityRoleSetsClient: %+v", err)
	}

	return &PlatformWorkloadIdentityRoleSetsClient{
		Client: client,
	}, nil
}
