package profilepatent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfilePatentClient struct {
	Client *msgraph.Client
}

func NewProfilePatentClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfilePatentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilepatent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfilePatentClient: %+v", err)
	}

	return &ProfilePatentClient{
		Client: client,
	}, nil
}
