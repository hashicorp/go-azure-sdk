package applicationfederatedidentitycredential

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationFederatedIdentityCredentialClient struct {
	Client *msgraph.Client
}

func NewApplicationFederatedIdentityCredentialClientWithBaseURI(api sdkEnv.Api) (*ApplicationFederatedIdentityCredentialClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationfederatedidentitycredential", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationFederatedIdentityCredentialClient: %+v", err)
	}

	return &ApplicationFederatedIdentityCredentialClient{
		Client: client,
	}, nil
}
