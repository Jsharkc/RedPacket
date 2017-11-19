# RedPacket API

Host : 

Port : 

##User Authentication 

1. 获取 Authentication token (假设已存在用户列表，模拟用户已登录状态，做测试用)

   ```
   URL: /gene/token
   Method: POST
   Input: 
   {
   	uid: 752413
   }
   Output:
   {
    	token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
    		eyJleHAiOjE1MTEwNzI1NzAsInVpZCI6MTIzfQ.
    		NcxJrRYtYwlzYkQra-KzjLN-ncJprOU1f6-LUQ_vDYA"
   }
   ```

## User Method

1. 用户查询余额

   ```
   URL: /user/balance
   Method: GET
   Output:
   {
    	balance: 220			// 用户钱包剩余金额 单位(分)
   }
   ```


## RedPacket Methods

所有操作在请求头加上 Authentication: Bearer ${token} (注：Bearer+空格+获取的 token)

1. 发口令红包 

   （注：限制只能发最大200元的红包，参数可在配置文件里调整，过24小时后剩余金额退到银行卡）

   ```
   URL: /redpack/send
   Method: POST
   Input: 
   {
   	total: 7524,		// 红包总金额，单位(分) (注：因为红包资金来源于银行卡，假设他资金充足)
   	num: 10				// 红包个数
   	blessing: "祝你生日快乐！"  // 祝福语
   }
   Output:
   {
    	pwd: "3dF731Ge"		// 抢红包的口令
   }
   ```

2. 抢红包（抢）

   ```
   URL: /redpack/grab
   Method: POST
   Input: 
   {
   	rpid: 2446,				// 红包ID
   	pwd: "3dF731Ge"			// 红包口令
   }
   Output:
   {
    	money: 220				// 抢到的金额
   }
   ```

3. 用户查看自己抢到的红包列表

   ```
   URL: /user/grab/list
   Method: GET
   Output:
   [
     {
       userid: 752413,				// 发红包人的 ID, 可用于获取昵称之类的信息
       blessing: "祝你生日快乐！",   // 祝福语	
       money: 220				    // 抢到了金额
     }, {
       userid: 234689,				   // 发红包人的 ID, 可用于获取昵称之类的信息
       blessing: "红红火火，恍恍惚惚",   // 祝福语	
       money: 720				       // 抢到了金额
     }
   ]
   ```

   ​


​