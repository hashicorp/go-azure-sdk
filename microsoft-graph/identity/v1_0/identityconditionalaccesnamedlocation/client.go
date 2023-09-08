package identityconditionalaccesnamedlocation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConditionalAccesNamedLocationClient struct {
	Client *msgraph.Client
}

func NewIdentityConditionalAccesNamedLocationClientWithBaseURI(api sdkEnv.Api) (*IdentityConditionalAccesNamedLocationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityconditionalaccesnamedlocation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityConditionalAccesNamedLocationClient: %+v", err)
	}

	return &IdentityConditionalAccesNamedLocationClient{
		Client: client,
	}, nil
}
