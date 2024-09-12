package appconsentrequestuserconsentrequestapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestUserConsentRequestApprovalClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestUserConsentRequestApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestUserConsentRequestApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestuserconsentrequestapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestUserConsentRequestApprovalClient: %+v", err)
	}

	return &AppConsentRequestUserConsentRequestApprovalClient{
		Client: client,
	}, nil
}
