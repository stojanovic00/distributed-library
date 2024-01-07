#!/bin/bash
kubectl apply -f persistent_volumes.yml
kubectl apply -f pv_claims.yml
kubectl apply -f config_maps.yml
kubectl apply -f deployments.yml
kubectl apply -f services.yml
kubectl apply -f ingress.yml


