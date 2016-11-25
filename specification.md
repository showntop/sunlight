# api specification
==========

apps share and perfect

##create user

```
method: post
url:    /users
request body:
    {
        email: "",
        mobile: "",
        password: ""
    }
    ps: email & mobile 
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```

## user login

```
method: post
url:    /sessions
request body:
    {
        field_value: username/mobile/email,
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```
