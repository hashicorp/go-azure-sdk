package directorycertificateauthoritycertificatebasedapplicationconfiguration

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

type ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CertificateBasedApplicationConfigurationCollectionResponse
}

type ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult struct {
	Items []models.CertificateBasedApplicationConfigurationCollectionResponse
}

// ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurations ...
func (c DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient) ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurations(ctx context.Context) (result ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/certificateAuthorities/certificateBasedApplicationConfigurations",
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
		Values *[]models.CertificateBasedApplicationConfigurationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsComplete retrieves all the results into a single object
func (c DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient) ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsComplete(ctx context.Context) (ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult, error) {
	return c.ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate(ctx, models.CertificateBasedApplicationConfigurationCollectionResponseOperationPredicate{})
}

// ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient) ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate models.CertificateBasedApplicationConfigurationCollectionResponseOperationPredicate) (result ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult, err error) {
	items := make([]models.CertificateBasedApplicationConfigurationCollectionResponse, 0)

	resp, err := c.ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurations(ctx)
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

	result = ListDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult{
		Items: items,
	}
	return
}
