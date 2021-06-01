---
page_title: "playgroundtech_application resource - terraform-provider-playgroundtech"
subcategory: ""
description: |-
  
---

# playgroundtech_application (resource)

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

- **email** (String)  
  Phone number applicant want to be contacted on.
- **linkedin** (String)  
  Email applicant want to be contacted on.
- **phone_number** (String)  
  Linkedin of applicant.

### Optional

- **github** (String)
  Github of applicant to be used as further reference.
- **homepage** (String)  
  Homepage of applicant to be used as further reference.


