package ransomwarereports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RansomwareReport
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RansomwareReport
}

type ListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// List ...
func (c RansomwareReportsClient) List(ctx context.Context, id VolumeId) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCustomPager{},
		Path:       fmt.Sprintf("%s/ransomwareReports", id.ID()),
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
		Values *[]RansomwareReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c RansomwareReportsClient) ListComplete(ctx context.Context, id VolumeId) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, RansomwareReportOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RansomwareReportsClient) ListCompleteMatchingPredicate(ctx context.Context, id VolumeId, predicate RansomwareReportOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]RansomwareReport, 0)

	resp, err := c.List(ctx, id)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
