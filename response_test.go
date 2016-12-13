package amazonpa

import (
	"encoding/xml"
	"testing"
)

func assertEqualStr(t *testing.T, v1, v2 string, msg ...string) {
	if v1 != v2 {
		t.Error(msg)
	}
}

func assertEqualInt(t *testing.T, v1, v2 int, msg ...string) {
	if v1 != v2 {
		t.Error(msg)
	}
}

func assertEqualFloat(t *testing.T, v1, v2 float64, msg ...string) {
	if v1 != v2 {
		t.Error(msg)
	}
}

func assertEqualBool(t *testing.T, v1, v2 bool, msg ...string) {
	if v1 != v2 {
		t.Error(msg)
	}
}

func TestParseItemLookupResponse(t *testing.T) {
	const responseData = `
    <?xml version="1.0"?>
    <ItemLookupResponse xmlns="http://webservices.amazon.com/AWSECommerceService/2013-08-01">
        <OperationRequest>
            <RequestId>54b32dd2-f525-4987-9295-b1f2cda8d5f5</RequestId>
            <Arguments>
                <Argument Name="ItemId" Value="B003TGG2EA"/>
                <Argument Name="Operation" Value="ItemLookup"/>
            </Arguments>
            <RequestProcessingTime>0.0837001980000000</RequestProcessingTime>
        </OperationRequest>
        <Items>
            <Request>
                <IsValid>True</IsValid>
                <ItemLookupRequest>
                    <IdType>ASIN</IdType>
                    <ItemId>B003TGG2EA</ItemId>
                    <ResponseGroup>Large</ResponseGroup>
                    <VariationPage>All</VariationPage>
                </ItemLookupRequest>
            </Request>
            <Item>
                <ASIN>B003TGG2EA</ASIN>
                <DetailPageURL>https://www.amazon.it/Grohe-32843000-Cosmopolitan-Miscelatore-Monocomando/dp/B003TGG2EA%3FSubscriptionId%3DAKIAIZL74FSKHXDX66WQ%26tag%3Dgoldbot-21%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB003TGG2EA</DetailPageURL>
                <ItemLinks>
                    <ItemLink>
                        <Description>Add To Wishlist</Description>
                        <URL>https://www.amazon.it/gp/registry/wishlist/add-item.html%3Fasin.0%3DB003TGG2EA%26SubscriptionId%3DAKIAIZL74FSKHXDX66WQ%26tag%3Dgoldbot-21%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB003TGG2EA</URL>
                    </ItemLink>
                </ItemLinks>
                <SalesRank>214</SalesRank>
                <SmallImage>
                    <URL>http://ecx.images-amazon.com/images/I/3143fcCtWnL._SL75_.jpg</URL>
                    <Height Units="pixels">75</Height>
                    <Width Units="pixels">52</Width>
                </SmallImage>
                <MediumImage>
                    <URL>http://ecx.images-amazon.com/images/I/3143fcCtWnL._SL160_.jpg</URL>
                    <Height Units="pixels">160</Height>
                    <Width Units="pixels">110</Width>
                </MediumImage>
                <LargeImage>
                    <URL>http://ecx.images-amazon.com/images/I/3143fcCtWnL.jpg</URL>
                    <Height Units="pixels">500</Height>
                    <Width Units="pixels">344</Width>
                </LargeImage>
                <ImageSets>
                    <ImageSet Category="primary">
                        <SwatchImage>
                            <URL>http://ecx.images-amazon.com/images/I/3143fcCtWnL._SL30_.jpg</URL>
                            <Height Units="pixels">30</Height>
                            <Width Units="pixels">21</Width>
                        </SwatchImage>
                    </ImageSet>
                </ImageSets>
                <ItemAttributes>
                    <Binding>Tools &amp; Home Improvement</Binding>
                    <Brand>Grohe</Brand>
                    <CatalogNumberList>
                        <CatalogNumberListElement>32843000-FG</CatalogNumberListElement>
                        <CatalogNumberListElement>Z00401109</CatalogNumberListElement>
                    </CatalogNumberList>
                    <Color>Cromo</Color>
                    <EAN>4005176874840</EAN>
                    <EANList>
                        <EANListElement>4005176874840</EANListElement>
                    </EANList>
                    <Feature>Comfort superiore grazie alla bocca alta di erogazione</Feature>
                    <Feature>Rotazione morbida della bocca di erogazione</Feature>
                    <Feature>Cromatura Grohe StarLight per una brillantezza inalterata nel tempo</Feature>
                    <Feature>Installazione Semplice e Veloce</Feature>
                    <Feature>Angolo di rotazione della bocca di erogazione di 0/150/360 gradi</Feature>
                    <ItemDimensions>
                        <Height Units="Centesimi pollici">528</Height>
                        <Length Units="Centesimi pollici">1437</Length>
                        <Weight Units="Centesimi libbre">100</Weight>
                        <Width Units="Centesimi pollici">795</Width>
                    </ItemDimensions>
                    <Label>Grohe</Label>
                    <ListPrice>
                        <Amount>18500</Amount>
                        <CurrencyCode>EUR</CurrencyCode>
                        <FormattedPrice>EUR 185,00</FormattedPrice>
                    </ListPrice>
                    <Manufacturer>Grohe</Manufacturer>
                    <Model>32843000</Model>
                    <MPN>32843000</MPN>
                    <PackageDimensions>
                        <Height Units="Centesimi pollici">520</Height>
                        <Length Units="Centesimi pollici">1449</Length>
                        <Weight Units="Centesimi libbre">476</Weight>
                        <Width Units="Centesimi pollici">819</Width>
                    </PackageDimensions>
                    <PackageQuantity>1</PackageQuantity>
                    <PartNumber>32843000</PartNumber>
                    <ProductGroup>Bricolage</ProductGroup>
                    <ProductTypeName>TOOLS</ProductTypeName>
                    <Publisher>Grohe</Publisher>
                    <ReleaseDate>2015-06-30</ReleaseDate>
                    <Studio>Grohe</Studio>
                    <Title>Grohe 32843000 Eurosmart Cosmopolitan Miscelatore Monocomando Lavello, Cromo</Title>
                    <Warranty>2 anni</Warranty>
                </ItemAttributes>
                <OfferSummary>
                    <LowestNewPrice>
                        <Amount>6490</Amount>
                        <CurrencyCode>EUR</CurrencyCode>
                        <FormattedPrice>EUR 64,90</FormattedPrice>
                    </LowestNewPrice>
                    <TotalNew>16</TotalNew>
                    <TotalUsed>0</TotalUsed>
                    <TotalCollectible>0</TotalCollectible>
                    <TotalRefurbished>0</TotalRefurbished>
                </OfferSummary>
                <Offers>
                    <TotalOffers>1</TotalOffers>
                    <TotalOfferPages>1</TotalOfferPages>
                    <MoreOffersUrl>https://www.amazon.it/gp/offer-listing/B003TGG2EA%3FSubscriptionId%3DAKIAIZL74FSKHXDX66WQ%26tag%3Dgoldbot-21%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB003TGG2EA</MoreOffersUrl>
                    <Offer>
                        <OfferAttributes>
                            <Condition>New</Condition>
                        </OfferAttributes>
                        <OfferListing>
                            <OfferListingId>q1yue46%2Bp%2B5TjX%2BbCjq7jXnSFu0XuDFm0ZJTb3Mwv8epyiw3I1Rk4S1j9ceypzivpkkkaaQ7vzBSoJddMvmpPhN8RQeiqDm6WVi2YIj5WFwd6DIWJkiruQ%3D%3D</OfferListingId>
                            <Price>
                                <Amount>6490</Amount>
                                <CurrencyCode>EUR</CurrencyCode>
                                <FormattedPrice>EUR 64,90</FormattedPrice>
                            </Price>
                            <AmountSaved>
                                <Amount>12010</Amount>
                                <CurrencyCode>EUR</CurrencyCode>
                                <FormattedPrice>EUR 120,10</FormattedPrice>
                            </AmountSaved>
                            <PercentageSaved>65</PercentageSaved>
                            <Availability>Generalmente spedito in 24 ore</Availability>
                            <AvailabilityAttributes>
                                <AvailabilityType>now</AvailabilityType>
                                <MinimumHours>0</MinimumHours>
                                <MaximumHours>0</MaximumHours>
                            </AvailabilityAttributes>
                            <IsEligibleForSuperSaverShipping>1</IsEligibleForSuperSaverShipping>
                            <IsEligibleForPrime>1</IsEligibleForPrime>
                        </OfferListing>
                    </Offer>
                </Offers>
                <CustomerReviews>
                    <IFrameURL>https://www.amazon.it/reviews/iframe?akid=AKIAIZL74FSKHXDX66WQ&amp;alinkCode=xm2&amp;asin=B003TGG2EA&amp;atag=goldbot-21&amp;exp=2016-12-14T06%3A32%3A23Z&amp;v=2&amp;sig=G1XUpkUhFIGmT5OeUO1OyAy4T4YrYCpxDFNVR%2FYic5Q%3D</IFrameURL>
                    <HasReviews>true</HasReviews>
                </CustomerReviews>
                <EditorialReviews>
                    <EditorialReview>
                        <Source>Product Description</Source>
                        <Content>Grohe Eurosmart Cosmopolitan Miscelatore monocomando&lt;br&gt;per lavello&lt;br&gt;bocca alta&lt;br&gt;cartuccia a dischi ceramici da 35 mm con tecnologia SilkMove &lt;br&gt;dotato di serie di limitatore di portata regolabile&lt;br&gt;bocca girevole&lt;br&gt;angolo di rotazione selezionabile tra: 0&#xB0; / 150&#xB0; / 360&#xB0;&lt;/p&gt;&lt;br&gt;&#xA0;&lt;/p&gt;&lt;br&gt;&#xA0;&lt;/p&gt;</Content>
                        <IsLinkSuppressed>0</IsLinkSuppressed>
                    </EditorialReview>
                </EditorialReviews>
                <SimilarProducts>
                    <SimilarProduct>
                        <ASIN>B00RTG0DZK</ASIN>
                        <Title>Grohe 33300002 Eurosmart New Miscelatore Monocomando per Vasca/Doccia, Cromo</Title>
                    </SimilarProduct>
                </SimilarProducts>
                <BrowseNodes>
                    <BrowseNode>
                        <BrowseNodeId>3120323031</BrowseNodeId>
                        <Name>Rubinetti per lavelli da cucina</Name>
                        <Ancestors>
                            <BrowseNode>
                                <BrowseNodeId>3119756031</BrowseNodeId>
                                <Name>Rubinetti da cucina</Name>
                                <Ancestors>
                                    <BrowseNode>
                                        <BrowseNodeId>3119621031</BrowseNodeId>
                                        <Name>Impianti per la cucina</Name>
                                        <Ancestors>
                                            <BrowseNode>
                                                <BrowseNodeId>3119607031</BrowseNodeId>
                                                <Name>Attrezzature per cucine e bagni</Name>
                                                <Ancestors>
                                                    <BrowseNode>
                                                        <BrowseNodeId>2454161031</BrowseNodeId>
                                                        <Name>Categorie</Name>
                                                        <IsCategoryRoot>1</IsCategoryRoot>
                                                        <Ancestors>
                                                            <BrowseNode>
                                                                <BrowseNodeId>2454160031</BrowseNodeId>
                                                                <Name>Fai da te</Name>
                                                            </BrowseNode>
                                                        </Ancestors>
                                                    </BrowseNode>
                                                </Ancestors>
                                            </BrowseNode>
                                        </Ancestors>
                                    </BrowseNode>
                                </Ancestors>
                            </BrowseNode>
                        </Ancestors>
                    </BrowseNode>
                </BrowseNodes>
            </Item>
        </Items>
    </ItemLookupResponse>
    `
	var response ItemLookupResponse
	xml.Unmarshal([]byte(responseData), &response)

	// OperationRequest
	assertEqualStr(t, response.OperationRequest.RequestID, "54b32dd2-f525-4987-9295-b1f2cda8d5f5", "Bad RequestID")
	assertEqualStr(t, response.OperationRequest.Arguments[0].Name, "ItemId", "Bad Argument Name")
	assertEqualStr(t, response.OperationRequest.Arguments[0].Value, "B003TGG2EA", "Bad Argument Value")
	assertEqualFloat(t, response.OperationRequest.RequestProcessingTime, 0.0837001980000000, "Bad RequestProcessingTime")

	// Items/Request
	assertEqualBool(t, response.Items.Request.IsValid, true, "Bad IsValid")
	assertEqualStr(t, response.Items.Request.ItemLookupRequest.IDType, "ASIN", "Bad IDType")
	assertEqualStr(t, response.Items.Request.ItemLookupRequest.ItemID, "B003TGG2EA", "Bad ItemID")
	assertEqualStr(t, response.Items.Request.ItemLookupRequest.ResponseGroup, "Large", "Bad ResponseGroup")
	assertEqualStr(t, response.Items.Request.ItemLookupRequest.VariationPage, "All", "Bad VariationPage")

	//Items/Item
	assertEqualStr(t, response.Items.Item.ASIN, "B003TGG2EA", "Bad ASIN")
	assertEqualStr(t, response.Items.Item.DetailPageURL, "https://www.amazon.it/Grohe-32843000-Cosmopolitan-Miscelatore-Monocomando/dp/B003TGG2EA%3FSubscriptionId%3DAKIAIZL74FSKHXDX66WQ%26tag%3Dgoldbot-21%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3DB003TGG2EA", "Bad DetailPageURL")
	assertEqualInt(t, response.Items.Item.SalesRank, 214, "Bad SalesRank")

	assertEqualStr(t, response.Items.Item.SmallImage.URL, "http://ecx.images-amazon.com/images/I/3143fcCtWnL._SL75_.jpg", "Bad SmallImage/URL")
	assertEqualInt(t, int(response.Items.Item.SmallImage.Height), 75, "Bad SmallImage/Height")
	assertEqualInt(t, int(response.Items.Item.SmallImage.Width), 52, "Bad SmallImage/Width")

	assertEqualStr(t, response.Items.Item.MediumImage.URL, "http://ecx.images-amazon.com/images/I/3143fcCtWnL._SL160_.jpg", "Bad SmallImage/URL")
	assertEqualInt(t, int(response.Items.Item.MediumImage.Height), 160, "Bad SmallImage/Height")
	assertEqualInt(t, int(response.Items.Item.MediumImage.Width), 110, "Bad SmallImage/Width")

	assertEqualStr(t, response.Items.Item.LargeImage.URL, "http://ecx.images-amazon.com/images/I/3143fcCtWnL.jpg", "Bad SmallImage/URL")
	assertEqualInt(t, int(response.Items.Item.LargeImage.Height), 500, "Bad SmallImage/Height")
	assertEqualInt(t, int(response.Items.Item.LargeImage.Width), 344, "Bad SmallImage/Width")

	// Items/Item/ItemAttributes
	assertEqualStr(t, response.Items.Item.ItemAttributes.Binding, "Tools & Home Improvement", "Bad Binding")
	assertEqualStr(t, response.Items.Item.ItemAttributes.Brand, "Grohe", "Bad Binding")
	assertEqualStr(t, response.Items.Item.ItemAttributes.Color, "Cromo", "Bad Color")
	assertEqualStr(t, response.Items.Item.ItemAttributes.Creator, "", "Bad Creator")
	assertEqualStr(t, response.Items.Item.ItemAttributes.EAN, "4005176874840", "Bad EAN")

	assertEqualInt(t, int(response.Items.Item.ItemAttributes.ListPrice.Amount), 18500, "Bad Price/Amount")
	assertEqualStr(t, response.Items.Item.ItemAttributes.ListPrice.CurrencyCode, "EUR", "Bad Price/Currency")
	assertEqualStr(t, response.Items.Item.ItemAttributes.ListPrice.FormattedPrice, "EUR 185,00", "Bad Price/Formatted Price")

	assertEqualStr(t, response.Items.Item.BrowseNodes.BrowseNode[0].BrowseNodeID, "3120323031", "Bad BrowseNode/BrowseNodeID")
	assertEqualStr(t, response.Items.Item.BrowseNodes.BrowseNode[0].Name, "Rubinetti per lavelli da cucina", "Bad BrowseNode/Name")

	assertEqualStr(t, response.Items.Item.BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].BrowseNodeID, "3119756031", "Bad Ancestors/BrowseNodeID")
	assertEqualStr(t, response.Items.Item.BrowseNodes.BrowseNode[0].Ancestors.BrowseNode[0].Name, "Rubinetti da cucina", "Bad Ancestors/Name")
}
