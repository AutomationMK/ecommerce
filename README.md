# Ecommerce Web Application

## Description

This is a simple web application for handling stripe payments
A go backend is used and is inspired by following along
tsawler's go web applications intermediate udemy course.

This project is mostly focused on showcasing the golang 
backend features including the following backend methods...
* Creation and handling of http routes
* Creation and implementation of Middleware for routes
* Serving static files to the website
* Implementing a template rendering function for go tmpl files
* Integrating javascript and css to be served as static files
* Writing tests for all backend functions
* Connecting to a Postgress database to persist user data
* Implementing sending mail notifications
* Authenticating users and setting up an Admin user
* Connecting to stripe API for credit card payments
* Handling selling of products online
* Setting of recurring payments using Stripe Plans
* Handling stripe refunds of payments
* Canceling subscriptions

Anywhere along the way I implemented any personal changes on
my end based on my own knowledge of web design/development.
Such as

* Using tailwindcss for generating the styles of the website

## Getting Started

### Dependencies

* [chi](https://github.com/go-chi/chi) v5.2.5 (for a routing framework)
* [stripe-go](https://github.com/stripe/stripe-go) v84.4.1 (go library for stripe API)
* [cors](https://github.com/go-chi/cors) v1.2.2 (http middleware for go)
* [tailwindcss](https://github.com/tailwindlabs/tailwindcss) v4.2.0 (as a css framework tool)
* [pgx](https://github.com/jackc/pgx/v5) v5.9.1 (Posgresql driver for go)
* [scs](https://github.com/alexedwards/scs/v2) v2.9.0 (HTTP session management for go)

### Installing manually

TODO: Add more information on installing and running the source code as the project progresses

### Authors

Max Kranker

## Version History

<details>
<summary>Click To Expand Releases</summary>

- 0.1.0
  - Change main header title
- 0.2.0
  - Setup API Handler CreateCustomerAndSubcribe
- 0.3.0
  - Add BronzePlanReciept Handler
- 0.4.0
  - Add /receipt/bronze Route
- 0.5.0
  - Add JS Session Storage and Redirect To /receipt/bronze
- 0.6.0
  - Add Initial receipt-plan Template Page
- 0.7.0
  - Add LoginPage Handler
- 0.8.0
  - Add /login Route
- 0.9.0
  - Add Initial Login Page Template
- 0.10.0
  - Add Login Link To Nav Of Base Layout
- 0.11.0
  - Add Test Javascript AJAX API Fetch For Authentication
- 0.12.0
  - Add readJSON API Helper Function
- 0.13.0
  - Add badRequest API Helper Function
- 0.13.1
  - Remove JS hidePayButton() and amountToCharge From Template
- 0.14.0
  - Add Initial CreateAuthToken API Handler
- 0.15.0
  - Add /api/authenticate Route
- 0.16.0
  - Add writeJSON API Helper Function
- 0.16.1
  - Fix badRequest To Set Header Name As "Content-Type"

</details>

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
