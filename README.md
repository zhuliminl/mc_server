# mc_server


service repository pattern 笔记
1. repository 对外不应该包含 db 的语义, just as memory to use
   设计思路: 设计 API 接口 => 定义 dto => 设计 service => 开始写，最后写 controller

