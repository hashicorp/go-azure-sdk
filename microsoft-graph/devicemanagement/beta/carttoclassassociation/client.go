package carttoclassassociation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CartToClassAssociationClient struct {
	Client *msgraph.Client
}

func NewCartToClassAssociationClientWithBaseURI(sdkApi sdkEnv.Api) (*CartToClassAssociationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "carttoclassassociation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CartToClassAssociationClient: %+v", err)
	}

	return &CartToClassAssociationClient{
		Client: client,
	}, nil
}
