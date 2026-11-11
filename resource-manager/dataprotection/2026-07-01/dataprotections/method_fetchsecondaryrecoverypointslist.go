package dataprotections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FetchSecondaryRecoveryPointsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureBackupRecoveryPointResource
}

type FetchSecondaryRecoveryPointsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureBackupRecoveryPointResource
}

type FetchSecondaryRecoveryPointsListOperationOptions struct {
	Filter *string
}

func DefaultFetchSecondaryRecoveryPointsListOperationOptions() FetchSecondaryRecoveryPointsListOperationOptions {
	return FetchSecondaryRecoveryPointsListOperationOptions{}
}

func (o FetchSecondaryRecoveryPointsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o FetchSecondaryRecoveryPointsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FetchSecondaryRecoveryPointsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type FetchSecondaryRecoveryPointsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *FetchSecondaryRecoveryPointsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// FetchSecondaryRecoveryPointsList ...
func (c DataprotectionsClient) FetchSecondaryRecoveryPointsList(ctx context.Context, id ProviderLocationId, input FetchSecondaryRPsRequestParameters, options FetchSecondaryRecoveryPointsListOperationOptions) (result FetchSecondaryRecoveryPointsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &FetchSecondaryRecoveryPointsListCustomPager{},
		Path:          fmt.Sprintf("%s/fetchSecondaryRecoveryPoints", id.ID()),
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
		Values *[]AzureBackupRecoveryPointResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// FetchSecondaryRecoveryPointsListComplete retrieves all the results into a single object
func (c DataprotectionsClient) FetchSecondaryRecoveryPointsListComplete(ctx context.Context, id ProviderLocationId, input FetchSecondaryRPsRequestParameters, options FetchSecondaryRecoveryPointsListOperationOptions) (FetchSecondaryRecoveryPointsListCompleteResult, error) {
	return c.FetchSecondaryRecoveryPointsListCompleteMatchingPredicate(ctx, id, input, options, AzureBackupRecoveryPointResourceOperationPredicate{})
}

// FetchSecondaryRecoveryPointsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataprotectionsClient) FetchSecondaryRecoveryPointsListCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, input FetchSecondaryRPsRequestParameters, options FetchSecondaryRecoveryPointsListOperationOptions, predicate AzureBackupRecoveryPointResourceOperationPredicate) (result FetchSecondaryRecoveryPointsListCompleteResult, err error) {
	items := make([]AzureBackupRecoveryPointResource, 0)

	resp, err := c.FetchSecondaryRecoveryPointsList(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = FetchSecondaryRecoveryPointsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
