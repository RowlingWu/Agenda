# Login and Logout


##**AGENDA命令**
###login
用户登录login -u 用户名 -p 密码
-u = 用户名
-p = 密码
###logout
用户登出 logout

***
## **GOTEST**
###测试结果
 ![1][1]
 
 
###测试案例
- 登录
  >第一个测试所输入的用户名或密码错误，登录返回假，则符合预期结果
   第二个测试所输入的用户名与密码正确，登录返回真，则符合预期结果
   第三个测试因为已经登录，登录返回假，则符合预期结果
  ***
 >![2][2]

***
- 登出
  >第一个测试因为已经登录，登录返回真，则符合预期结果
   第二个测试因为没有登录，登录返回假，则符合预期结果
   ***
  >![3][3]




  [1]: https://imgsa.baidu.com/forum/w%3D580/sign=f1e8776eaa8b87d65042ab1737092860/3dc9a0efce1b9d16bf0e5f3ef8deb48f8d5464b7.jpg
  
  [2]: https://imgsa.baidu.com/forum/w%3D580/sign=c4f551e48cd6277fe912323018391f63/ece0c91b9d16fdfaeca503b1bf8f8c5495ee7bb7.jpg
  
  [3]: https://imgsa.baidu.com/forum/w%3D580/sign=e57f77b92c7f9e2f70351d002f31e962/85149a16fdfaaf51b72a44e0875494eef11f7ab7.jpg