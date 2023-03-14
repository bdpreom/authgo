Token controller = one endpoint will be used to gernetate jwt,user has to send a list of valid email amd password
User Controller = regiser a user new user will be created
Secured Controller = dummy controller secured by jwt token,showcase the ability of the middleware that we will build to restrict access to only the requests that have an actual valid JWT in the request header.

go get -u github.com/gin-gonic/gin

GORM :

As you can see, our User Model will have a Name, Username, Email, and password. Here, the Username and Email will be unique. This means, that once we complete our application and try to register new users with the same username or email, the code won’t allow you to do it. The best part is that you don’t have to write any code specifically for this. Everything is handled by GORM

he gorm.Model specification adds some default properties to the Model,
like id,
created date,
modified date, and
deleted date.

go get gorm.io/gorm
go get gorm.io/driver/mysql

go get golang.org/x/crypto/bcrypt

change the config file for db entry
