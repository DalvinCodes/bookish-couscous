# bookish-couscous
OAuth2/OpenID AuthZ/AuthN Server

Bookish-Couscous serves as the upstream authorization and authentication provider for our services. 

The authentication server will use the following providers:
- Google 
- Facebook
- Microsoft
- Apple

Frameworks
- [OAuth2](https://github.com/golang/oauth2)
- [Go-Fiber](https://github.com/gofiber/fiber)
  - A blazing fast HTTP engine
- [GoDotEnv](https://github.com/joho/godotenv)
  - Loads environment variables from .env files