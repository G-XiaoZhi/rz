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
ApiVersion: judge/v0.1
Kind: ShareJudge
metedata:
  name: buy-avgs-ma5_ma250
  labels:
    action: buy
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  and_stgs:   # 这里定义list，不同策略是and操作
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
  or_stgs:   # 这里定义list，不同策略是or操作
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
  filter_stgs:
    filter_type: 1         # 1=纯数值 2=基本指标 3=指标微调
    filter_value: ["pe_ttm"]  # 数组形式，基于type灵活计算
    
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
  strategies:
    - stg_name: xxx
      stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
    - stg_name: yyy
      stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
    - stg_name: zzz
      stg_op: gt
      left_type: 1         # 1=纯数值 2=基本指标 3=指标微调 4=形态指标 ，支持后续扩展
      left_value: ["close", "*", 0.95]  # 数组形式，基于type灵活计算，left和right的type一致也可以不一致
      right_type: 1
      right_value: 1
  judgePolicy:
    - policyName: firstJudge
      operattion: And
      stgs: [xxx,yyy]
      policy: []
      finishPolicy: IfTrue | IfFalse | NotSure | End  # 依据这个策略的结果判断是否返回，也就是后续策略是否有继续执行的必要
    - policyName: secondeJudge
      operattion: Or
      stgs: [zzz]
      policy: [firstJudge]
      finishPolicy: IfTrue | IfFalse | NotSure | End  # 依据这个策略的结果判断是否返回，也就是后续策略是否有继续执行的必要

  # 这里的策略是，先计算firstPolicy里的stg [xxx & yyy]，这里可以设置finishPolicy=IfTrue
  # 然后计算secondPolicy里的stg和policy，zzz or firstPolicy
  policyContainer: [firstPolicy, secondPolicy]

  # examples:
  # finish表示判断结束，逻辑组合减少中间不必要的计算，从而支持各种复杂的组合计算
  # stgs: [xx, yy, zz, ww, tt]
  # 策略组合 xx & yy & zz & ww & tt， 一个judge，设置stgs 和 operation=and，finish=end
  # 策略组合 xx | yy | zz | ww | tt，一个judge，设置stgs 和 operation=or，finish=end
  # 策略组合 （xx | yy | zz） & （ ww | tt），两个judge，first: stgs=xx,yy,zz,op=or,finish=iffalse; second:stgs=ww,tt,op=or,finish=End;（这里其实不需要再做 first & second，因first为false直接返回了，直接返回second的结果即可）
  # 策略组合 （xx & yy ） | （ zz & ww & tt)，两个judge，first: stgs=xx,yy,op=and,finish=iftrue; second:stgs=zz,ww,tt,op=and,finish=End（这里其实不需要再做 first|second，因first为true直接返回了，直接返回second的结果即可）
  # 策略组合 xx &((yy & zz) | (ww & tt))，四个judge，first: stgs=xx,op=and,finish=iffalse; second:stgs=[yy,zz],op=and,finish=NotSure,third:stgs=[ww,tt],op=and,finish=NotSure;fourth:policy=[second,third],op=or,finish=End（这里其实不需要再做 xx & fourth，因为xx肯定为true，然后直接返回fourth的结果就是最终的结果）
```


```yaml
ApiVersion: judge/v1
Kind: IndexJudge   # 判断大盘状态
metedata:
  name: sh_index_state
  labels:
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  and_stgs:   # 这里定义list，不同策略是and操作
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
  or_stgs:   # 这里定义list，不同策略是or操作
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
ApiVersion: trade/v1
Kind: ShareTrade   # 交易系统的参数定义，买入和卖出
metedata:
  name: sale_shares
  labels:
    policy: stg_20210810
spec:
  inputShareListRef:  input_share_name    # input share yaml中metedata定义的name
  stgs:   # 这里定义list，不同策略是or操作
    buy_price: close / pre_close   # 支持自定义，程序中通过case选择使用的买入价格
    sale_price: close / pre_close   # 支持自定义，程序中通过case选择使用的卖出价格
    buy_num_stg: most / one    # 股票买入数量策略，一次买入最大 / 一次最多一个
```


```yaml
ApiVersion: stgConfig/v1
Kind: StgConfig   # 交易策略的初始化设置，支持反复
metedata:
  name: stg-config
  labels:
    policy: stg_20210810
spec:
  init_money: 100000
  start_date: 20210101
  end_date: 20210810
    
```


