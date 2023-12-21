#!/bin/bash
docker build -t bankapp:v1 .
docker tag bankapp:v1 moazrefat/bankapp:v1
docker push moazrefat/bankapp:v1