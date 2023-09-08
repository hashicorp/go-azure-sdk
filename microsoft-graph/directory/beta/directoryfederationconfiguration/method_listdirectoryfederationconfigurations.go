package directoryfederationconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryFederationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityProviderBaseCollectionResponse
}

type ListDirectoryFederationConfigurationsCompleteResult struct {
	Items []models.IdentityProviderBaseCollectionResponse
}

// ListDirectoryFederationConfigurations ...
func (c DirectoryFederationConfigurationClient) ListDirectoryFederationConfigurations(ctx context.Context) (result ListDirectoryFederationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/federationConfigurations",
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
		Values *[]models.IdentityProviderBaseCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryFederationConfigurationsComplete retrieves all the results into a single object
func (c DirectoryFederationConfigurationClient) ListDirectoryFederationConfigurationsComplete(ctx context.Context) (ListDirectoryFederationConfigurationsCompleteResult, error) {
	return c.ListDirectoryFederationConfigurationsCompleteMatchingPredicate(ctx, models.IdentityProviderBaseCollectionResponseOperationPredicate{})
}

// ListDirectoryFederationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryFederationConfigurationClient) ListDirectoryFederationConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate models.IdentityProviderBaseCollectionResponseOperationPredicate) (result ListDirectoryFederationConfigurationsCompleteResult, err error) {
	items := make([]models.IdentityProviderBaseCollectionResponse, 0)

	resp, err := c.ListDirectoryFederationConfigurations(ctx)
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

	result = ListDirectoryFederationConfigurationsCompleteResult{
		Items: items,
	}
	return
}
