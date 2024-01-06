#!/bin/bash
kubectl create configmap central-config --from-env-file=.env.central
kubectl create configmap central-db-config --from-env-file=.env.postgre-central-db


kubectl create configmap local-ns-config --from-env-file=.env.local-ns
kubectl create configmap local-bg-config --from-env-file=.env.local-bg
kubectl create configmap local-ni-config --from-env-file=.env.local-ni

kubectl create configmap local-db-ns-config --from-env-file=.env.postgre-local-db-ns
kubectl create configmap local-db-bg-config --from-env-file=.env.postgre-local-db-bg
kubectl create configmap local-db-ni-config --from-env-file=.env.postgre-local-db-ni




