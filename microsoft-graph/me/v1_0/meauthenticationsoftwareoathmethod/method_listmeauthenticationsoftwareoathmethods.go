package meauthenticationsoftwareoathmethod

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

type ListMeAuthenticationSoftwareOathMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SoftwareOathAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationSoftwareOathMethodsCompleteResult struct {
	Items []models.SoftwareOathAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationSoftwareOathMethods ...
func (c MeAuthenticationSoftwareOathMethodClient) ListMeAuthenticationSoftwareOathMethods(ctx context.Context) (result ListMeAuthenticationSoftwareOathMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/softwareOathMethods",
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
		Values *[]models.SoftwareOathAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeAuthenticationSoftwareOathMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationSoftwareOathMethodClient) ListMeAuthenticationSoftwareOathMethodsComplete(ctx context.Context) (ListMeAuthenticationSoftwareOathMethodsCompleteResult, error) {
	return c.ListMeAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx, models.SoftwareOathAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationSoftwareOathMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationSoftwareOathMethodClient) ListMeAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.SoftwareOathAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationSoftwareOathMethodsCompleteResult, err error) {
	items := make([]models.SoftwareOathAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationSoftwareOathMethods(ctx)
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

	result = ListMeAuthenticationSoftwareOathMethodsCompleteResult{
		Items: items,
	}
	return
}
