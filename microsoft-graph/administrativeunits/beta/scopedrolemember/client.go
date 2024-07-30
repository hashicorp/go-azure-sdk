package scopedrolemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedRoleMemberClient struct {
	Client *msgraph.Client
}

func NewScopedRoleMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*ScopedRoleMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "scopedrolemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScopedRoleMemberClient: %+v", err)
	}

	return &ScopedRoleMemberClient{
		Client: client,
	}, nil
}
