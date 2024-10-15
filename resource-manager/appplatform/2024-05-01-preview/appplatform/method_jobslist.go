package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JobResource
}

type JobsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []JobResource
}

type JobsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *JobsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// JobsList ...
func (c AppPlatformClient) JobsList(ctx context.Context, id commonids.SpringCloudServiceId) (result JobsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &JobsListCustomPager{},
		Path:       fmt.Sprintf("%s/jobs", id.ID()),
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
		Values *[]JobResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// JobsListComplete retrieves all the results into a single object
func (c AppPlatformClient) JobsListComplete(ctx context.Context, id commonids.SpringCloudServiceId) (JobsListCompleteResult, error) {
	return c.JobsListCompleteMatchingPredicate(ctx, id, JobResourceOperationPredicate{})
}

// JobsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppPlatformClient) JobsListCompleteMatchingPredicate(ctx context.Context, id commonids.SpringCloudServiceId, predicate JobResourceOperationPredicate) (result JobsListCompleteResult, err error) {
	items := make([]JobResource, 0)

	resp, err := c.JobsList(ctx, id)
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

	result = JobsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
