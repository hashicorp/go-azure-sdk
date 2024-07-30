package chatpermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewChatPermissionGrantClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatPermissionGrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatpermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatPermissionGrantClient: %+v", err)
	}

	return &ChatPermissionGrantClient{
		Client: client,
	}, nil
}
