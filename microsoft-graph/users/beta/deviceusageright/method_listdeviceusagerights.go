package deviceusageright

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

type ListDeviceUsageRightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UsageRight
}

type ListDeviceUsageRightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UsageRight
}

type ListDeviceUsageRightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceUsageRightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceUsageRights ...
func (c DeviceUsageRightClient) ListDeviceUsageRights(ctx context.Context, id UserIdDeviceId) (result ListDeviceUsageRightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceUsageRightsCustomPager{},
		Path:       fmt.Sprintf("%s/usageRights", id.ID()),
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
		Values *[]beta.UsageRight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceUsageRightsComplete retrieves all the results into a single object
func (c DeviceUsageRightClient) ListDeviceUsageRightsComplete(ctx context.Context, id UserIdDeviceId) (ListDeviceUsageRightsCompleteResult, error) {
	return c.ListDeviceUsageRightsCompleteMatchingPredicate(ctx, id, UsageRightOperationPredicate{})
}

// ListDeviceUsageRightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceUsageRightClient) ListDeviceUsageRightsCompleteMatchingPredicate(ctx context.Context, id UserIdDeviceId, predicate UsageRightOperationPredicate) (result ListDeviceUsageRightsCompleteResult, err error) {
	items := make([]beta.UsageRight, 0)

	resp, err := c.ListDeviceUsageRights(ctx, id)
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

	result = ListDeviceUsageRightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
