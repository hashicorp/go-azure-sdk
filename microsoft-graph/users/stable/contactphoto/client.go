package contactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactPhotoClient struct {
	Client *msgraph.Client
}

func NewContactPhotoClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactPhotoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactPhotoClient: %+v", err)
	}

	return &ContactPhotoClient{
		Client: client,
	}, nil
}
