package mecalendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
