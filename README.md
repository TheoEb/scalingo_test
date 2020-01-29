# scalingo_test
Technical test for Scalingo in Go.

# Installation
```git clone https://github.com/TheoEb/scalingo_test```


# How to run
### server
Get a github API key and put it in a file named "github_api.key" at the root of the repo.

```
cd backend
GITHUB_KEY_FILE=<path_to_the_key_file> go run main.go
```

### web client

```
cd frontend
npm install --silent
npm install react-scripts@3.0.1 -g --silent
PORT=8766 npm start
```

Open a browser on localhost:8766

### With docker

```docker-compose up -d```

Open a browser on localhost:8766

### Screenshot
![Screenshot](assets/screenshot.png)