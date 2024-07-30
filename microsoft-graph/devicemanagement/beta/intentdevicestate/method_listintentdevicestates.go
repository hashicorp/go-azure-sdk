package intentdevicestate

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

type ListIntentDeviceStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntentDeviceState
}

type ListIntentDeviceStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntentDeviceState
}

type ListIntentDeviceStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentDeviceStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentDeviceStates ...
func (c IntentDeviceStateClient) ListIntentDeviceStates(ctx context.Context, id DeviceManagementIntentId) (result ListIntentDeviceStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentDeviceStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceStates", id.ID()),
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
		Values *[]beta.DeviceManagementIntentDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentDeviceStatesComplete retrieves all the results into a single object
func (c IntentDeviceStateClient) ListIntentDeviceStatesComplete(ctx context.Context, id DeviceManagementIntentId) (ListIntentDeviceStatesCompleteResult, error) {
	return c.ListIntentDeviceStatesCompleteMatchingPredicate(ctx, id, DeviceManagementIntentDeviceStateOperationPredicate{})
}

// ListIntentDeviceStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentDeviceStateClient) ListIntentDeviceStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentId, predicate DeviceManagementIntentDeviceStateOperationPredicate) (result ListIntentDeviceStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntentDeviceState, 0)

	resp, err := c.ListIntentDeviceStates(ctx, id)
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

	result = ListIntentDeviceStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
