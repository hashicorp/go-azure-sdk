package useronenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteResourceClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteResourceClient: %+v", err)
	}

	return &UserOnenoteResourceClient{
		Client: client,
	}, nil
}
