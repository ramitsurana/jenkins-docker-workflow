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

* Cloudformation JSON Template

## Installation 

Simply follow the below commands:
````
$ git clone https://github.com/ramitsurana/jenkins-docker-workflow/
$ cd jenkins-docker-workflow/cloudformation-test/
$ chmod +x cf-test
$ mv cf-test /usr/local/bin/
$ cf-test
Checking Prerequisites ..
JQ installed on linux

 Enter path of your file: ecs-pipelline.json
Checking Parameters
Parameters OK
Checking Resources
Resources OK
Checking Mappings
Mappings OK
Checking Metadata
Metadata OK
Do you wish to check your services (y/n) ?n
exit status 1
````

## License

MIT License
