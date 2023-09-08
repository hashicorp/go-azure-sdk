package domainfederationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainFederationConfigurationClient struct {
	Client *msgraph.Client
}

func NewDomainFederationConfigurationClientWithBaseURI(api sdkEnv.Api) (*DomainFederationConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "domainfederationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainFederationConfigurationClient: %+v", err)
	}

	return &DomainFederationConfigurationClient{
		Client: client,
	}, nil
}
