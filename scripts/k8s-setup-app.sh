helm repo add gruntwork https://helmcharts.gruntwork.io
helm repo update
helm upgrade cars-api gruntwork/k8s-service -f deployments/api-services/cars.yaml --install --wait --atomic
helm upgrade rental-api gruntwork/k8s-service -f deployments/api-services/rental.yaml --install --wait --atomic
helm upgrade payment-api gruntwork/k8s-service -f deployments/api-services/payment.yaml --install --wait --atomic
helm upgrade gateway gruntwork/k8s-service -f deployments/api-services/gateway.yaml --install --wait --atomic
helm upgrade retryer gruntwork/k8s-service -f deployments/api-services/retryer.yaml --install --wait --atomic
