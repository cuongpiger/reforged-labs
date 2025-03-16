# reforged-lab

<hr>

###### 🌈 Table of contents

1. [Environment](#environment)
2. [Deployment](#deployment)

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