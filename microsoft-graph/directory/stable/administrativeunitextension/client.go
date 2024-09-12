package administrativeunitextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrativeUnitExtensionClient struct {
	Client *msgraph.Client
}

func NewAdministrativeUnitExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*AdministrativeUnitExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "administrativeunitextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrativeUnitExtensionClient: %+v", err)
	}

	return &AdministrativeUnitExtensionClient{
		Client: client,
	}, nil
}
