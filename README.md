# CircleCI Demo Monorepo
An example monorepo pipeline in CircleCI with Go services.

## What's happening in this project?
1. The `gatekeeper` workflow consists of a single `gatekeeper` job that runs by default on all commits and other triggers.
    - Inside the `gatekeeper`, we are scanning for which folders have changed and then making an API call using our [v2 API to trigger the pipeline](https://circleci.com/docs/api/v2/#trigger-a-new-pipeline).
    - When making the API trigger, we must specify parameters to set `gatekeeper` to `false` so that it doesn't retrigger the same workflow in a loop, and only trigger the desired service workflow(s).
2. The API call triggers another workflow, `service`, which will have jobs that run specifically for said service(s). The name of the service is passed as a parameter in the API call.

## Notes, Known Shortcomings, Desired Features
- This only works if each service in the monorepo follows the same steps / process (e.g. whatever is in `service` job). If different services require different tasks, then you'd need to define different jobs and potentially workflows for them.
- This also only works if the name of the service matches the name of the folder. We could decouple that, but the functionality is shown already - it would just be more config.
- We're using built-in pipeline parameters to compare current commit to previous commit. Ideal is to compare against last commit **built on CircleCI**, not just the last commit in the log (since multiple commits can be pushed, and only latest one is built). 
- Script in `gatekeeper` job does not cover the scenario where a branch does not have a previous commit. I believe that will break.
- Need to implement comparing tags

