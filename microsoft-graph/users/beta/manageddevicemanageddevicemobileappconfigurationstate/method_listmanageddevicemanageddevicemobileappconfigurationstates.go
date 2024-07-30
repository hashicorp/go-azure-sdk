package manageddevicemanageddevicemobileappconfigurationstate

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

type ListManagedDeviceManagedDeviceMobileAppConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceMobileAppConfigurationState
}

type ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceMobileAppConfigurationState
}

type ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceManagedDeviceMobileAppConfigurationStates ...
func (c ManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListManagedDeviceManagedDeviceMobileAppConfigurationStates(ctx context.Context, id UserIdManagedDeviceId) (result ListManagedDeviceManagedDeviceMobileAppConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager{},
		Path:       fmt.Sprintf("%s/managedDeviceMobileAppConfigurationStates", id.ID()),
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
		Values *[]beta.ManagedDeviceMobileAppConfigurationState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceManagedDeviceMobileAppConfigurationStatesComplete retrieves all the results into a single object
func (c ManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListManagedDeviceManagedDeviceMobileAppConfigurationStatesComplete(ctx context.Context, id UserIdManagedDeviceId) (ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult, error) {
	return c.ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx, id, ManagedDeviceMobileAppConfigurationStateOperationPredicate{})
}

// ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id UserIdManagedDeviceId, predicate ManagedDeviceMobileAppConfigurationStateOperationPredicate) (result ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceMobileAppConfigurationState, 0)

	resp, err := c.ListManagedDeviceManagedDeviceMobileAppConfigurationStates(ctx, id)
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

	result = ListManagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
