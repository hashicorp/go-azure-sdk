package azurebackuprecoverypointresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureBackupRecoveryPointResource
}

type RecoveryPointsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureBackupRecoveryPointResource
}

type RecoveryPointsListOperationOptions struct {
	Filter *string
}

func DefaultRecoveryPointsListOperationOptions() RecoveryPointsListOperationOptions {
	return RecoveryPointsListOperationOptions{}
}

func (o RecoveryPointsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecoveryPointsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecoveryPointsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecoveryPointsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecoveryPointsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecoveryPointsList ...
func (c AzureBackupRecoveryPointResourcesClient) RecoveryPointsList(ctx context.Context, id BackupInstanceId, options RecoveryPointsListOperationOptions) (result RecoveryPointsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecoveryPointsListCustomPager{},
		Path:          fmt.Sprintf("%s/recoveryPoints", id.ID()),
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

// RecoveryPointsListComplete retrieves all the results into a single object
func (c AzureBackupRecoveryPointResourcesClient) RecoveryPointsListComplete(ctx context.Context, id BackupInstanceId, options RecoveryPointsListOperationOptions) (RecoveryPointsListCompleteResult, error) {
	return c.RecoveryPointsListCompleteMatchingPredicate(ctx, id, options, AzureBackupRecoveryPointResourceOperationPredicate{})
}

// RecoveryPointsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AzureBackupRecoveryPointResourcesClient) RecoveryPointsListCompleteMatchingPredicate(ctx context.Context, id BackupInstanceId, options RecoveryPointsListOperationOptions, predicate AzureBackupRecoveryPointResourceOperationPredicate) (result RecoveryPointsListCompleteResult, err error) {
	items := make([]AzureBackupRecoveryPointResource, 0)

	resp, err := c.RecoveryPointsList(ctx, id, options)
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

	result = RecoveryPointsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
