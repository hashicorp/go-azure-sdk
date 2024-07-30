package devicelocalcredential

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

type ListDeviceLocalCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceLocalCredentialInfo
}

type ListDeviceLocalCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceLocalCredentialInfo
}

type ListDeviceLocalCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceLocalCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceLocalCredentials ...
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentials(ctx context.Context) (result ListDeviceLocalCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceLocalCredentialsCustomPager{},
		Path:       "/directory/deviceLocalCredentials",
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
		Values *[]stable.DeviceLocalCredentialInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceLocalCredentialsComplete retrieves all the results into a single object
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentialsComplete(ctx context.Context) (ListDeviceLocalCredentialsCompleteResult, error) {
	return c.ListDeviceLocalCredentialsCompleteMatchingPredicate(ctx, DeviceLocalCredentialInfoOperationPredicate{})
}

// ListDeviceLocalCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentialsCompleteMatchingPredicate(ctx context.Context, predicate DeviceLocalCredentialInfoOperationPredicate) (result ListDeviceLocalCredentialsCompleteResult, err error) {
	items := make([]stable.DeviceLocalCredentialInfo, 0)

	resp, err := c.ListDeviceLocalCredentials(ctx)
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

	result = ListDeviceLocalCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
