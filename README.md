# jubilant-engine

# installation

docker build -t peterjbishop/jubilant-engine:latest . 
<!-- build and tag image  -->

docker push peterjbishop/jubilant-engine 
<!-- should appear in Docker under Images -->

docker run -d --name=jubilant-engine-container -p 8080:8080 peterjbishop/jubilant-engine
<!-- run in detached mode, under the name jubilant-engine-container, expose port 8080 to 8080 on your machine -->

