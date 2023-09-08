package identityapiconnector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIdentityApiConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityApiConnectorCollectionResponse
}

type ListIdentityApiConnectorsCompleteResult struct {
	Items []models.IdentityApiConnectorCollectionResponse
}

// ListIdentityApiConnectors ...
func (c IdentityApiConnectorClient) ListIdentityApiConnectors(ctx context.Context) (result ListIdentityApiConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/apiConnectors",
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
		Values *[]models.IdentityApiConnectorCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityApiConnectorsComplete retrieves all the results into a single object
func (c IdentityApiConnectorClient) ListIdentityApiConnectorsComplete(ctx context.Context) (ListIdentityApiConnectorsCompleteResult, error) {
	return c.ListIdentityApiConnectorsCompleteMatchingPredicate(ctx, models.IdentityApiConnectorCollectionResponseOperationPredicate{})
}

// ListIdentityApiConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityApiConnectorClient) ListIdentityApiConnectorsCompleteMatchingPredicate(ctx context.Context, predicate models.IdentityApiConnectorCollectionResponseOperationPredicate) (result ListIdentityApiConnectorsCompleteResult, err error) {
	items := make([]models.IdentityApiConnectorCollectionResponse, 0)

	resp, err := c.ListIdentityApiConnectors(ctx)
	if err != nil {
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

	result = ListIdentityApiConnectorsCompleteResult{
		Items: items,
	}
	return
}
