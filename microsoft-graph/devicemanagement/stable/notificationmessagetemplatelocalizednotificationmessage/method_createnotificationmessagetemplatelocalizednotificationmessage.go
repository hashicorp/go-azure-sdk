package notificationmessagetemplatelocalizednotificationmessage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateNotificationMessageTemplateLocalizedNotificationMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.LocalizedNotificationMessage
}

type CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions() CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions {
	return CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions{}
}

func (o CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateNotificationMessageTemplateLocalizedNotificationMessage - Create localizedNotificationMessage. Create a new
// localizedNotificationMessage object.
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) CreateNotificationMessageTemplateLocalizedNotificationMessage(ctx context.Context, id stable.DeviceManagementNotificationMessageTemplateId, input stable.LocalizedNotificationMessage, options CreateNotificationMessageTemplateLocalizedNotificationMessageOperationOptions) (result CreateNotificationMessageTemplateLocalizedNotificationMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/localizedNotificationMessages", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.LocalizedNotificationMessage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
