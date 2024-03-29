deploy_k8s() {
    kubectl apply -f ./kubernetes/application/namespace.yaml
    kubectl apply -f ./kubernetes/application
    ./kubernetes/application/create-application-config.sh
    ./kubernetes/application/load-postgres-secret.sh
    ./kubernetes/application/load-user-cert.sh
    ./kubernetes/application/load-test-data-config.sh
    kubectl apply -f ./kubernetes/application/db
    kubectl apply -f ./kubernetes/application/test-data-service
    kubectl apply -f ./kubernetes/application/user-service
    kubectl apply -f ./kubernetes/application/book-service
    kubectl apply -f ./kubernetes/application/transaction-service
    kubectl apply -f ./kubernetes/application/web-service
}

delete_k8s() {
    kubectl delete -f ./kubernetes/application/web-service
    kubectl delete -f ./kubernetes/application/transaction-service
    kubectl delete -f ./kubernetes/application/book-service
    kubectl delete -f ./kubernetes/application/user-service
    kubectl delete -f ./kubernetes/application/test-data-service
    kubectl delete -f ./kubernetes/application/db
    kubectl delete -f ./kubernetes/application
}

if [ "$1" = "delete" ]; then
    delete_k8s
elif [ "$1" = "deploy" ]; then
    deploy_k8s
else
    echo "Usage: $0 [deploy|delete]"
    exit 1
fi
