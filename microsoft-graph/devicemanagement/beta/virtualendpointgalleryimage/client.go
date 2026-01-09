package virtualendpointgalleryimage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointGalleryImageClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointGalleryImageClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointGalleryImageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointgalleryimage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointGalleryImageClient: %+v", err)
	}

	return &VirtualEndpointGalleryImageClient{
		Client: client,
	}, nil
}
