package appconsentappconsentrequestuserconsentrequestapprovalstage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentAppConsentRequestUserConsentRequestApprovalStageClient struct {
	Client *msgraph.Client
}

func NewAppConsentAppConsentRequestUserConsentRequestApprovalStageClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentAppConsentRequestUserConsentRequestApprovalStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentappconsentrequestuserconsentrequestapprovalstage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentAppConsentRequestUserConsentRequestApprovalStageClient: %+v", err)
	}

	return &AppConsentAppConsentRequestUserConsentRequestApprovalStageClient{
		Client: client,
	}, nil
}
