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

type DismissRecommendationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Recommendation
}

type DismissRecommendationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDismissRecommendationOperationOptions() DismissRecommendationOperationOptions {
	return DismissRecommendationOperationOptions{}
}

func (o DismissRecommendationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissRecommendationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissRecommendationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissRecommendation - Invoke action dismiss. Dismiss a recommendation object that you consider to be inapplicable
// to your tenant and update its status to dismissed .
func (c RecommendationClient) DismissRecommendation(ctx context.Context, id beta.DirectoryRecommendationId, input DismissRecommendationRequest, options DismissRecommendationOperationOptions) (result DismissRecommendationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dismiss", id.ID()),
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
