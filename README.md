# 前提

比特币上交易花费需要手续费，而现在的手续费动辄数百元并不便宜，所以我们接下来的操作都在测网上进行。

## 效果

比特币浏览器里查询的交易

![image-20220318181529852](比特币刻字-简化版.assets/image-20220318181529852.png)

这里可以看到 `OP_RETURN`解码后的`你好,zhaojiuzhou`就是在比特链上刻字后的样子

# 步骤

## 生成自己的私钥和钱包地址

* 代码 [wallet_gen](https://github.com/coder-zhouge/btc_mark/blob/main/cmd/wallet_gen/main.go)
* 结果

![image-20220318160942066](比特币刻字-简化版.assets/image-20220318160942066.png)

> 私钥和钱包会保存在地址运行程序时候的工作目录下`testnet3_wallet.dat`下，如果需要重新生成，需要删除这个文件再运行程序

## 乞讨一些比特币 [testnet-faucet](https://testnet-faucet.mempool.co/)

![image-20220318161942504](比特币刻字-简化版.assets/image-20220318161942504.png)

![image-20220318162230397](比特币刻字-简化版.assets/image-20220318162230397.png)

> 用完后剩余的记得归还人家
>
> 如果这个水龙头失效了，那么再自己`google`一个吧

## 查看自己地址上的信息 

地址为: https://api.blockcypher.com/v1/btc/test3/addrs/mzD2vD8SBhiiJhD4FFZzC8U3jWLc65WxqH/full

> 读者要用自己的钱包地址替换串`mzD2vD8SBhiiJhD4FFZzC8U3jWLc65WxqH`

![image-20220318163255019](比特币刻字-简化版.assets/image-20220318163255019.png)

从`mark1`可以看到未确认的金额有`10W`聪，静等确认，我们就可以继续下面的操作了

## 生成交易`transaction`

* 代码 [mark](https://github.com/coder-zhouge/btc_mark/blob/main/cmd/mark/main.go)

* 替换代码中的下列值

  ![image-20220318174633053](比特币刻字-简化版.assets/image-20220318174633053.png)

  * `walletAddress`为你自己的钱包地址
  * `costTxHash`为`blockcypher`浏览器中信息那个图的`mark2`
  * `costTxOutputN`为`blockcypher`浏览器中信息那个图的第二个输入(index为1)
  * `costTxOutputNLockScript`为`blockcypher`浏览器中信息那个图的`mark3`
  * `priKeyWIF`为你自己钱包对应的私钥

* 运行程序，得到如下输出

  ![image-20220318175139445](比特币刻字-简化版.assets/image-20220318175139445.png)

  将`01000000010ca2`一直到`bc8100000000`这个十六进制串复制到地址`https://live.blockcypher.com/btc/pushtx/`的`Transaction Hex`输入框里，`Network`选择`Bitcoin Testnet`

  ![image-20220318175412720](比特币刻字-简化版.assets/image-20220318175412720.png)

  点击`Broadcast Transaction`按钮，一切正常后可以看到这个界面

  ![image-20220318174242518](比特币刻字-简化版.assets/image-20220318174242518.png)

## 查看战果

打开测链浏览器 `https://live.blockcypher.com/btc-testnet`, 输入你的地址进行查询

![image-20220318180957476](比特币刻字-简化版.assets/image-20220318180957476.png)

可以看到最新交易的txID

![image-20220318181106171](比特币刻字-简化版.assets/image-20220318181106171.png)

为`901a0c7ae485ef9d3c2e1e4e4ed6b14c15f3aa2eb73a6031c0471748ee1d9f52`

拼接成地址 `https://blockchair.com/zh/bitcoin/testnet/transaction/901a0c7ae485ef9d3c2e1e4e4ed6b14c15f3aa2eb73a6031c0471748ee1d9f52`

浏览器打开，就可以看到我们刻的字了

![image-20220318181323418](比特币刻字-简化版.assets/image-20220318181323418.png)

