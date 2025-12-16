package monitorresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorsListLinkableEnvironmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LinkableEnvironmentResponse
}

type MonitorsListLinkableEnvironmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []LinkableEnvironmentResponse
}

type MonitorsListLinkableEnvironmentsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MonitorsListLinkableEnvironmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MonitorsListLinkableEnvironments ...
func (c MonitorResourcesClient) MonitorsListLinkableEnvironments(ctx context.Context, id MonitorId, input LinkableEnvironmentRequest) (result MonitorsListLinkableEnvironmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &MonitorsListLinkableEnvironmentsCustomPager{},
		Path:       fmt.Sprintf("%s/listLinkableEnvironments", id.ID()),
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
		Values *[]LinkableEnvironmentResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MonitorsListLinkableEnvironmentsComplete retrieves all the results into a single object
func (c MonitorResourcesClient) MonitorsListLinkableEnvironmentsComplete(ctx context.Context, id MonitorId, input LinkableEnvironmentRequest) (MonitorsListLinkableEnvironmentsCompleteResult, error) {
	return c.MonitorsListLinkableEnvironmentsCompleteMatchingPredicate(ctx, id, input, LinkableEnvironmentResponseOperationPredicate{})
}

// MonitorsListLinkableEnvironmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonitorResourcesClient) MonitorsListLinkableEnvironmentsCompleteMatchingPredicate(ctx context.Context, id MonitorId, input LinkableEnvironmentRequest, predicate LinkableEnvironmentResponseOperationPredicate) (result MonitorsListLinkableEnvironmentsCompleteResult, err error) {
	items := make([]LinkableEnvironmentResponse, 0)

	resp, err := c.MonitorsListLinkableEnvironments(ctx, id, input)
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

	result = MonitorsListLinkableEnvironmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
