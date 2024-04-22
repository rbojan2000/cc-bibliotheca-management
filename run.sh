#!/bin/bash

minikube start

kubectl apply -f kubernetes/central-library/central-library-pv.yaml
kubectl apply -f kubernetes/central-library/central-library-pvc.yaml
kubectl apply -f kubernetes/central-library/central-library-deployment.yaml
kubectl apply -f kubernetes/central-library/central-library-service.yaml

kubectl apply -f kubernetes/nis/nis-library-pv.yaml
kubectl apply -f kubernetes/nis/nis-library-pvc.yaml
kubectl apply -f kubernetes/nis/nis-library-deployment.yaml
kubectl apply -f kubernetes/nis/nis-library-service.yaml

kubectl apply -f kubernetes/belgrade/belgrade-library-pv.yaml
kubectl apply -f kubernetes/belgrade/belgrade-library-pvc.yaml
kubectl apply -f kubernetes/belgrade/belgrade-library-deployment.yaml
kubectl apply -f kubernetes/belgrade/belgrade-library-service.yaml

kubectl apply -f kubernetes/novisad/novisad-library-pv.yaml
kubectl apply -f kubernetes/novisad/novisad-library-pvc.yaml
kubectl apply -f kubernetes/novisad/novisad-library-deployment.yaml
kubectl apply -f kubernetes/novisad/novisad-library-service.yaml

kubectl apply -f kubernetes/ingress.yaml

echo "#### kubectl host #####"
kubectl get ingres
