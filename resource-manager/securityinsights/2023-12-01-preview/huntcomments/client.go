package huntcomments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntCommentsClient struct {
	Client *resourcemanager.Client
}

func NewHuntCommentsClientWithBaseURI(sdkApi sdkEnv.Api) (*HuntCommentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "huntcomments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HuntCommentsClient: %+v", err)
	}

	return &HuntCommentsClient{
		Client: client,
	}, nil
}
