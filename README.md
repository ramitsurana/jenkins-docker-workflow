Inspired by [Stelligent ECS CloudFormation Solution](https://stelligent.com/2016/05/26/automating-ecs-provisioning-in-cloudformation-part-1/) and [Stelligent Jenkins ECS Solution](https://stelligent.com/2016/08/24/containerized-ci-solutions-in-aws-part-1-jenkins-in-ecs/)

YAML Reference: https://github.com/microservices-today/ngp-ecs/tree/master

# Jenkins-docker-workflow

This repo is a sample POC for Jenkins-Docker-ECS Workflow.

## Introduction

We are going to build up a ECS based solution using the AWS CloudFormation Service.

## Prerequisites

* AWS Account
* AWS CLI Setup

## Setup
 
For getting started, the CFN template will setup nodes for an ECS cluster using an auto-scaling group.Some of the parameters used in the json file are:

* InstanceType:

This is used to define the type of EC2 insttance that you are going to use while using an ECS cluster.

* ECS Cluster:

This is the ECS Cluster parameters that are important to define as it determines the size of ECS Cluster you are building.

* Bootstrap:

This parameter is used to bootstrap the application you want to deploy along with the several constraints that you wish to define.

## License

MIT License
