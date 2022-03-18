# bookish-couscous
OAuth2/OpenID AuthZ/AuthN Server

Bookish-Couscous serves as the upstream authorization and authentication provider for our services. 

The authentication server will use the following providers:
- Google 
- Facebook
- Microsoft
- Apple

### Dependencies
- [OAuth2](https://github.com/golang/oauth2)
- [Go-Fiber](https://github.com/gofiber/fiber)
  - A blazing fast HTTP engine
- [GoDotEnv](https://github.com/joho/godotenv)
  - Loads environment variables from .env files
- [Zap](https://github.com/uber-go/zap)
  - Blazing fast, structured, leveled logging in Go.
- [Yaml](https://github.com/go-yaml/yaml)
  - Used to decode/encode YAML files
- [Viper](https://github.com/spf13/viper)
  - Viper is a complete configuration solution for Go applications including 12-Factor apps.