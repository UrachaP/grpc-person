# grpc-person

### gRPC for test in my project

คำสั่งในการ generate จากไฟล์ .proto ให้เป็นภาษา go ประกอบด้วย 3 ส่วนคือ
<br>--go_out คือ output for the Go code generated
<br>--go-grpc_out คือ output for the Go code generating for gRPC
<br>และส่วนสุดท้ายคือไฟล์ proto ที่เราต้องการนำไป generate
```go
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ pkg/pb_gen/Person.proto
```