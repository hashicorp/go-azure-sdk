package domainsharedemaildomaininvitation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainSharedEmailDomainInvitationClient struct {
	Client *msgraph.Client
}

func NewDomainSharedEmailDomainInvitationClientWithBaseURI(api sdkEnv.Api) (*DomainSharedEmailDomainInvitationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "domainsharedemaildomaininvitation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainSharedEmailDomainInvitationClient: %+v", err)
	}

	return &DomainSharedEmailDomainInvitationClient{
		Client: client,
	}, nil
}
