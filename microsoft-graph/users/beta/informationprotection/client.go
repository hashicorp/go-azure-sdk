package informationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionClient: %+v", err)
	}

	return &InformationProtectionClient{
		Client: client,
	}, nil
}
