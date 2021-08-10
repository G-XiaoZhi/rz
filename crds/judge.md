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
  policy: stg_20210810   # 一个策略会对应多个yaml，然后给yaml打policy标签，用于识别这些yaml组成一个完整的策略
spec:
  origin_type: market / judgeShareRef     # 股票数据来源：市场分类，股票判断的返回
  value: zb / kcb / cyb / judge_name   # 主板 / 科创板 / 创业板 / judge yaml中metedata定义的name
  
```



```yaml
ApiVersion: judge/v1
Kind: ShareJudge
metedata:
  name: buy-avgs-ma5_ma250
  labels:
    action: buy
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  stgs_logical_op: and / or
  stgs：   # 这里定义list，不同策略是and / or操作
    - stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
    - stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
```


```yaml
ApiVersion: judge/v1
Kind: SHIndexJudge   # 判断大盘状态
metedata:
  name: sh_index_state
  labels:
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  stgs_logical_op: and / or
  stgs：   # 这里定义list，不同策略是or操作
    - stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
```


```yaml
ApiVersion: trade/v1
Kind: ShareTrade   # 交易系统的参数定义，买入和卖出
metedata:
  name: sale_shares
  labels:
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  stgs：   # 这里定义list，不同策略是or操作
    hold_share_num：3
    
```
