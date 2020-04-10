docker build -t navalta3030/go-jwt -f ./../Dockerfile ./../
docker push navalta3030/go-jwt:latest

kubectl apply -f .