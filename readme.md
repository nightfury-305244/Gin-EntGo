# Project Title
Gin + EntGo

## Description
Sample Golang project that use entgo and gin framework for database connection and api integration.

## Installation
```bash
git clone https://github.com/yourusername/yourprojectname.git
cd yourprojectname
npm install
```

## About

```
1. "/test" GET API

It test the server is running on. 
It automatically set "example" as 12345, and test it. If it's running smoothly, it return 12345.
```

```
2. '/json' POST API

It get user name and password as a parameter and check it's in the database.
If the database have one, it return 200 code and input value.
```

```
3. '/user' POST API

This is the register funtion. You can register with user name, email and password.
This function set these data to database after validation.
```