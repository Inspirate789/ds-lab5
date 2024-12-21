helm upgrade postgres oci://registry-1.docker.io/bitnamicharts/postgresql -f deployments/postgres/values.yaml --install --wait --atomic
helm upgrade kafka oci://registry-1.docker.io/bitnamicharts/kafka -f deployments/kafka/values.yaml --install --wait --atomic
helm upgrade keycloak oci://registry-1.docker.io/bitnamicharts/keycloak -f deployments/keycloak/values.yaml --install --wait --atomic
