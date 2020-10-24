## Pre-req

### Create User:
```bash
curl -X POST -d '{"username": "username","email": "email@email.com", "password": "password"}' localhost:8888/api/v1/users
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

### Delete User & Application:
To fetch your `:ID` and `Token` you can login with following:  

```bash
curl -d '{"username": "username","email": "email@email.com", "password": "password"}' localhost:8888/api/v1/login
```  

After receiving your `:id` and `$TOKEN` you run following command with replacement of your `:id` and `Token`.  
This will delete your user and any job application made by you.  
```bash
curl -X DELETE localhost:8888/api/v1/users/:id -H 'Authorization: Bearer {$TOKEN}' 
```

