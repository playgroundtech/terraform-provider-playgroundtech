---
page_title: "PlaygroundTech Provider"
subcategory: ""
description: |-
  
---

# PlaygroundTech Provider
This provider can be used to send job applications to [PlaygroundTech](https://playgroundtech.io).

### Create User:
In order to make request to the api you need to create a user.  
Use the following command and add your information:
```bash
curl -X POST -d '{"email": "email@email.com", "password": "password"}' https://api.playgroundtech.io/v1/users
```

### Delete User & Eventual Application:
To fetch your `:ID` and `Token` you can login with following:

```bash
curl -d '{"email": "email@email.com", "password": "password"}' https://api.playgroundtech.io/v1/login
```  

After receiving your `:id` and `$TOKEN` you run following command with replacement of your `:id` and `Token`.  
This will delete your user and any job application made by you.
```bash
curl -X DELETE https://api.playgroundtech.io/v1/users/:id -H 'Authorization: Bearer {$TOKEN}' 
```

## Schema
```hcl
provider "playgroundtech" {
  email    = "email@email.com"
  password = "password"
}
```
### Optional

- **email** (String)  
  Can be set by declaring it in the provider block. Or by exporting it as an environment variable.  
  `export PLAYGROUNDTECH_EMAIL=your@email.com`
- **password** (String, Sensitive)  
  Can be set by declaring it in the provider block. Or by exporting it as an environment variable.  
  `export PLAYGROUNDTECH_PASSWORD=your_secret_password`
