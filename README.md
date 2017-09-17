IAP: Golang In-App-Purchase Package for Apple iTune and Google Play
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/iap/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/iap?status.svg)](https://godoc.org/github.com/kkdai/iap)  [![Build Status](https://travis-ci.org/kkdai/iap.svg?branch=master)](https://travis-ci.org/kkdai/iap)


### IAP


Install
---------------
`go get github.com/kkdai/iap`


Usage
---------------

```go
package main

import (
	"fmt"

	. "github.com/kkdai/iap"
)

func main() {
	//WIP
}
```


Want to know more detail about Google Play IAP
---------------

Refer my blog [here](https://golang.kktix.cc/events/gtg27)

Want to know more detail about Apple iTune IAP
---------------

Refer my blog [here](http://www.evanlin.com/server-side-iap-verification-apple-store/)

Reference (Android Google Play)
---------------

- [OAuth2 server account document](https://developers.google.com/accounts/docs/OAuth2ServiceAccount) 
- Using [Google Developer API - IAP API](http://developer.android.com/google/play/billing/gp-purchase-status-api.html#subscriptions_api_overview) to connecting to Google Play using OAuth and REST communication.
- [http://developer.android.com/google/play/billing/billing_integrate.html](http://developer.android.com/google/play/billing/billing_integrate.html)
- [http://stackoverflow.com/questions/16067180/server-side-verification-of-google-play-in-app-billing-version-3-purchase](http://stackoverflow.com/questions/16067180/server-side-verification-of-google-play-in-app-billing-version-3-purchase)
- [http://developer.android.com/google/play/billing/billing_integrate.html#Purchase](http://developer.android.com/google/play/billing/billing_integrate.html#Purchase)
- [https://developers.google.com/android-publisher/getting_started](https://developers.google.com/android-publisher/getting_started)
- [https://developers.google.com/accounts/docs/OAuth2](https://developers.google.com/accounts/docs/OAuth2)
- [https://developers.google.com/accounts/docs/OAuth2ServiceAccount](https://developers.google.com/accounts/docs/OAuth2ServiceAccount)
- [https://developer.android.com/google/play/billing/gp-purchase-status-api.html#using](https://developer.android.com/google/play/billing/gp-purchase-status-api.html#using)
- [http://robertomurray.co.uk/blog/2013/server-side-google-play-in-app-billing-receipt-validation-and-testing/](http://robertomurray.co.uk/blog/2013/server-side-google-play-in-app-billing-receipt-validation-and-testing/)



Reference (Apple iTune)
---------------

- [Validating Receipts With the App Store](https://developer.apple.com/library/ios/releasenotes/General/ValidateAppStoreReceipt/Chapters/ValidateRemotely.html)
- [iOS7 - receipts not validating at sandbox - error 21002 (java.lang.IllegalArgumentException)](http://stackoverflow.com/questions/19222845/ios7-receipts-not-validating-at-sandbox-error-21002-java-lang-illegalargume/19963192#19963192)
- [Validate Mac App Store receipt server side](http://stackoverflow.com/questions/22957165/validate-mac-app-store-receipt-server-side)
- [Validating iOS In-App Purchases With Laravel](http://blog.goforyt.com/validating-ios-app-purchases-laravel/)


License
---------------

This package is licensed under MIT license. See LICENSE for details.

