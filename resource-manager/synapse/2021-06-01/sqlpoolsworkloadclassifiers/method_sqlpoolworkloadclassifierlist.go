package sqlpoolsworkloadclassifiers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolWorkloadClassifierListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadClassifier
}

type SqlPoolWorkloadClassifierListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadClassifier
}

type SqlPoolWorkloadClassifierListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolWorkloadClassifierListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolWorkloadClassifierList ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierList(ctx context.Context, id WorkloadGroupId) (result SqlPoolWorkloadClassifierListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlPoolWorkloadClassifierListCustomPager{},
		Path:       fmt.Sprintf("%s/workloadClassifiers", id.ID()),
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
		Values *[]WorkloadClassifier `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolWorkloadClassifierListComplete retrieves all the results into a single object
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierListComplete(ctx context.Context, id WorkloadGroupId) (SqlPoolWorkloadClassifierListCompleteResult, error) {
	return c.SqlPoolWorkloadClassifierListCompleteMatchingPredicate(ctx, id, WorkloadClassifierOperationPredicate{})
}

// SqlPoolWorkloadClassifierListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierListCompleteMatchingPredicate(ctx context.Context, id WorkloadGroupId, predicate WorkloadClassifierOperationPredicate) (result SqlPoolWorkloadClassifierListCompleteResult, err error) {
	items := make([]WorkloadClassifier, 0)

	resp, err := c.SqlPoolWorkloadClassifierList(ctx, id)
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

	result = SqlPoolWorkloadClassifierListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
