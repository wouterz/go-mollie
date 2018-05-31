// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

// Mollie holds references to all sections of the API
// use the access functions to retrieve instances of this structure
type Mollie struct {
	issuers   *IssuerAPI
	methods   *MethodAPI
	payments  *PaymentAPI
	customers *CustomerAPI
}

type OauthMollie struct {
	connects *ConnectAPI
}

// Get generates a new API structure with the provided API-Key
func Get(apikey string) Mollie {
	c := core{apiKey: apikey}

	return Mollie{
		issuers:   newIssuers(&c)
		methods:   newMethods(&c),
		payments:  newPayments(&c),
		customers: newCustomers(&c)
	}
}

func GetOauth(clientID, clientSecret, redirect string, scopes ...string) {
	return OauthMollie{
		connects: newConnects(clientID, clientSecret, redirect, scopes),
	}
}

// Issuers returns a reference to the IssuerAPI
func (m Mollie) Issuers() *IssuerAPI {
	return m.issuers
}

// Methods returns a reference to the MethodAPI
func (m Mollie) Methods() *MethodAPI {
	return m.methods
}

// Payments returns a reference to the PaymentAPI
func (m Mollie) Payments() *PaymentAPI {
	return m.payments
}

// Customers returns a reference to the CustomerAPI
func (m Mollie) Customers() *CustomerAPI {
	return m.customers
}

// Connects returns a reference to the ConnectAPI
func (m Mollie) Connects() *ConnectAPI {
	return m.connects
}
