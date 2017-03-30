#!/usr/bin/env bash
set -e

echo "In configure-ecs.sh"
script_dir="$(dirname "$0")"
bin_dir="$(dirname $0)/../bin"

echo The value of arg 0 = $0
echo The value of arg 1 = $1
echo The value of arg 2 = $2 
echo The value of arg 3 = $3 
echo The value of arg 4 = $4  
echo The value of arg 5 = $5
echo The value of arg script_dir = $script_dir
echo UPDATED 201606072316

MY_STACK=$1
MY_DOCKERHUB_REPO=$2
DOCKER_USER_NAME=$3
DOCKER_PASS=$4
MY_URL=$5

echo The value of MY_STACK is $MY_STACK 
echo The value of MY_DOCKERHUB_REPO is $MY_DOCKERHUB_REPO

# Unique ID for Docker tag
uuid=$(date +%s)
dockerhub_repo="$MY_DOCKERHUB_REPO"
ecs_stack_name="$MY_STACK"

ecs_template_url="https://s3.amazonaws.com/ramitsurana-cloudformation/ecs-pipeline.json"
#ecs_template_url="$MY_URL"

echo The value of arg uuid = $uuid
eval $(docker login)

# Build, Tag and Deploy Docker
docker build -t $dockerhub_repo:$uuid .
docker login --username=$DOCKER_USER_NAME --password=$DOCKER_PASS
docker push $dockerhub_repo

aws cloudformation update-stack --stack-name $ecs_stack_name --template-url $ecs_template_url --region us-east-1 --capabilities="CAPABILITY_IAM" --parameters ParameterKey=AppName,UsePreviousValue=true ParameterKey=ECSRepoName,UsePreviousValue=true ParameterKey=DesiredCapacity,UsePreviousValue=true ParameterKey=KeyName,UsePreviousValue=true ParameterKey=RepositoryBranch,UsePreviousValue=true ParameterKey=RepositoryName,UsePreviousValue=true ParameterKey=InstanceType,UsePreviousValue=true ParameterKey=MaxSize,UsePreviousValue=true ParameterKey=S3ArtifactBucket,UsePreviousValue=true ParameterKey=S3ArtifactObject,UsePreviousValue=true ParameterKey=SSHLocation,UsePreviousValue=true ParameterKey=YourIP,UsePreviousValue=true ParameterKey=ImageTag,ParameterValue=$uuid ParameterKey=ECSCFNURL,ParameterValue=$ecs_template_url ParameterKey=GitHubOwner,UsePreviousValue=true ParameterKey=GitHubToken,UsePreviousValue=true

sleep 10
