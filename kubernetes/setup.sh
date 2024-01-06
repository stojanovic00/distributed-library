#!/bin/bash
kubectl apply -f pv_claims.yml
kubectl apply -f persistent_volumes.yml
./make_config_map.sh
kubectl apply -f deployments.yml

