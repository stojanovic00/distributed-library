#!/bin/bash
kubectl delete -f ingress.yml
kubectl delete -f services.yml
kubectl delete -f deployments.yml
kubectl delete -f config_maps.yml
kubectl delete -f pv_claims.yml
kubectl delete -f persistent_volumes.yml

