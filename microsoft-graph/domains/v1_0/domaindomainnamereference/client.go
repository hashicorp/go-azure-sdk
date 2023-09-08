package domaindomainnamereference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDomainNameReferenceClient struct {
	Client *msgraph.Client
}

func NewDomainDomainNameReferenceClientWithBaseURI(api sdkEnv.Api) (*DomainDomainNameReferenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "domaindomainnamereference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainDomainNameReferenceClient: %+v", err)
	}

	return &DomainDomainNameReferenceClient{
		Client: client,
	}, nil
}
