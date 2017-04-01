# Testing Cloudformation Templates

## Problem

Testing and verifying cloudformation templates can be troublesome.

## Solutions

Basic cloudformation can be divided into 4 categories such as:

* Parameters
* Mappings 
* Resources 
* Group

## Prerequisites

* Golang
* Jq

## How to use the test ?

````
$ go run main.go 

 Enter path of your file: ecs-pipeline.json
Checking Parameters
Parameters OK
Checking Resources
Resources OK
Checking Mappings
Mappings OK
Checking Metadata
Metadata OK
Basic Checks completed.
OPTIONS [.Metadata][.Resources] 
List of services: 
 * .EcsCluster 
 * .EcsElb 
 * .ECSAutoScalingGroup 
 * .ContainerInstances
 ````

## License

MIT License
