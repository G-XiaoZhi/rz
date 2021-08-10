## how to choose

- 通过yaml定义描述
- 选择的对象可以是股票、大盘等
- 基于条件计算，只返回符合条件的结果
- 结果对象的使用是另外的逻辑处理


```yaml
ApiVersion: inputShare/v1
Kind: InputShare
metedata:
  name: choose_zb_share
spec:
  origin_type: market / judgeShareRef     # 股票数据来源：市场分类，股票判断的返回
  value: zb / kcb / cyb / judge_name   # 主板 / 科创板 / 创业板 / judge yaml中metedata定义的name
  
```



```yaml
ApiVersion: judge/v1
Kind: ShareAndJudge
metedata:
  name: buy-avgs-ma5_ma250
  labels:
    action: buy
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  stgs：
    - stg_op: gt
      left_type: 1         # 1=指标+微调 2=形态指标
      left_value: ["close", "*", 0.95]  # 数组支持
      right_type: 1
      right_value: 1
  


```

```yaml
ApiVersion: judge/v1
Kind: SHIndexJudge

```
