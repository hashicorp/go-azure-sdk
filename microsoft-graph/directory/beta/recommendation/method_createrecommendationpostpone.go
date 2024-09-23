package recommendation

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

type CreateRecommendationPostponeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Recommendation
}

type CreateRecommendationPostponeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateRecommendationPostponeOperationOptions() CreateRecommendationPostponeOperationOptions {
	return CreateRecommendationPostponeOperationOptions{}
}

func (o CreateRecommendationPostponeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRecommendationPostponeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRecommendationPostponeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRecommendationPostpone - Invoke action postpone. Postpone action on a recommendation object to a specified
// future date and time by marking its status as postponed. On the date and time provided, Microsoft Entra ID will
// automatically update the status of the recommendation object to active again.
func (c RecommendationClient) CreateRecommendationPostpone(ctx context.Context, id beta.DirectoryRecommendationId, input CreateRecommendationPostponeRequest, options CreateRecommendationPostponeOperationOptions) (result CreateRecommendationPostponeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
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

	var model beta.Recommendation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
