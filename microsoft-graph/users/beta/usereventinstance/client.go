package usereventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceClient: %+v", err)
	}

	return &UserEventInstanceClient{
		Client: client,
	}, nil
}
