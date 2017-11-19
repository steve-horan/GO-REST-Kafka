# REST API used to stop and start kafka containers for edge case testing.

If you choose to build this in the docker container, first compile the code:

go build

If on a Mac:

env GOOS=linux GOARCH=amd64 GOARM=7 go build

If Windows:

LOL

docker run -p 4444:4444 -v /var/run/docker.sock:/var/run/docker.sock -it go-rest-kafka