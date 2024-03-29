![Lib Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/lib-test.yml/badge.svg)
![User Service Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/user-service-test.yml/badge.svg)
![Book Service Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/book-service-test.yml/badge.svg)
![Transaction Service Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/transaction-service-test.yml/badge.svg)
![Web Service Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/web-service-test.yml/badge.svg)
![Test Data Service Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/test-data-service-test.yml/badge.svg)
![Reverse Proxy Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/reverse-proxy-test.yml/badge.svg)
![Load Balancer Badge](https://github.com/akatranlp/hsfl-master-ai-cloud-engineering/actions/workflows/load-balancer-test.yml/badge.svg)

[![codecov](https://codecov.io/gh/akatranlp/hsfl-master-ai-cloud-engineering/graph/badge.svg?token=UMTYYPZ8TM)](https://codecov.io/gh/akatranlp/hsfl-master-ai-cloud-engineering)

# VerseVault: Unleash Your Words, Earn Your Worth!

Welcome to VerseVault, your one-stop destination for unleashing your creative potential and turning your passion for writing into a profitable venture. VerseVault is a dynamic platform where aspiring writers and wordsmiths of all levels can craft, publish, and monetize their written works.

Here's how it works:

Write & Publish: Start by sharing your thoughts, stories, poems, or expertise with our global community. Whether you're a seasoned writer or just getting started, VerseVault provides you with the tools to create captivating content.

Monetize Your Writing: Your words have value, and at VerseVault, we believe you should be rewarded for your talent. List your texts for sale, and other users can purchase them using VV-Coins, our virtual currency.

Earn Real Money: Once you've accumulated VV-Coins, you have the option to convert them into real money. It's your hard work, and we ensure you reap the benefits.

Explore & Discover: Readers can explore an array of engaging content on VerseVault, making it a hub for discovering new voices and fresh perspectives. From articles and essays to fiction and non-fiction, there's something for everyone.

VerseVault empowers writers to not only share their stories but also earn a living doing what they love. Join us today, and let your creativity flourish while turning your passion for writing into a rewarding career. Start your journey on VerseVault now and monetize your words like never before!

## How to deploy our application

Create own `.env` file from `.env-example`. This includes every configuration for local or kubernetes deployments.
Especially `AUTH_IS_ACTIVE` can be changed from true to false or vise versa, to enable or disable the need of an accessToken when talking to the service endpoints.

Create own `./src/reverse-proxy/config.yaml` from `./src/reverse-proxy/config-example.yaml`. You don't need to edit this file in the normal case.

The docker-compose-files include this automatically and in the scripts under `./kubernetes/application` the `.env` is loaded and the values from it are used to generate secrets and configMaps.

You need to create an RSA-KeyPair, which is used to sign and validate the JWT-Tokens our user-service produces.

```bash
mkdir -p ./src/user-service/certs
cd ./src/user-service/certs
openssl genrsa -out access-key.pem 2048
openssl rsa -in access-key.pem -outform PEM -pubout -out access-public.pem
openssl genrsa -out refresh-key.pem 2048
openssl rsa -in refresh-key.pem -outform PEM -pubout -out refresh-public.pem
```

### Deploy locally for dev or testing

Always execute the following command to start backend and frontend in docker

```bash
# with gRPC Communication between services:
docker compose up -d --build

# without
docker compose -f docker-compose-no-grpc.yaml up -d --build

# with our own loadbalancer and orchestrator
docker compose -f docker-compose-loadbalance.yaml up -d --build
```

If you want to develop on the frontend execute this commands to get hot-reload on it.

```bash
cd ./src/web-service
pnpm install
pnpm dev
```

### Deploy production version on kubernetes

You first need to complete the [steps up above](#how-to-deploy-our-application) and then execute this bunch of commands in the specified order to deploy the whole application. These comands are also in another script under
`./scripts/k8s-deploy.sh`. Use it with the argument `deploy` to deploy the application to your current k8s-cluster.

```bash
kubectl apply -f ./kubernetes/application/namespace.yaml
kubectl apply -f ./kubernetes/application
./kubernetes/application/create-application-config.sh
./kubernetes/application/load-postgres-secret.sh
./kubernetes/application/load-user-cert.sh
./kubernetes/application/load-test-data-config.sh
kubectl apply -f ./kubernetes/application/db
kubectl apply -f ./kubernetes/application/test-data-service
kubectl apply -f ./kubernetes/application/user-service
kubectl apply -f ./kubernetes/application/book-service
kubectl apply -f ./kubernetes/application/transaction-service
kubectl apply -f ./kubernetes/application/web-service
```

## How to deploy monitoring software on the kubernetes cluster

As an addition to our application, we also have manifests to monitor your cluster with prometheus and grafana.
Deploy them with the following commands.

```bash
kubectl apply -f ./kubernetes/monitoring
kubectl apply -f ./kubernetes/monitoring/prometheus
kubectl apply -f ./kubernetes/monitoring/grafana
kubectl apply -f ./kubernetes/monitoring/kube-state-metrics
```

Our grafana dashboard is located at `./kubernetes/monitoring/grafana/dashboard.json`, which includes graphs to check the current memory-usage and cpu-utility aswell as the incoming and outgoing network traffic.

## Authors

Fabian Petersen\
fabian.petersen@stud.hs-flensburg.de\
Hochschule Flensburg

Pascal Friedrichsen\
pascal.friedrichsen@stud.hs-flensburg.de\
Hochschule Flensburg

Dominik Heckner\
dominik.heckner@stud.hs-flensburg.de\
Hochschule Flensburg
