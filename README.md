# Tag Along App

Welcome to Tag Along!! A web application foced on build a great community. Our app supports features which enhances:
* Community Building
* Event Planning
* Social Interactions


## Instructions to run the App

Before following the below instructions, make sure that you have [Go](https://go.dev/) installed. Once you have Go configured in your system, follow the below commands to run the application.

1. Download the code.
2. Navigate to the folder containing the `go.mod` file.
3. Download the dependencies by running the below command:
    ```bash
    go get github.com/prasenjeetpaul/tag_along_app
    ```

    If the above command fails, try the below two commands to install the dependencies manually:
    ```bash
    go get github.com/google/uuid
    go get github.com/gin-gonic/gin
    ```
4. Once all the dependencies are downloaded, run the application using the below command from the same folder:
    ```bash
    go run .
    ```
5. Application API would start on `localhost:8080`. You can access the API directly by a web browser or by tools like Postman, etc
