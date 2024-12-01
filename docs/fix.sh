#!/bin/sh

sed -i "s/x-nullable/nullable/g" ./docs/docs.go
sed -i "s/x-omitempty/omitempty/g" ./docs/docs.go
sed -i "s/x-example/example/g" ./docs/docs.go
sed -i "s/x-nullable/nullable/g" ./docs/swagger.json
sed -i "s/x-omitempty/omitempty/g" ./docs/swagger.json
sed -i "s/x-example/example/g" ./docs/swagger.json
sed -i "s/x-nullable/nullable/g" ./docs/swagger.yaml
sed -i "s/x-omitempty/omitempty/g" ./docs/swagger.yaml
sed -i "s/x-example/example/g" ./docs/swagger.yaml