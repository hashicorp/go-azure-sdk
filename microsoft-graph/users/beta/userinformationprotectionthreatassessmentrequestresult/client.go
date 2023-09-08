package userinformationprotectionthreatassessmentrequestresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionThreatAssessmentRequestResultClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionThreatAssessmentRequestResultClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionthreatassessmentrequestresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionThreatAssessmentRequestResultClient: %+v", err)
	}

	return &UserInformationProtectionThreatAssessmentRequestResultClient{
		Client: client,
	}, nil
}
