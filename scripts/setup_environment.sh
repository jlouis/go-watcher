#!/bin/sh

## This is replaced by the CodeDeploy agent under install
ENV='%%%SYSTEM%%%'

aws --region eu-west-1 s3 cp s3://sgn-deployments/cicada/config.sh /opt/cicada/config.sh
chmod +x /opt/cicada/config.sh

