package mechatpermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewMeChatPermissionGrantClientWithBaseURI(api sdkEnv.Api) (*MeChatPermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatpermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatPermissionGrantClient: %+v", err)
	}

	return &MeChatPermissionGrantClient{
		Client: client,
	}, nil
}
