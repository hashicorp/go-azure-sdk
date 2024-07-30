package manageddeviceencryptionstate

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

type ListManagedDeviceEncryptionStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceEncryptionState
}

type ListManagedDeviceEncryptionStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceEncryptionState
}

type ListManagedDeviceEncryptionStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceEncryptionStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceEncryptionStates ...
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStates(ctx context.Context) (result ListManagedDeviceEncryptionStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceEncryptionStatesCustomPager{},
		Path:       "/deviceManagement/managedDeviceEncryptionStates",
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
		Values *[]beta.ManagedDeviceEncryptionState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceEncryptionStatesComplete retrieves all the results into a single object
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStatesComplete(ctx context.Context) (ListManagedDeviceEncryptionStatesCompleteResult, error) {
	return c.ListManagedDeviceEncryptionStatesCompleteMatchingPredicate(ctx, ManagedDeviceEncryptionStateOperationPredicate{})
}

// ListManagedDeviceEncryptionStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStatesCompleteMatchingPredicate(ctx context.Context, predicate ManagedDeviceEncryptionStateOperationPredicate) (result ListManagedDeviceEncryptionStatesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceEncryptionState, 0)

	resp, err := c.ListManagedDeviceEncryptionStates(ctx)
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

	result = ListManagedDeviceEncryptionStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
