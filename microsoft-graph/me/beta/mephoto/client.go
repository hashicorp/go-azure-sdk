package mephoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePhotoClient struct {
	Client *msgraph.Client
}

func NewMePhotoClientWithBaseURI(api sdkEnv.Api) (*MePhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mephoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePhotoClient: %+v", err)
	}

	return &MePhotoClient{
		Client: client,
	}, nil
}
