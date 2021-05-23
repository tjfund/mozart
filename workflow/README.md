# Workflow

Workflow takes DSL in common format such as: ASL or Serverless. Translate them on the fly to the underneath 
workflow engine implementation and upload them to the corresponding workflow engine.

It can also work in bypass mode to directly upload DSL to the native workflow engine. For example, ASL can be 
deployed into step functions as is. Netflix Conductor DSL can be deployed in to Netflix Conductor as is.

## Note

* Currently, only ASL is supported.
* ASL can be deployed to local Step Machine for testing as well.
