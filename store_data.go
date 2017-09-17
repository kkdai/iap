//The MIT License (MIT)
//Copyright (c) 2015 Evan Lin
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package iap

type AppleIAP struct {
	Quantity        string `json:"quantity"`
	Product_id      string `json:"product_id"`
	Transaction_id  string `json:"transaction_id"`
	Is_trial_period string `json:"is_trial_period"`
	Purchase_date   string `json:"purchase_date"`
	Expires_date    string `json:"expires_date"`
}

type AppleIAP_Check struct {
	Apple_iap      AppleIAP
	Transaction_id string
	User_id        string
	Server_Time    string
}

/*
{
	"status":0,
	"environment":"Sandbox",
	"receipt":
		{
			"receipt_type":"ProductionSandbox",
			"adam_id":0, "app_item_id":0,
			"bundle_id":"com.test.testProd",
			"application_version":"1",
			"download_id":0,
			"version_external_identifier":0,
			"request_date":"2015-02-07 03:23:48 Etc/GMT",
			"request_date_ms":"1423279428347",
			"request_date_pst":"2015-02-06 19:23:48 America/Los_Angeles",
			"original_purchase_date":"2013-08-01 07:00:00 Etc/GMT",
			"original_purchase_date_ms":"1375340400000",
			"original_purchase_date_pst":"2013-08-01 00:00:00 America/Los_Angeles",
			"original_application_version":"1.0",
			"in_app": [ {}
		}
*/
type AppleReceiptResult struct {
	Receipt_Type        string     `json:"receipt_type"`
	Bundle_ID           string     `json:"bundle_id"`
	Application_Version string     `json:"application_version"`
	IAPList             []AppleIAP `json:"in_app"`
}

type AppleStoreResult struct {
	Status        int                `json:"status"`
	Environment   string             `json:"environment"`
	ReceiptResult AppleReceiptResult `json:"receipt"`
}

// {
//  "kind": "androidpublisher#subscriptionPurchase",
//  "startTimeMillis": "1426484398029",
//  "expiryTimeMillis": "1426657148051",
//  "autoRenewing": true
// }
//GooglePlayResult :
type GooglePlayResult struct {
	Kind               string `json:"kind"`
	StartTimeMillis    string `json:"startTimeMillis"`
	ExpiryTimeMillis   string `json:"expiryTimeMillis"`
	PurchaseTimeMillis string `json:"purchaseTimeMillis"`
	AutoRenewing       bool   `json:"autoRenewing"`
}

//AppleRequest apple store request template
type AppleRequest struct {
	Receipt_Data string `json:"receipt-data"`
	Password     string `json:"password,omitempty"`
}

//AppleReceipt :
type AppleReceipt struct {
	Bundle_ID      string `json:"bundle_id"`
	Product_ID     string `json:"product_id"`
	Transaction_ID string `json:"transaction_id"`
	Receipt_Data   string `json:"receipt-data"`
}

/*
[
  {
    INAPP_PURCHASE_DATA : {
      "orderId": "1234",
      "packageName": "com.test.iaptest",
      "productId": "iap_test",
      "purchaseTime": 1234,
      "purchaseState": 0,
      "developerPayload": "TEST_DEV_PAYLOAD",
      "purchaseToken": ""
    },

    INAPP_DATA_SIGNATURE : "1234",
    RESPONSE_CODE : 0
  }
]
*/

//Google_IAP_Data :
type Google_IAP_Data struct {
	OrderId          string `json:"orderId"`
	PackageName      string `json:"packageName"`
	ProductId        string `json:"productId"`
	PurchaseTime     int64  `json:"purchaseTime"`
	PurchaseState    int    `json:"purchaseState"`
	DeveloperPayload string `json:"developerPayload"`
	PurchaseToken    string `json:"purchaseToken"`
}

//GoogleIAP_Check :
type GoogleIAP_Check struct {
	Google_iap    Google_IAP_Data
	PurchaseToken string
	User_id       string
	Server_Time   string
}

//GoogleReceipt :
type GoogleReceipt struct {
	INAPP_PURCHASE_DATA  Google_IAP_Data `json:"inapp_purchase_data"`
	INAPP_DATA_SIGNATURE string          `json:"inapp_data_sigature"`
	RESPONSE_CODE        int             `json:"response_code"`
}
