这是一个可查询进度的、可能会产生随机错误的DSL workflow。
产生错误时，可以通过retry task 进行重试。

进度根据

运行示例：
1) You need a Temporal service running. See README.md for more details.
2) Run
```
go run worker/main.go
```
to start worker for dsl workflow.
3) Run
```
go run starter/main.go -dslConfig=dsl/workflow1.yaml
```
to submit start request for workflow defined in `workflow1.yaml` file.
4) Run
```bash
go run query/main.go <workflow-id>
```
to query progress of the workflow. 

5) Run
```bash
go run retry/main.go <workflow-id>
```
to retry the failed workflow. 