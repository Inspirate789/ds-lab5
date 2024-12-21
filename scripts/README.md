# TODO

# Инициализация кластера (вне пайплайна)

### 1. Деплой Postgres, Kafka и Keycloak
```shell
./scripts/k8s-setup-storage.sh
```

### 2. Настраиваем Keycloak по инструкции https://www.keycloak.org/getting-started/getting-started-kube
- Логин: `user`
- Пароль: вывод команды `kubectl get secret keycloak -o jsonpath='{.data.admin-password}' | base64 --decode`

### 3. Настраиваем Keycloak по инструкции https://www.keycloak.org/getting-started/getting-started-kube
Создаём пользователя с кредами TODO
realm ds-labs

### 4. После создания Client берём его Client Secret и записываем в секрет проекта `TODO`

# Деплой приложения (в пайплайне)

### 5. Деплой сервисов приложения
```shell
./scripts/k8s-setup-app.sh
```

### 6. Прогон API тестов
```shell
newman run ./postman/collection.json -e postman/environment.json \
    --env-var serviceUrl=http://gateway.default.svc.cluster.local \
    --env-var identityProviderUrl=http://keycloak.default.svc.cluster.local/realms/ds-labs/protocol/openid-connect/token \
    --env-var username=aboba \
    --env-var password=aboba \
    --env-var clientId=aboba \
    --env-var clientSecret=tKEDd2gbem7UOaztsyclidAm6sGCNpCr \
    --bail
```
