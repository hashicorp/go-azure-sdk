package extensionproperty

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionPropertyClient struct {
	Client *msgraph.Client
}

func NewExtensionPropertyClientWithBaseURI(sdkApi sdkEnv.Api) (*ExtensionPropertyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "extensionproperty", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExtensionPropertyClient: %+v", err)
	}

	return &ExtensionPropertyClient{
		Client: client,
	}, nil
}
