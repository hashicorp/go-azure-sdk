package metransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewMeTransitiveMemberOfClientWithBaseURI(api sdkEnv.Api) (*MeTransitiveMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTransitiveMemberOfClient: %+v", err)
	}

	return &MeTransitiveMemberOfClient{
		Client: client,
	}, nil
}
