# **Project NANDEMO**

This is a demo project showcasing:

- OAuth2 login (Google / Microsoft)
- Event-driven microservice architecture
- Golang backend microservices
- React frontend

Tech stack:
- Go
- Docker
- React + TypeScript
- GCP Pub/Sub

Initial services:
- Gateway Service
- Order Service
- Merchant Service


## **About the Makefile**

You will be able to spin up the dockerised services, including gateway, order and merchant services by running:

*make up* in your terminal.

To run integration tests, use *make test*, but **only after** you have successfully spun up all the services.

You may use **curl http://localhost:8080/health** to check if the services are up.

If you like, add more integration tests in the appropriate folder and register them to get them run inside the **docker-compose.yml** file.

After using or testing, you may use *make down* in the terminal to stop the services.

Using *docker system prune -f* command can clean up unused Docker data. It removes stopped containers, unused networks, dangling images and build cahce.

## Logging

Use **docker logs {container name}** to get logs.

This project uses Go 1.24.

## More

There will be a private project, **Nandemo-ii**, with more features coming up in the future.
