#!/bin/bash
docker build -t bankapp:db .
docker tag bankapp:db moazrefat/bankapp:db
docker push moazrefat/bankapp:db