package userauthenticationsoftwareoathmethod

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

type ListUserByIdAuthenticationSoftwareOathMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SoftwareOathAuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationSoftwareOathMethodsCompleteResult struct {
	Items []models.SoftwareOathAuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationSoftwareOathMethods ...
func (c UserAuthenticationSoftwareOathMethodClient) ListUserByIdAuthenticationSoftwareOathMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationSoftwareOathMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/softwareOathMethods", id.ID()),
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

// ListUserByIdAuthenticationSoftwareOathMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationSoftwareOathMethodClient) ListUserByIdAuthenticationSoftwareOathMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationSoftwareOathMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx, id, models.SoftwareOathAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationSoftwareOathMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationSoftwareOathMethodClient) ListUserByIdAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SoftwareOathAuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationSoftwareOathMethodsCompleteResult, err error) {
	items := make([]models.SoftwareOathAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationSoftwareOathMethods(ctx, id)
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

	result = ListUserByIdAuthenticationSoftwareOathMethodsCompleteResult{
		Items: items,
	}
	return
}
