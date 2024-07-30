package derivedcredential

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

type ListDerivedCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementDerivedCredentialSettings
}

type ListDerivedCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementDerivedCredentialSettings
}

type ListDerivedCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDerivedCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDerivedCredentials ...
func (c DerivedCredentialClient) ListDerivedCredentials(ctx context.Context) (result ListDerivedCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDerivedCredentialsCustomPager{},
		Path:       "/deviceManagement/derivedCredentials",
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
		Values *[]beta.DeviceManagementDerivedCredentialSettings `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDerivedCredentialsComplete retrieves all the results into a single object
func (c DerivedCredentialClient) ListDerivedCredentialsComplete(ctx context.Context) (ListDerivedCredentialsCompleteResult, error) {
	return c.ListDerivedCredentialsCompleteMatchingPredicate(ctx, DeviceManagementDerivedCredentialSettingsOperationPredicate{})
}

// ListDerivedCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DerivedCredentialClient) ListDerivedCredentialsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementDerivedCredentialSettingsOperationPredicate) (result ListDerivedCredentialsCompleteResult, err error) {
	items := make([]beta.DeviceManagementDerivedCredentialSettings, 0)

	resp, err := c.ListDerivedCredentials(ctx)
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

	result = ListDerivedCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
