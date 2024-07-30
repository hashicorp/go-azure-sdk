package synchronizationjob

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

type ListSynchronizationJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SynchronizationJob
}

type ListSynchronizationJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SynchronizationJob
}

type ListSynchronizationJobsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationJobs ...
func (c SynchronizationJobClient) ListSynchronizationJobs(ctx context.Context, id ServicePrincipalId) (result ListSynchronizationJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSynchronizationJobsCustomPager{},
		Path:       fmt.Sprintf("%s/synchronization/jobs", id.ID()),
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
		Values *[]stable.SynchronizationJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationJobsComplete retrieves all the results into a single object
func (c SynchronizationJobClient) ListSynchronizationJobsComplete(ctx context.Context, id ServicePrincipalId) (ListSynchronizationJobsCompleteResult, error) {
	return c.ListSynchronizationJobsCompleteMatchingPredicate(ctx, id, SynchronizationJobOperationPredicate{})
}

// ListSynchronizationJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationJobClient) ListSynchronizationJobsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate SynchronizationJobOperationPredicate) (result ListSynchronizationJobsCompleteResult, err error) {
	items := make([]stable.SynchronizationJob, 0)

	resp, err := c.ListSynchronizationJobs(ctx, id)
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

	result = ListSynchronizationJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
