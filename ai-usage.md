## Example 1: Create Environment

### Agent Input
```json
{
  "tool": "create_environment",
  "input": {
    "name": "dev",
    "team": "platform"
  }
}

API Response
{
  "environmentId": "env-123",
  "status": "CREATED"
}


## Example 2: Deploy Application

### CLI Command
```platformctl deploy service --name payments-api

Structured Output
{
  "service": "payments-api",
  "status": "DEPLOYED"
}