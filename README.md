# Review chatbot

## Description

This project has the goal of defining a way to control workflow executions based on chat interactions.

### Prerequisites

- Go 1.21 or later
- Docker

## Installation

 1. Clone this repository:
 ```bash
    git clone https://github.com/brendontj/review-chatbot.git
 ```

 2. Install the dependencies:

 ```bash
    go mod download
 ```

 3. Run docker container for the DB:

 ```bash
    docker-compose up -d
 ```

## Usage

1. Run the project:

```bash
    go run main.go
```

2. Access the chat page in your browser:

```
    localhost:8000
```

3. To initiate a pre-defined workflow you need to send a signal to the running application:

Write in the chat `review-signal` , and then the review workflow will start soon


## Further developments 

- Make the persistence transactional
- Adjust the use case for more than one workflow execution
- Adjust the use case to run different workflow types
- Add tests
- Remove hard-coded db credentials 