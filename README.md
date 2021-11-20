## System requirements 
You need to have [Docker, And Docker-Compose](https://www.docker.com) 
installed in order to build and run the application.
No additional tools required.

## To install docker-compose Run the Following

1. Downlaod using this command
    ```shell script
     sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    ```
2. Grant The binery execute priv
    ```shell script
     sudo chmod +x /usr/local/bin/docker-compose
    ```

## How to build and run
1. Build the images
    ```shell script
    docker-compose build
    ```
2. Spin up the continers:
    ```shell script
    docker-compose up
    ```
Access the application via http://localhost:8080
