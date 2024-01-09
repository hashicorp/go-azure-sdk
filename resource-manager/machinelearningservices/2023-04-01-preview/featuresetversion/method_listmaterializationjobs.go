package featuresetversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMaterializationJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FeaturesetJob
}

type ListMaterializationJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FeaturesetJob
}

type ListMaterializationJobsOperationOptions struct {
	FeatureWindowEnd   *string
	FeatureWindowStart *string
	Filters            *string
	Skip               *string
}

func DefaultListMaterializationJobsOperationOptions() ListMaterializationJobsOperationOptions {
	return ListMaterializationJobsOperationOptions{}
}

func (o ListMaterializationJobsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMaterializationJobsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListMaterializationJobsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.FeatureWindowEnd != nil {
		out.Append("featureWindowEnd", fmt.Sprintf("%v", *o.FeatureWindowEnd))
	}
	if o.FeatureWindowStart != nil {
		out.Append("featureWindowStart", fmt.Sprintf("%v", *o.FeatureWindowStart))
	}
	if o.Filters != nil {
		out.Append("filters", fmt.Sprintf("%v", *o.Filters))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	return &out
}

// ListMaterializationJobs ...
func (c FeaturesetVersionClient) ListMaterializationJobs(ctx context.Context, id FeatureSetVersionId, options ListMaterializationJobsOperationOptions) (result ListMaterializationJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/listMaterializationJobs", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]FeaturesetJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMaterializationJobsComplete retrieves all the results into a single object
func (c FeaturesetVersionClient) ListMaterializationJobsComplete(ctx context.Context, id FeatureSetVersionId, options ListMaterializationJobsOperationOptions) (ListMaterializationJobsCompleteResult, error) {
	return c.ListMaterializationJobsCompleteMatchingPredicate(ctx, id, options, FeaturesetJobOperationPredicate{})
}

// ListMaterializationJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FeaturesetVersionClient) ListMaterializationJobsCompleteMatchingPredicate(ctx context.Context, id FeatureSetVersionId, options ListMaterializationJobsOperationOptions, predicate FeaturesetJobOperationPredicate) (result ListMaterializationJobsCompleteResult, err error) {
	items := make([]FeaturesetJob, 0)

	resp, err := c.ListMaterializationJobs(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListMaterializationJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
