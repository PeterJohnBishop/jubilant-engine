# jubilant-engine

# notes

docker build -t peterjbishop/jubilant-engine:latest . 
<!-- build and tag image  -->

docker push peterjbishop/jubilant-engine 
<!-- should appear in Docker under Images -->

docker pull peterjbishop/jubilant-engine 

docker run -d --name=jubilant-engine-container -p 8080:8080 peterjbishop/jubilant-engine
<!-- run in detached mode, under the name jubilant-engine-container, expose port 8080 to 8080 on your machine -->

<!-- 
{ make some changes }

docker build -t peterjbishop/jubilant-engine:latest . 
docker push peterjbishop/jubilant-engine 
docker pull peterjbishop/jubilant-engine 
docker run -d --name=jubilant-engine-container -p 8080:8080 peterjbishop/jubilant-engine

{ make some changes }
-->

minikube start
<!-- start up minikube of course -->

kubectl apply -f deployment_go.yaml
<!-- Create the app or service described in the .yaml file -->

kubectl expose deployment jubilant-engine --type=NodePort --port=8080
<!-- create a service to forward external raffic to port 8080 where the app is listening -->

minikube service jubilant-engine
<!-- starts kubernetes pods and exposes the traffic -->
