package customsecurityattributeaudit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeAuditClient struct {
	Client *msgraph.Client
}

func NewCustomSecurityAttributeAuditClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomSecurityAttributeAuditClient, error) {
	client, err := msgraph.NewClient(sdkApi, "customsecurityattributeaudit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomSecurityAttributeAuditClient: %+v", err)
	}

	return &CustomSecurityAttributeAuditClient{
		Client: client,
	}, nil
}
