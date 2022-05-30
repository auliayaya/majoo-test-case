# Majoo Test Backend

Hai everyone, this is repo for completing Majoo Test Case as Backend Engineer Interview.
Please take a look to my blog post for the details [link](https://auliaillahi.my.id/test-case-backend-majoo/)
For postman doc please download here [link](https://www.getpostman.com/collections/af2e0027b22680c51579)

## Notes

Here i explain a little bit about what i've done.
I create 3 Endpoint:

- Login
  - `/login`
- Merchant Omzet (Calculate daily revenue on november with pagination and date params with authenticated user)
  - `/merchant/omzet/:id?start_date=2021-11-01&end_date=2021-11-30&page_size=10&page_number=1`
- Outlet Omzet (Calculate daily revenue on november with pagination and date params with authenticated user)
  - `/outlet/omzet/:id?start_date=2021-11-01&end_date=2021-11-30&page_size=10&page_number=1`

you can test it on my live API using postman configuration doc that i put on top of this description.
here is my live API [https://majoo-api.auliaillahi.my.id/api/v1/checkhealth](https://majoo-api.auliaillahi.my.id/api/v1/checkhealth)

But if you wanna try this repo on your local machine, you can follow the instruction in subcontent getting started

## Getting Started

- Clone this repository
- Rename app.example to app.env
- Put your database configuration there
- Then run `make server` on your command line terminal
