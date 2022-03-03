# Build automate app with docker + route53 + redis + nginx

[![FIBO Logo](https://fibo.cloud/assets/images/logo.svg)](https://fibo.cloud/)

Domain automation [Template]
## Суулгах зүйлс

1. [Аккоунт үүсгэх](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)

2. [Docker суулгах]
```
curl -l https://get.docker.com | bash

```

3. Шинэ folder үүсгэх

4. Clone хийх

```
git clone https://github.com/johandui/serverless-notify
```
5. .env file үүсгэх
AWS_ACCESS="REPLACE_AWS_ACCESS"
AWS_SECRET="REPLACE_AWS_SECRET"
AWS_HOST="REPLACE_AWS_ROUTE53_HOST"
AWS_DOMAIN="REPLACE_DOMAIN"
AWS_IP="REPLACE_PUBLIC_IP"
AWS_REDIS="REPLACE_REDIS_URL"

6. Build хийх
```
docker build -t domain-automation .
```

7. Deploy хийх
```
docker run -dt -v /var/run/docker.sock:/var/run/docker.sock .
```

----
