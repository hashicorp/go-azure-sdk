package scopedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedMemberClient struct {
	Client *msgraph.Client
}

func NewScopedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*ScopedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "scopedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScopedMemberClient: %+v", err)
	}

	return &ScopedMemberClient{
		Client: client,
	}, nil
}
