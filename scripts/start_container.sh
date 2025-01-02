#!/bin/bash

#Pull the docker image
docker pull bxtgeek/age-calc

#Run the docker container
docker run -d -p 8000:8080 bxtgeek/age-calc