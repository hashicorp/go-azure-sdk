package mecontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactPhotoClient struct {
	Client *msgraph.Client
}

func NewMeContactPhotoClientWithBaseURI(api sdkEnv.Api) (*MeContactPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactPhotoClient: %+v", err)
	}

	return &MeContactPhotoClient{
		Client: client,
	}, nil
}
