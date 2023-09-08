package meprofilepatent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfilePatentClient struct {
	Client *msgraph.Client
}

func NewMeProfilePatentClientWithBaseURI(api sdkEnv.Api) (*MeProfilePatentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilepatent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfilePatentClient: %+v", err)
	}

	return &MeProfilePatentClient{
		Client: client,
	}, nil
}
