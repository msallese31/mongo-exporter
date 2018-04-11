# mongo-exporter
super unnecessary microservice to listen for requests to export mongo data

# Build/Publish the Docker Image

```
./buildTagPush.sh
```

# Run the image locally 

```
docker run mongo-exporter
```

# Make a request

```
curl localhost:8080
```