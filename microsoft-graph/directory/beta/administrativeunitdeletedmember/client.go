package administrativeunitdeletedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrativeUnitDeletedMemberClient struct {
	Client *msgraph.Client
}

func NewAdministrativeUnitDeletedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*AdministrativeUnitDeletedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "administrativeunitdeletedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrativeUnitDeletedMemberClient: %+v", err)
	}

	return &AdministrativeUnitDeletedMemberClient{
		Client: client,
	}, nil
}
