#!/usr/bin/env bash

minikube delete && minikube start && eval $(minikube docker-env) && helm init && make
