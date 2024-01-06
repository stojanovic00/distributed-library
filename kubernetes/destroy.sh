#!/bin/bash
kubectl delete -f pv_claims.yml
kubectl delete -f persistent_volumes.yml