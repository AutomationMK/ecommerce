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
- 0.17.0
  - Change CreateAuthToken To Use writeJSON Helper Function
- 0.18.0
  - Add GetUserByEmail DBModel Method
- 0.19.0
  - Add invalidCredentials API Helper Function
- 0.20.0
  - Add passwordMatches API Helper Function
- 0.20.1
  - Fix badRequest To Set Header To Status 400
- 0.21.0
  - Add User Database Query and Password Validation To CreateAuthToken
- 0.22.0
  - Add bcrypt Golang Module
- 0.23.0
  - Add Token Type and GenerateToken Function For Authentication
- 0.23.1
  - Fix 0.23.0 Version In README.md File
- 0.24.0
  - Add Token Generation and Token To JSON Response In CreateAuthToken
- 0.25.0
  - Add Database Migrations To Create Tokens Table
- 0.26.0
  - Add InsertToken DBModel Method
- 0.27.0
  - Add InsertToken DBModel Method To CreateAuthToken
- 0.28.0
  - Update InsertToken To Delete Existing Tokens
- 0.29.0
  - Add showSuccess and showError JS Functions and Add Data To Storage
- 0.30.0
  - Add Login JS Logic To Base Layout
- 0.31.0
  - Redirect To Home Page After Successfull Login
- 0.32.0
  - Add Initial CheckAuthentication API Handler
- 0.33.0
  - Add /api/is-authenticated API Route
- 0.34.0
  - Add Initial checkAuth JS Function To Base Layout
- 0.35.0
  - Add checkAuth JS Function To Terminal Template Page
- 0.36.0
  - Add authenticateToken Helper Function and Update CheckAuthentication
- 0.37.0
  - Add GetUserForToken DBModel Method
- 0.38.0
  - Update authenticateToken To Use GetUserForToken
- 0.39.0
  - Migrate To Add Expiry Column To Tokens Table
- 0.40.0
  - Update InsertToken and GetUserForToken To Use Expiry
- 0.41.0
  - Add Auth API Middleware Function
- 0.42.0
  - Add VirtualTerminalPaymentSucceeded API Handler
- 0.43.0
  - Remove stripe-js Partial From VirtualTerminal Handler
- 0.44.0
  - Add /api/admin Routes
- 0.45.0
  - Replace stripe-js Partial With Custom Script and Add Receipt Section
- 0.45.1
  - Fix VirtualTerminalPaymentSucceeded To Have Payment Intent and Method
- 0.46.0
  - Add Auth Middleware For Frontend 
- 0.47.0
  - Add /admin Routes Using Auth Middleware
- 0.47.1
  - Fix Charge Another Card A Tag To Use /admin/virtual-terminal
- 0.47.2
  - Fix Virtual Terminal Base Nav A Tag To Use /admin/virtual-terminal
- 0.48.0
  - Add Authenticate DBModel Method
- 0.49.0
  - Add PostLoginPage Frontend Handler
- 0.50.0
  - Add /login Post Route
- 0.51.0
  - Add JS login_form Submit For Authentication On Frontend
- 0.52.0
  - Add Logout Frontend Handler
- 0.53.0
  - Add /logout Route
- 0.54.0
  - Redirect To /logout In Base Layout JS Logout Function
- 0.55.0
  - Add sessions Table For SCS Session Storage
- 0.56.0
  - Add github.com/alexedwards/scs/pgxstore Module
- 0.57.0
  - Change DBModel DB Object To Be *pgxpool.Pool Type
- 0.57.1
  - Fix conn.Close To Not Have Context As Argument In Backend
- 0.58.0
  - Add Database Pool Connection To session.Store
- 0.58.1
  - Fix Login and Subscription Templates To Use Flex For Section
- 0.59.0
  - Add ForgotPassword Frontend Handler
- 0.60.0
  - Add /forgot-password GET Route
- 0.61.0
  - Add forgot-password Template Page
- 0.62.0
  - Add /forgot-password Link At Bottom Of Login Form

</details>

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
