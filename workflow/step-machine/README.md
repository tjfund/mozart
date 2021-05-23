# Step Machine

`machine` is an implementation of the AWS State Machine specification. The primary goal of this implementation is 
to enable testing of state machines and code together.

## Note

This part of code can be used as local execution as well as validation code before deploying to AWS Step Functions.

### Continuing Development

Step at the moment is still very beta, and its API will likely change more before it stabilizes. If you have ideas 
for improvements please reach out.

Some of the TODOs left for the library are:

1. Support for Parallel States
1. Better Validations e.g. making sure all states are reachable and executable
1. Client side visualization of state machine and execution using GraphViz

Unless we need more changes, we can continuously point to coinbase's repo.