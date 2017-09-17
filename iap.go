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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

var appleReceipt AppleReceipt
var googleReceipt []GoogleReceipt
var APPLE_SECRET string

func init() {

}

func IsReceiptToAppleStore(platform string) bool {
	return platform == "iOS" || platform == "ios"
}

func IsReceiptToGooglePlay(platform string) bool {
	return platform == "Android" || platform == "android"
}

func ConnectToGooglePlay(request Google_IAP_Data) (error, *GooglePlayResult) {
	conf := &jwt.Config{
		Email: "test@gmail.com",

		PrivateKey: []byte(`-----BEGIN PRIVATE KEY-----
1234			
-----END PRIVATE KEY-----`),
		Scopes: []string{
			"https://www.googleapis.com/auth/androidpublisher",
		},
		TokenURL: google.JWTTokenURL,
		// If you would like to impersonate a user, you can
		// create a transport with a subject. The following GET
		// request will be made on the behalf of user@example.com.
		// Optional.
		// Subject: "user@example.com",
	}
	// Initiate an http.Client, the following GET request will be
	// authorized and authenticated on the behalf of user@example.com.
	client := conf.Client(oauth2.NoContext)

	//"https://www.googleapis.com/androidpublisher/v2/applications/com.test.iaptest/purchases/subscriptions/iap_test/tokens/aebabfcipjkjijlkodginnbc.AO-J1OxeU0GVN3AySAcp4acX7YFpCC1k0MYAZv4HnhJF7UoxRGmuPSlSYeOqrqWEGzXmw43ZpEmjEbpMbSfUw75caKaRr-3E4ndIN-eSnstwbEnF_w_qgJw"
	//connect_string := fmt.Sprintf("https://www.googleapis.com/androidpublisher/v2/applications/%s/purchases/subscriptions/%s/tokens/%s", request.PackageName, request.ProductId, request.PurchaseToken)
	connect_string := fmt.Sprintf("https://www.googleapis.com/androidpublisher/v2/applications/%s/purchases/products/%s/tokens/%s", request.PackageName, request.ProductId, request.PurchaseToken)
	res, err := client.Get(connect_string)
	if err != nil {
		log.Println("err:", err)
		return err, nil
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	appResult := &GooglePlayResult{}

	err = json.Unmarshal(body, &appResult)
	log.Printf("Receipt return %+v \n", appResult)

	time_duration, _ := strconv.ParseInt(appResult.StartTimeMillis, 10, 64)
	time_purchase := time.Unix(time_duration/1000, 0)

	time_duration, _ = strconv.ParseInt(appResult.ExpiryTimeMillis, 10, 64)
	time_expired := time.Unix(time_duration/1000, 0)

	log.Println("StartTimeMillis:", time_purchase.Local(), "\n Expired:", time_expired)
	return nil, appResult
}

func ConnectToAppleStore(sanbox bool, request AppleRequest) (error, *AppleStoreResult) {
	var url string
	if sanbox {
		log.Println("use sandbox..")
		url = "https://sandbox.itunes.apple.com/verifyReceipt"
	} else {
		log.Println("use product..")
		url = "https://buy.itunes.apple.com/verifyReceipt"
	}

	client := &http.Client{}

	jsonstr, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err:", err)
		return err, nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	appResult := &AppleStoreResult{}
	err = json.Unmarshal(body, &appResult)
	log.Printf("Receipt return %+v \n", appResult)
	if err != nil {
		log.Println("json parse error")
		return errors.New("data failed"), nil
	}
	return nil, appResult
}

func ValidaIAP(platform string) {
	var production bool
	if IsReceiptToAppleStore(platform) {
		//iOS
		appleRequest := AppleRequest{Receipt_Data: appleReceipt.Receipt_Data, Password: APPLE_SECRET}
		err, storeResult := ConnectToAppleStore(production, appleRequest)

		if err != nil || storeResult == nil {
			log.Println("Receipt failed." + err.Error())
			return
		}

		//Check receipt and return purchase date and type.
		if appleReceipt.Bundle_ID != "expect id" {
			return
		}

		findResult := false
		indexResult := -1
		for index, obj := range storeResult.ReceiptResult.IAPList {
			if obj.Product_id == appleReceipt.Product_ID && obj.Transaction_id == appleReceipt.Transaction_ID {
				findResult = true
				indexResult = index
				break
			}
		}

		if findResult == false && indexResult == -1 {
			return //ValidaIAP(user_receipt_data, db, false)
		}

		log.Println("ExpireDate :", storeResult.ReceiptResult.IAPList[indexResult].Expires_date)
		// if CheckExpiresDate(storeResult.ReceiptResult.IAPList[indexResult].Expires_date) {
		// 	//Update user if original status is subscription.
		// 	UpdateUserExpired(db, user_receipt_data.User_Id)
		// 	return nil, errors.New("The receipt already expired, update user back to free.")
		// }
	} else if IsReceiptToGooglePlay(platform) {
		//Android
		if len(googleReceipt) == 0 || googleReceipt[0].INAPP_PURCHASE_DATA.PackageName != "com.test.testapp" {
			// "Input Google receipt error."
		}

		err, google_result := ConnectToGooglePlay(googleReceipt[0].INAPP_PURCHASE_DATA)
		if err != nil || google_result == nil {
			log.Println("Receipt failed." + err.Error())
			// "Receipt failed. Err:" + err.Error()
		}

		log.Println(google_result)
		time_duration_int64, _ := strconv.ParseInt(google_result.StartTimeMillis, 10, 64)
		if time_duration_int64 == googleReceipt[0].INAPP_PURCHASE_DATA.PurchaseTime {
			//It is normally purchase and consume.
			log.Println("Purchase success, handle purchase receipt.")
		} else {
			log.Println("Purchase failed..")
			// return nil, errors.New("Receipt not complete purchase, failed.")
		}

		// 	time_purchase := GetUTCTimeZoneFromTime(time.Unix(time_duration_int64/1000, 0))
		// 	log.Println("time_purchase", time_purchase.String())

		// 	time_expired_int64, _ := strconv.ParseInt(google_result.ExpiryTimeMillis, 10, 64)
		// 	time_expired := GetUTCTimeZoneFromTime(time.Unix(time_expired_int64/1000, 0))
		// 	log.Println("time_expired", time_expired.String())

		// 	if CheckExpiresDate(time_expired.String()) {
		// 		//Update user if original status is subscription.
		// 		return nil, errors.New("The receipt already expired, update user back to free.")
		// 	}

		// 	subType_int, err = AndroidUpdateUserSubscribedByReceiptResult(db, user_receipt_data.Platform, user_receipt_data.User_Id, user_receipt_data.GooglePlayReceipt[0].INAPP_PURCHASE_DATA.ProductId, time_purchase.String(), time_expired.String())
		// 	if err != nil {
		// 		return nil, errors.New("Receipt failed. Err:" + err.Error())
		// 	}
		// 	ret.Subscribe_Type = subType_int
		// 	ret.Subscribe_Date = time_purchase.String()
		// }
	}

	return
}

func ValidaIAP_V2(platform string) (string, error) {
	var production bool
	production = true

	if IsReceiptToAppleStore(platform) {
		//iOS
		appleRequest := AppleRequest{Receipt_Data: appleReceipt.Receipt_Data, Password: APPLE_SECRET}
		err, storeResult := ConnectToAppleStore(production, appleRequest)

		if err != nil || storeResult == nil {
			log.Println("Receipt failed." + err.Error())
			return "", errors.New("Receipt failed. Err:" + err.Error())
		}

		//Check receipt and return purchase date and type.
		if appleReceipt.Bundle_ID != storeResult.ReceiptResult.Bundle_ID {
			return "", errors.New("Bundle ID is not match...")
		}

		findResult := false
		indexResult := -1
		for index, obj := range storeResult.ReceiptResult.IAPList {
			log.Println("PID:", obj.Product_id, " want:", appleReceipt.Product_ID, " find:", obj.Product_id == appleReceipt.Product_ID)
			log.Println("TID:", obj.Transaction_id, " want:", appleReceipt.Transaction_ID, " find:", obj.Transaction_id == appleReceipt.Transaction_ID)
			if obj.Product_id == appleReceipt.Product_ID && obj.Transaction_id == appleReceipt.Transaction_ID {
				findResult = true
				indexResult = index
				log.Println("Receipt:", obj)
				break
			}
		}

		if findResult == false && indexResult == -1 {
			//check aganist sandbox serve because this is how Apple rolls
			if production {
				return ValidaIAP_V2("ios/android")
			}

			return "", errors.New("Cannot find product ID in IAP list")
		}
	} else if IsReceiptToGooglePlay("Platform") {
		//Android
		if len(googleReceipt) == 0 || googleReceipt[0].INAPP_PURCHASE_DATA.PackageName != "com.test.testapp" {
			return "nil", errors.New("Input Google receipt error.")
		}

		err, google_result := ConnectToGooglePlay(googleReceipt[0].INAPP_PURCHASE_DATA)
		if err != nil || google_result == nil {
			log.Println("Receipt failed." + err.Error())
			return "nil", errors.New("Receipt failed. Err:" + err.Error())
		}

		log.Println(google_result)
		time_duration_int64, _ := strconv.ParseInt(google_result.PurchaseTimeMillis, 10, 64)
		if time_duration_int64 == googleReceipt[0].INAPP_PURCHASE_DATA.PurchaseTime {
			//It is normally purchase and consume.
			log.Println("Purchase success, handle purchase receipt.")
		} else {
			log.Println("Purchase failed..")
			return "nil", errors.New("Receipt not complete purchase, failed.")
		}
		//Update data
	}

	return "RET", nil
}
