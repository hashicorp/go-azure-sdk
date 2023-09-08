package useronenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteResourceContentClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteResourceContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteResourceContentClient: %+v", err)
	}

	return &UserOnenoteResourceContentClient{
		Client: client,
	}, nil
}
