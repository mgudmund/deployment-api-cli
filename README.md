# Deployment API Cli
![Go CI](https://github.com/mgudmund/deployment-api-cli/workflows/Go%20CI/badge.svg)

This is an MVP of a cli for our deployment metrics API. 

You need to have a product-spec.yaml in the directory where the command is executed. 

Sample can be found in this repo. 

## Startup Workflow

1. Sign-up using the sign-up command. You'll get a token that you will need to put in an environment variable called DEPLOYMENT_API_TOKEN
2. Run commands needed to track metrics

## Deployment workflow

1. Create new deployment using the create command. Save the deployment ID for next steps. 
2. Run your steps needed to deploy you application
3. Update the deployment using the ID you got in step 1, for each step in the deployment(created, running, done, cancelled, failed).
4. When deployment flow is done, the deployment ID can be discarded. 

  
