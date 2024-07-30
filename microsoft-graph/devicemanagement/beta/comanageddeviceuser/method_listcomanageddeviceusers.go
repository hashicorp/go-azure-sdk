package comanageddeviceuser

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

type ListComanagedDeviceUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.User
}

type ListComanagedDeviceUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.User
}

type ListComanagedDeviceUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceUsers ...
func (c ComanagedDeviceUserClient) ListComanagedDeviceUsers(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceUsersCustomPager{},
		Path:       fmt.Sprintf("%s/users", id.ID()),
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
		Values *[]beta.User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceUsersComplete retrieves all the results into a single object
func (c ComanagedDeviceUserClient) ListComanagedDeviceUsersComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceUsersCompleteResult, error) {
	return c.ListComanagedDeviceUsersCompleteMatchingPredicate(ctx, id, UserOperationPredicate{})
}

// ListComanagedDeviceUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceUserClient) ListComanagedDeviceUsersCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate UserOperationPredicate) (result ListComanagedDeviceUsersCompleteResult, err error) {
	items := make([]beta.User, 0)

	resp, err := c.ListComanagedDeviceUsers(ctx, id)
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

	result = ListComanagedDeviceUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
