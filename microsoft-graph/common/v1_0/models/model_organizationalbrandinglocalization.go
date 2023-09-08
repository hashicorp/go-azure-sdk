package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationalBrandingLocalization struct {
	BackgroundColor                   *string                          `json:"backgroundColor,omitempty"`
	BackgroundImage                   *string                          `json:"backgroundImage,omitempty"`
	BackgroundImageRelativeUrl        *string                          `json:"backgroundImageRelativeUrl,omitempty"`
	BannerLogo                        *string                          `json:"bannerLogo,omitempty"`
	BannerLogoRelativeUrl             *string                          `json:"bannerLogoRelativeUrl,omitempty"`
	CdnList                           *[]string                        `json:"cdnList,omitempty"`
	CustomAccountResetCredentialsUrl  *string                          `json:"customAccountResetCredentialsUrl,omitempty"`
	CustomCSS                         *string                          `json:"customCSS,omitempty"`
	CustomCSSRelativeUrl              *string                          `json:"customCSSRelativeUrl,omitempty"`
	CustomCannotAccessYourAccountText *string                          `json:"customCannotAccessYourAccountText,omitempty"`
	CustomCannotAccessYourAccountUrl  *string                          `json:"customCannotAccessYourAccountUrl,omitempty"`
	CustomForgotMyPasswordText        *string                          `json:"customForgotMyPasswordText,omitempty"`
	CustomPrivacyAndCookiesText       *string                          `json:"customPrivacyAndCookiesText,omitempty"`
	CustomPrivacyAndCookiesUrl        *string                          `json:"customPrivacyAndCookiesUrl,omitempty"`
	CustomResetItNowText              *string                          `json:"customResetItNowText,omitempty"`
	CustomTermsOfUseText              *string                          `json:"customTermsOfUseText,omitempty"`
	CustomTermsOfUseUrl               *string                          `json:"customTermsOfUseUrl,omitempty"`
	Favicon                           *string                          `json:"favicon,omitempty"`
	FaviconRelativeUrl                *string                          `json:"faviconRelativeUrl,omitempty"`
	HeaderBackgroundColor             *string                          `json:"headerBackgroundColor,omitempty"`
	HeaderLogo                        *string                          `json:"headerLogo,omitempty"`
	HeaderLogoRelativeUrl             *string                          `json:"headerLogoRelativeUrl,omitempty"`
	Id                                *string                          `json:"id,omitempty"`
	LoginPageLayoutConfiguration      *LoginPageLayoutConfiguration    `json:"loginPageLayoutConfiguration,omitempty"`
	LoginPageTextVisibilitySettings   *LoginPageTextVisibilitySettings `json:"loginPageTextVisibilitySettings,omitempty"`
	ODataType                         *string                          `json:"@odata.type,omitempty"`
	SignInPageText                    *string                          `json:"signInPageText,omitempty"`
	SquareLogo                        *string                          `json:"squareLogo,omitempty"`
	SquareLogoDark                    *string                          `json:"squareLogoDark,omitempty"`
	SquareLogoDarkRelativeUrl         *string                          `json:"squareLogoDarkRelativeUrl,omitempty"`
	SquareLogoRelativeUrl             *string                          `json:"squareLogoRelativeUrl,omitempty"`
	UsernameHintText                  *string                          `json:"usernameHintText,omitempty"`
}
