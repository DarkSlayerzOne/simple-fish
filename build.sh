docker build -t darkxlayers/simple-fish:latest -t darkxlayers/simple-fish:$SHA .
docker push darkxlayers/simple-fish:latest
docker push darkxlayers/simple-fish:$SHA
kubectl apply -f kubs
kubectl set image deployments/go-fish-deployment server=darkxlayers/simple-fish:$SHA