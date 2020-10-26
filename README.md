## Pre-req

First, build and install the provider.

```shell
$ make install
```
In order to make request to the api you need to create a user.  
Use the following command and to your information:
```bash
curl -X POST -d '{"username": "username","email": "email@email.com", "password": "password"}' 13.48.134.210/api/v1/users
```

### Provider Docs:
Required Providers:
```hcl
terraform {
  required_providers {
    playgroundtech = {
      versions = ["0.2"]
      source   = "playgroundtech.se/edu/playgroundtech"
    }
  }
}
```

In order to make calls against the Playgroundtech provider you need to authenticate.    
This can be done either by setting the following environment variables `PLAYGROUNDTECH_EMAIL` and `PLAYGROUNDTECH_PASSWORD`.  
It's also possible to set them by defining the following block in your `.tf` file.  
```hcl
provider "playgroundtech" {
  email    = "email@email.com"
  password = "password"
}
```
### Playground Resource:  
Manages a job application to Playgroundtech.  

#### Example Usage:  
```hcl
resource "playgroundtech_application" "test" {
  phone_number = "0123456789"
  email        = "email@email.com"
  linkedin     = "linkedin.com"
  github       = "github.com"
  homepage     = "homepage.com"
}
```

#### Argument Reference:  
- `phone_number` | (Required) - String    
Phone number applicant want to be contacted on.  

- `email` | (Required) - String    
Email applicant want to be contacted on.  
  
- `linkedin` | (Required) - String  
Linkedin of applicant.  

- `github` | (Optional) - String  
Github of applicant to be used as further reference.  

- `homepage` | (Optional) - String  
Homepage of applicant to be used as further reference.  
  
  
### Delete User & Application:
To fetch your `:ID` and `Token` you can login with following:  

```bash
curl -d '{"username": "username","email": "email@email.com", "password": "password"}' 13.48.134.210/api/v1/login
```  

After receiving your `:id` and `$TOKEN` you run following command with replacement of your `:id` and `Token`.  
This will delete your user and any job application made by you.  
```bash
curl -X DELETE 13.48.134.210/api/v1/users/:id -H 'Authorization: Bearer {$TOKEN}' 
```

