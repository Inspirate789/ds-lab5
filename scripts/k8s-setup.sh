helm upgrade postgres oci://registry-1.docker.io/bitnamicharts/postgresql -f deployments/postgres/values.yaml --install --wait --atomic
helm upgrade kafka oci://registry-1.docker.io/bitnamicharts/kafka -f deployments/kafka/values.yaml --install --wait --atomic
helm upgrade keycloak oci://registry-1.docker.io/bitnamicharts/keycloak -f deployments/keycloak/values.yaml --install --wait --atomic
helm repo add gruntwork https://helmcharts.gruntwork.io
helm repo update
helm upgrade cars-api gruntwork/k8s-service -f deployments/api-services/cars.yaml --install --wait --atomic
helm upgrade rental-api gruntwork/k8s-service -f deployments/api-services/rental.yaml --install --wait --atomic
helm upgrade payment-api gruntwork/k8s-service -f deployments/api-services/payment.yaml --install --wait --atomic
helm upgrade gateway gruntwork/k8s-service -f deployments/api-services/gateway.yaml --install --wait --atomic
helm upgrade retryer gruntwork/k8s-service -f deployments/api-services/retryer.yaml --install --wait --atomic
