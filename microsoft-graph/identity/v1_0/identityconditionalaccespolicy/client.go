package identityconditionalaccespolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConditionalAccesPolicyClient struct {
	Client *msgraph.Client
}

func NewIdentityConditionalAccesPolicyClientWithBaseURI(api sdkEnv.Api) (*IdentityConditionalAccesPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityconditionalaccespolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityConditionalAccesPolicyClient: %+v", err)
	}

	return &IdentityConditionalAccesPolicyClient{
		Client: client,
	}, nil
}
