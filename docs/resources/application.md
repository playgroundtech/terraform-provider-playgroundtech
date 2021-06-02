---
page_title: "playgroundtech_application resource - terraform-provider-playgroundtech"
subcategory: ""
description: |-
  
---

# playgroundtech_application

## Schema

```hcl
provider "playgroundtech" {
  email    = "email@email.com"
  password = "password"
}

resource "playgroundtech_application" "test" {
  phone_number = "0123456789"
  email        = "email@email.com"
  linkedin     = "linkedin.com"
  github       = "github.com"
  homepage     = "homepage.com"
}
```

### Required
- **phone_number** (String)   
  Phone number applicant want to be contacted on.
- **email** (String)  
  Email applicant want to be contacted on.
- **linkedin** (String)   
  Linkedin of applicant.

### Optional

- **github** (String)    
  Github of applicant to be used as further reference.
- **homepage** (String)    
  Homepage of applicant to be used as further reference.


