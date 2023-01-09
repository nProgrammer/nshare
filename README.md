# nShare API 
### Simple API for file sharing

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

