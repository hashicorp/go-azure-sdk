package mememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMemberOfClient struct {
	Client *msgraph.Client
}

func NewMeMemberOfClientWithBaseURI(api sdkEnv.Api) (*MeMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMemberOfClient: %+v", err)
	}

	return &MeMemberOfClient{
		Client: client,
	}, nil
}
