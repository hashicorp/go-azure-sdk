package directoryonpremisessynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryOnPremisesSynchronizationClient struct {
	Client *msgraph.Client
}

func NewDirectoryOnPremisesSynchronizationClientWithBaseURI(api sdkEnv.Api) (*DirectoryOnPremisesSynchronizationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryonpremisessynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryOnPremisesSynchronizationClient: %+v", err)
	}

	return &DirectoryOnPremisesSynchronizationClient{
		Client: client,
	}, nil
}
