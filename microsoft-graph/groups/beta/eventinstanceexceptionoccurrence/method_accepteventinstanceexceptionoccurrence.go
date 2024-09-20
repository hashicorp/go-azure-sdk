package eventinstanceexceptionoccurrence

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

type AcceptEventInstanceExceptionOccurrenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AcceptEventInstanceExceptionOccurrenceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAcceptEventInstanceExceptionOccurrenceOperationOptions() AcceptEventInstanceExceptionOccurrenceOperationOptions {
	return AcceptEventInstanceExceptionOccurrenceOperationOptions{}
}

func (o AcceptEventInstanceExceptionOccurrenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AcceptEventInstanceExceptionOccurrenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AcceptEventInstanceExceptionOccurrenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AcceptEventInstanceExceptionOccurrence - Invoke action accept. Accept the specified event in a user calendar.
func (c EventInstanceExceptionOccurrenceClient) AcceptEventInstanceExceptionOccurrence(ctx context.Context, id beta.GroupIdEventIdInstanceIdExceptionOccurrenceId, input AcceptEventInstanceExceptionOccurrenceRequest, options AcceptEventInstanceExceptionOccurrenceOperationOptions) (result AcceptEventInstanceExceptionOccurrenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accept", id.ID()),
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
