package mescopedrolememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeScopedRoleMemberOfClient struct {
	Client *msgraph.Client
}

func NewMeScopedRoleMemberOfClientWithBaseURI(api sdkEnv.Api) (*MeScopedRoleMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mescopedrolememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeScopedRoleMemberOfClient: %+v", err)
	}

	return &MeScopedRoleMemberOfClient{
		Client: client,
	}, nil
}
