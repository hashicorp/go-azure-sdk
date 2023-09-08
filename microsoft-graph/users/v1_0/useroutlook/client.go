package useroutlook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookClient struct {
	Client *msgraph.Client
}

func NewUserOutlookClientWithBaseURI(api sdkEnv.Api) (*UserOutlookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookClient: %+v", err)
	}

	return &UserOutlookClient{
		Client: client,
	}, nil
}
