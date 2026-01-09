package scopedrolememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedRoleMemberOfClient struct {
	Client *msgraph.Client
}

func NewScopedRoleMemberOfClientWithBaseURI(sdkApi sdkEnv.Api) (*ScopedRoleMemberOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "scopedrolememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScopedRoleMemberOfClient: %+v", err)
	}

	return &ScopedRoleMemberOfClient{
		Client: client,
	}, nil
}
