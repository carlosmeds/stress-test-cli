# Stress Test CLI

This is a simple CLI application to perform stress tests on a given URL. The application is built using Go and can be run inside a Docker container.

## Building the Docker Image
To build the Docker image, use the following command:

`
docker build -t stress-test:latest -f Dockerfile .
`

## Running the Application
To run the application inside a Docker container, use the following command:

`
docker run stress-test -u http://google.com -r 1000 -c 5
`

## Command Line Flags
The application supports the following command line flags:

```
--url or -u: The URL to stress test.
--requests or -r: The number of requests to perform (default is 100).
--concurrency or -c: The number of multiple requests to make at a time (default is 1).
```

## Expected Output
When you run the application locally, you should see an output similar to the following:

```
---------------REPORT---------------
[URL]:      http://google.com 
[REQUESTS]: 100
[DURATION]: 574.077371ms

100 requests returned status code 200
```