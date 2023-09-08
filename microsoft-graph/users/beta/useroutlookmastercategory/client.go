package useroutlookmastercategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookMasterCategoryClient struct {
	Client *msgraph.Client
}

func NewUserOutlookMasterCategoryClientWithBaseURI(api sdkEnv.Api) (*UserOutlookMasterCategoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlookmastercategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookMasterCategoryClient: %+v", err)
	}

	return &UserOutlookMasterCategoryClient{
		Client: client,
	}, nil
}
