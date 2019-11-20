namespace go demo # 制定go的包是 nsgo
service myService {}

typedef i64 int64
# 单行注释
# 多行注释

# 常量
const i8 MATHPATH = 256
# 枚举
enum Operation {
    ADD=1,
    SUBSTRUCT=2,
    MULTIPLY=3,
    DIVIDE=4
}
# 结构体
struct Work {
    1:i32 num1=0,
    2:i32 num2,
    3:required Operation op,
    4:optional string comment
}
# 异常
exception InvalidOperation {
    1:i32 whatOp,
    2:string why
}
# 服务
service Calculator extends myService {
   void ping(),
   i32 add(1:i32 num1,2:i32 num2),
   i32 calculate(1:i32 logId,2:Work work) throws(1:InvalidOperation ouch),
   oneway void zip()
}