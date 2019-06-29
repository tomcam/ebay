package ebay

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// BrowseService handles communication with the Browse API
//
// eBay API docs: https://developer.ebay.com/api-docs/buy/browse/overview.html
type BrowseService service

// OptBrowseContextualLocation adds the header containing contextualLocation.
// It is strongly recommended that you use it when submitting Browse API methods.
//
// eBay API docs: https://developer.ebay.com/api-docs/buy/static/api-browse.html#Headers
func OptBrowseContextualLocation(country, zip string) func(*http.Request) {
	return func(req *http.Request) {
		v := req.Header.Get(headerEndUserCtx)
		if len(v) > 0 {
			v += ","
		}
		v += "contextualLocation=" + url.QueryEscape(fmt.Sprintf("country=%s,zip=%s", country, zip))
		req.Header.Set(headerEndUserCtx, v)
	}
}

// Item represents an eBay item.
type Item struct {
	AdditionalImages []struct {
		Height   string `json:"height"`
		ImageURL string `json:"imageUrl"`
		Width    string `json:"width"`
	} `json:"additionalImages"`
	AdultOnly       string   `json:"adultOnly"`
	AgeGroup        string   `json:"ageGroup"`
	BidCount        string   `json:"bidCount"`
	Brand           string   `json:"brand"`
	BuyingOptions   []string `json:"buyingOptions"`
	CategoryID      string   `json:"categoryId"`
	CategoryPath    string   `json:"categoryPath"`
	Color           string   `json:"color"`
	Condition       string   `json:"condition"`
	ConditionID     string   `json:"conditionId"`
	CurrentBidPrice struct {
		ConvertedFromCurrency string `json:"convertedFromCurrency"`
		ConvertedFromValue    string `json:"convertedFromValue"`
		Currency              string `json:"currency"`
		Value                 string `json:"value"`
	} `json:"currentBidPrice"`
	Description             string `json:"description"`
	EnabledForGuestCheckout string `json:"enabledForGuestCheckout"`
	EnergyEfficiencyClass   string `json:"energyEfficiencyClass"`
	Epid                    string `json:"epid"`
	EstimatedAvailabilities []struct {
		AvailabilityThreshold       string   `json:"availabilityThreshold"`
		AvailabilityThresholdType   string   `json:"availabilityThresholdType"`
		DeliveryOptions             []string `json:"deliveryOptions"`
		EstimatedAvailabilityStatus string   `json:"estimatedAvailabilityStatus"`
		EstimatedAvailableQuantity  string   `json:"estimatedAvailableQuantity"`
		EstimatedSoldQuantity       string   `json:"estimatedSoldQuantity"`
	} `json:"estimatedAvailabilities"`
	Gender string `json:"gender"`
	Gtin   string `json:"gtin"`
	Image  struct {
		Height   string `json:"height"`
		ImageURL string `json:"imageUrl"`
		Width    string `json:"width"`
	} `json:"image"`
	InferredEpid        string `json:"inferredEpid"`
	ItemAffiliateWebURL string `json:"itemAffiliateWebUrl"`
	ItemEndDate         string `json:"itemEndDate"`
	ItemID              string `json:"itemId"`
	ItemLocation        struct {
		AddressLine1    string `json:"addressLine1"`
		AddressLine2    string `json:"addressLine2"`
		City            string `json:"city"`
		Country         string `json:"country"`
		County          string `json:"county"`
		PostalCode      string `json:"postalCode"`
		StateOrProvince string `json:"stateOrProvince"`
	} `json:"itemLocation"`
	ItemWebURL       string `json:"itemWebUrl"`
	LocalizedAspects []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"localizedAspects"`
	MarketingPrice struct {
		DiscountAmount struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"discountAmount"`
		DiscountPercentage string `json:"discountPercentage"`
		OriginalPrice      struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"originalPrice"`
	} `json:"marketingPrice"`
	Material          string `json:"material"`
	MinimumPriceToBid struct {
		ConvertedFromCurrency string `json:"convertedFromCurrency"`
		ConvertedFromValue    string `json:"convertedFromValue"`
		Currency              string `json:"currency"`
		Value                 string `json:"value"`
	} `json:"minimumPriceToBid"`
	Mpn     string `json:"mpn"`
	Pattern string `json:"pattern"`
	Price   struct {
		ConvertedFromCurrency string `json:"convertedFromCurrency"`
		ConvertedFromValue    string `json:"convertedFromValue"`
		Currency              string `json:"currency"`
		Value                 string `json:"value"`
	} `json:"price"`
	PriceDisplayCondition string `json:"priceDisplayCondition"`
	PrimaryItemGroup      struct {
		ItemGroupAdditionalImages []struct {
			Height   string `json:"height"`
			ImageURL string `json:"imageUrl"`
			Width    string `json:"width"`
		} `json:"itemGroupAdditionalImages"`
		ItemGroupHref  string `json:"itemGroupHref"`
		ItemGroupID    string `json:"itemGroupId"`
		ItemGroupImage struct {
			Height   string `json:"height"`
			ImageURL string `json:"imageUrl"`
			Width    string `json:"width"`
		} `json:"itemGroupImage"`
		ItemGroupTitle string `json:"itemGroupTitle"`
		ItemGroupType  string `json:"itemGroupType"`
	} `json:"primaryItemGroup"`
	PrimaryProductReviewRating struct {
		AverageRating    string `json:"averageRating"`
		RatingHistograms []struct {
			Count  string `json:"count"`
			Rating string `json:"rating"`
		} `json:"ratingHistograms"`
		ReviewCount string `json:"reviewCount"`
	} `json:"primaryProductReviewRating"`
	Product struct {
		AdditionalImages []struct {
			Height   string `json:"height"`
			ImageURL string `json:"imageUrl"`
			Width    string `json:"width"`
		} `json:"additionalImages"`
		AdditionalProductIdentities []struct {
			ProductIdentity []struct {
				IdentifierType  string `json:"identifierType"`
				IdentifierValue string `json:"identifierValue"`
			} `json:"productIdentity"`
		} `json:"additionalProductIdentities"`
		AspectGroups []struct {
			Aspects []struct {
				LocalizedName   string   `json:"localizedName"`
				LocalizedValues []string `json:"localizedValues"`
			} `json:"aspects"`
			LocalizedGroupName string `json:"localizedGroupName"`
		} `json:"aspectGroups"`
		Brand       string   `json:"brand"`
		Description string   `json:"description"`
		Gtins       []string `json:"gtins"`
		Image       struct {
			Height   string `json:"height"`
			ImageURL string `json:"imageUrl"`
			Width    string `json:"width"`
		} `json:"image"`
		Mpns  []string `json:"mpns"`
		Title string   `json:"title"`
	} `json:"product"`
	ProductFicheWebURL    string `json:"productFicheWebUrl"`
	QuantityLimitPerBuyer string `json:"quantityLimitPerBuyer"`
	ReservePriceMet       string `json:"reservePriceMet"`
	ReturnTerms           struct {
		ExtendedHolidayReturnsOffered string `json:"extendedHolidayReturnsOffered"`
		RefundMethod                  string `json:"refundMethod"`
		RestockingFeePercentage       string `json:"restockingFeePercentage"`
		ReturnInstructions            string `json:"returnInstructions"`
		ReturnMethod                  string `json:"returnMethod"`
		ReturnPeriod                  struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"returnPeriod"`
		ReturnsAccepted         string `json:"returnsAccepted"`
		ReturnShippingCostPayer string `json:"returnShippingCostPayer"`
	} `json:"returnTerms"`
	Seller struct {
		FeedbackPercentage string `json:"feedbackPercentage"`
		FeedbackScore      string `json:"feedbackScore"`
		SellerAccountType  string `json:"sellerAccountType"`
		SellerLegalInfo    struct {
			Email                      string `json:"email"`
			Fax                        string `json:"fax"`
			Imprint                    string `json:"imprint"`
			LegalContactFirstName      string `json:"legalContactFirstName"`
			LegalContactLastName       string `json:"legalContactLastName"`
			Name                       string `json:"name"`
			Phone                      string `json:"phone"`
			RegistrationNumber         string `json:"registrationNumber"`
			SellerProvidedLegalAddress struct {
				AddressLine1    string `json:"addressLine1"`
				AddressLine2    string `json:"addressLine2"`
				City            string `json:"city"`
				Country         string `json:"country"`
				CountryName     string `json:"countryName"`
				County          string `json:"county"`
				PostalCode      string `json:"postalCode"`
				StateOrProvince string `json:"stateOrProvince"`
			} `json:"sellerProvidedLegalAddress"`
			TermsOfService string `json:"termsOfService"`
			VatDetails     []struct {
				IssuingCountry string `json:"issuingCountry"`
				VatID          string `json:"vatId"`
			} `json:"vatDetails"`
		} `json:"sellerLegalInfo"`
		Username string `json:"username"`
	} `json:"seller"`
	SellerItemRevision string `json:"sellerItemRevision"`
	ShippingOptions    []struct {
		AdditionalShippingCostPerUnit struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"additionalShippingCostPerUnit"`
		CutOffDateUsedForEstimate string `json:"cutOffDateUsedForEstimate"`
		MaxEstimatedDeliveryDate  string `json:"maxEstimatedDeliveryDate"`
		MinEstimatedDeliveryDate  string `json:"minEstimatedDeliveryDate"`
		QuantityUsedForEstimate   string `json:"quantityUsedForEstimate"`
		ShippingCarrierCode       string `json:"shippingCarrierCode"`
		ShippingCost              struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"shippingCost"`
		ShippingCostType              string `json:"shippingCostType"`
		ShippingServiceCode           string `json:"shippingServiceCode"`
		ShipToLocationUsedForEstimate struct {
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
		} `json:"shipToLocationUsedForEstimate"`
		TrademarkSymbol string `json:"trademarkSymbol"`
		Type            string `json:"type"`
	} `json:"shippingOptions"`
	ShipToLocations struct {
		RegionExcluded []struct {
			RegionName string `json:"regionName"`
			RegionType string `json:"regionType"`
		} `json:"regionExcluded"`
		RegionIncluded []struct {
			RegionName string `json:"regionName"`
			RegionType string `json:"regionType"`
		} `json:"regionIncluded"`
	} `json:"shipToLocations"`
	ShortDescription string `json:"shortDescription"`
	Size             string `json:"size"`
	SizeSystem       string `json:"sizeSystem"`
	SizeType         string `json:"sizeType"`
	Subtitle         string `json:"subtitle"`
	Taxes            []struct {
		EbayCollectAndRemitTax   string `json:"ebayCollectAndRemitTax"`
		IncludedInPrice          string `json:"includedInPrice"`
		ShippingAndHandlingTaxed string `json:"shippingAndHandlingTaxed"`
		TaxJurisdiction          struct {
			Region struct {
				RegionName string `json:"regionName"`
				RegionType string `json:"regionType"`
			} `json:"region"`
			TaxJurisdictionID string `json:"taxJurisdictionId"`
		} `json:"taxJurisdiction"`
		TaxPercentage string `json:"taxPercentage"`
		TaxType       string `json:"taxType"`
	} `json:"taxes"`
	Title                    string `json:"title"`
	TopRatedBuyingExperience string `json:"topRatedBuyingExperience"`
	UniqueBidderCount        string `json:"uniqueBidderCount"`
	UnitPrice                struct {
		ConvertedFromCurrency string `json:"convertedFromCurrency"`
		ConvertedFromValue    string `json:"convertedFromValue"`
		Currency              string `json:"currency"`
		Value                 string `json:"value"`
	} `json:"unitPrice"`
	UnitPricingMeasure string `json:"unitPricingMeasure"`
	Warnings           []struct {
		Category     string   `json:"category"`
		Domain       string   `json:"domain"`
		ErrorID      string   `json:"errorId"`
		InputRefIds  []string `json:"inputRefIds"`
		LongMessage  string   `json:"longMessage"`
		Message      string   `json:"message"`
		OutputRefIds []string `json:"outputRefIds"`
		Parameters   []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"parameters"`
		Subdomain string `json:"subdomain"`
	} `json:"warnings"`
}

// GetItem retrieves the details of a specific item.
// Fieldgroups valid values are "PRODUCT" and "COMPACT".
//
// eBay API docs: https://developer.ebay.com/api-docs/buy/browse/resources/item/methods/getItem
func (s *BrowseService) GetItem(ctx context.Context, itemID string, fieldgroups string, opts ...Opt) (Item, error) {
	u := fmt.Sprintf("buy/browse/v1/item/%s?fieldgroups=%s", itemID, fieldgroups)
	req, err := s.client.NewRequest(http.MethodGet, u, opts...)
	if err != nil {
		return Item{}, err
	}
	var it Item
	return it, s.client.Do(ctx, req, &it)
}
