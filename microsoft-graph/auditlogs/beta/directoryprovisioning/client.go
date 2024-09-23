package directoryprovisioning

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryProvisioningClient struct {
	Client *msgraph.Client
}

func NewDirectoryProvisioningClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryProvisioningClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryprovisioning", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryProvisioningClient: %+v", err)
	}

	return &DirectoryProvisioningClient{
		Client: client,
	}, nil
}
