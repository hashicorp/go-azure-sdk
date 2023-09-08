package applicationextensionproperty

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationExtensionPropertyClient struct {
	Client *msgraph.Client
}

func NewApplicationExtensionPropertyClientWithBaseURI(api sdkEnv.Api) (*ApplicationExtensionPropertyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationextensionproperty", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationExtensionPropertyClient: %+v", err)
	}

	return &ApplicationExtensionPropertyClient{
		Client: client,
	}, nil
}
