package impactedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateImpactedResourcePostponeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ImpactedResource
}

type CreateImpactedResourcePostponeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateImpactedResourcePostponeOperationOptions() CreateImpactedResourcePostponeOperationOptions {
	return CreateImpactedResourcePostponeOperationOptions{}
}

func (o CreateImpactedResourcePostponeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateImpactedResourcePostponeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateImpactedResourcePostponeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateImpactedResourcePostpone - Invoke action postpone. Postpone action on an impactedResource object to a specified
// future date and time by marking its status as postponed. On the specified date and time, Microsoft Entra ID will
// automatically mark the status of the impactedResource object to active.
func (c ImpactedResourceClient) CreateImpactedResourcePostpone(ctx context.Context, id beta.DirectoryImpactedResourceId, input CreateImpactedResourcePostponeRequest, options CreateImpactedResourcePostponeOperationOptions) (result CreateImpactedResourcePostponeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/postpone", id.ID()),
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

	var model beta.ImpactedResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
