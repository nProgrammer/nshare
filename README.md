# nShare API 
### Simple API for file sharing

## Preparation before first start-up

1. Remember to create directory called "data" or something like that - this directory will store every asset uploaded by users (alternatively you can use docker volume).
2. Remember to change DB password - in the realeses tabs you'll find DB repo with docker-compose and db-builder.sql script.

``API: docker-compose.yaml`` :
```yaml
version: '3.7'

services:
  nshare-api:
    container_name: api-nshare
    build: ./
    image: api-nshare
    restart: always
    network_mode: "host"
    volumes:
      - <PATH_TO_YOUR_DIRECTORY_OR_INIT_VOLUME_AND_IMPLEMENT_HERE>:/var/nshare-data/
    environment:
      - WORKINGPORT=8080
      - SSLMODE=false
      - DB_SERVERNAME=localhost
      - DB_SERVER_PORT=3309
      - DB_NAME=nShare
      - DB_PASSWD=example #! Remember to change password
      - DB_USER=root
      - FULL_ERRORS=true
```

```DB: docker-compose.yaml```
```yaml
version: '3.7'

services:
  db:
    container_name: mysql-nshare
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example #! Remember to change password
    ports:
      - "3309:3306"
    volumes:
      - ./db-vol:/var/lib/mysql
```

## Routes
1. **Check if API is running correctly**
    
    **method:** GET

    **url:** ```http://<SERVER>:<PORT>/status```

    **expected answer:**
    ```json
    "Works"
    ```

    ***

2. **Initializing user's session**
    
    **method:** POST

    **url:** ```http://<SERVER>:<PORT>/users/initSession```

    **body:** 
    ```json
    {
        "login":"default",
        "pass":"tluafed"
    }
    ```
    
    **response in case of correct authorization:**
    ```json
    {
        "id": 0,
        "userID": 1,
        "skey":"962aebcf892aa978805e690b3616b7119ca3bfd9a584451603eec7a9e7989a68"
    }
    ```

    **response in case of incorrect authorization:**
    ```json
    "Bad password or login"
    ```
    ***

