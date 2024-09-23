package managedappregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistrationClient struct {
	Client *msgraph.Client
}

func NewManagedAppRegistrationClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedAppRegistrationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "managedappregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedAppRegistrationClient: %+v", err)
	}

	return &ManagedAppRegistrationClient{
		Client: client,
	}, nil
}
