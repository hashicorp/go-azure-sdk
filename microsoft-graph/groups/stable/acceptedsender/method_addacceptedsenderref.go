package acceptedsender

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

type AddAcceptedSenderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddAcceptedSenderRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddAcceptedSenderRefOperationOptions() AddAcceptedSenderRefOperationOptions {
	return AddAcceptedSenderRefOperationOptions{}
}

func (o AddAcceptedSenderRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddAcceptedSenderRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddAcceptedSenderRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddAcceptedSenderRef - Create acceptedSender. Specify the user or group in @odata.id in the request body. Users in
// the accepted senders list can post to conversations of the group. Make sure you don't specify the same user or group
// in the accepted senders and rejected senders lists, otherwise you'll get an error.
func (c AcceptedSenderClient) AddAcceptedSenderRef(ctx context.Context, id stable.GroupId, input stable.ReferenceCreate, options AddAcceptedSenderRefOperationOptions) (result AddAcceptedSenderRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/acceptedSenders/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	return
}
