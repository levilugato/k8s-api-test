Levi Lugato test

The application is running on : `http://acdf8e79-default-levilugat-96ef-1589219485.us-east-1.elb.amazonaws.com/` 

Settings explained :

    I am using EKS + Fargate on aws to hostage it

    I am using an ALB load balancer as ingress to expose it

    I am using metrics server + hpa to scale the pods by CPU us

    I am not saving the names on a database but I have been created the PVC as required


The API has 2 endpoints :

`/` use GET to test : curl http://acdf8e79-default-levilugat-96ef-1589219485.us-east-1.elb.amazonaws.com

`/hello` use POST with a path param of your choise : curl -X POST http://http://acdf8e79-default-levilugat-96ef-1589219485.us-east-1.elb.amazonaws.com/hello/yournamehere

Thank you

