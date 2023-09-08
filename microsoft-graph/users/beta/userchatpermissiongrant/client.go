package userchatpermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewUserChatPermissionGrantClientWithBaseURI(api sdkEnv.Api) (*UserChatPermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatpermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatPermissionGrantClient: %+v", err)
	}

	return &UserChatPermissionGrantClient{
		Client: client,
	}, nil
}
