# reforged-lab

<hr>

###### 🌈 Table of contents

1. [Environment](#environment)
2. [Deployment](#deployment)
3. [Project details](#project-details)
4. [API details](#api-details)<br>
  4.1. [POST - `/api/v1/ads`](#post---apiv1ads-create-a-new-advertisement)

<hr>

# Environment

- GoLang 1.24.1
- MacOS 15.3, Ubuntu 24.04.1 LTS
- Docker Desktop (MacOS) or Docker Engine (Ubuntu)

# Deployment

- Clone this repository to your local machine:
    ```bash
    git clone https://github.com/cuongpiger/reforged-labs.git --depth 1 && cd reforged-labs
    ```

- Run the below command to deploy both PostgreSQL and API service:
    ```bash
    docker compose up -d
    ```

- Verify the deployment:
    ```bash
    docker ps
    ```
  You should see 2 containers running:
  ![docker-ps](./assets/01.png)

# Project details

- This application uses PostgreSQL as the database.
    - Using ORM to interact with the database.
    - The database used to store the Advertisement data.

- About the web-framework:
    - Using `gin-gonic` as the web-framework.
    - The API is RESTful.

- Following Clean Architecture principles.
    - The application is divided into 3 layers: `domain`, `usecase`, and `delivery`.
    - The `domain` layer contains the business logic.
    - The `usecase` layer contains the application logic.
    - The `delivery` layer contains the API logic.


# API details
## POST - `/api/v1/ads`: Create a new advertisement.

