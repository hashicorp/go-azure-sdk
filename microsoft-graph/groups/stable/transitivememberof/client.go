package transitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewTransitiveMemberOfClientWithBaseURI(sdkApi sdkEnv.Api) (*TransitiveMemberOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "transitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransitiveMemberOfClient: %+v", err)
	}

	return &TransitiveMemberOfClient{
		Client: client,
	}, nil
}
