package virtualendpointauditevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointAuditEventClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointAuditEventClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointAuditEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointauditevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointAuditEventClient: %+v", err)
	}

	return &VirtualEndpointAuditEventClient{
		Client: client,
	}, nil
}
