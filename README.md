# jubilant-engine

# docker container and kubernetes | creating a server container and testing in kubernetes cluster

install homebrew
brew install minikube
install docker engine/docker dekstop

docker build -t peterjbishop/jubilant-engine:latest . 
<!-- build and tag image  -->

docker push peterjbishop/jubilant-engine 
<!-- should appear in Docker under Images -->

# installation

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

# docker compose | creating a server container with postgres container and testing in kubernetes cluster

docker-compose down

PSQL_USER={username} PSQL_PASSWORD={password} PSQL_DBNAME={dbname} docker-compose up --build
<!-- to run in Docker with env variables set -->

docker push peterjbishop/jubilant-engine:latest

kubectl create secret generic db-secret \
  --from-literal=PSQL_USER={username} \
  --from-literal=PSQL_PASSWORD={password} \
  --from-literal=PSQL_DBNAME={dbname} \
  -n {your-namespace}
<!-- to set env variables in kubernetes -->

minikube status

minikube start

kubectl apply -f deployment_postgres.yaml <!-- if changes >

kubectl apply -f deployment_go.yaml <!-- if changes >

kubectl expose deployment jubilant-engine --type=NodePort --port=8080

minikube service jubilant-engine

kubectl get pods