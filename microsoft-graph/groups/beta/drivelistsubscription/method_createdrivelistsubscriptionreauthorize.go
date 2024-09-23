package drivelistsubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDriveListSubscriptionReauthorizeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDriveListSubscriptionReauthorizeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDriveListSubscriptionReauthorizeOperationOptions() CreateDriveListSubscriptionReauthorizeOperationOptions {
	return CreateDriveListSubscriptionReauthorizeOperationOptions{}
}

func (o CreateDriveListSubscriptionReauthorizeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveListSubscriptionReauthorizeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveListSubscriptionReauthorizeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveListSubscriptionReauthorize - Invoke action reauthorize. Reauthorize a subscription when you receive a
// reauthorizationRequired challenge.
func (c DriveListSubscriptionClient) CreateDriveListSubscriptionReauthorize(ctx context.Context, id beta.GroupIdDriveIdListSubscriptionId, options CreateDriveListSubscriptionReauthorizeOperationOptions) (result CreateDriveListSubscriptionReauthorizeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reauthorize", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	return
}
