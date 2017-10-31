# 运行结果 #

register表示注册用户：

![用户名为aaa的注册](http://img.blog.csdn.net/20171030235203152?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


此时打开userInfo文档可以看到在原有记录的末尾新增一条记录：

![新增记录](http://img.blog.csdn.net/20171031000514741?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


且在用户名存在的情况下报错拒绝请求：

![用户名已注册](http://img.blog.csdn.net/20171031000651307?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

其中第一行为log日志。


query的设计较为简单，将文件内容读入一个user数组中，返回该结果，在进行输出即可，实验结果如下：

![query结果](http://img.blog.csdn.net/20171031001122742?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


# go test 测试结果

## register 用户注册
## qu 用户查询
register -u 用户名 -p 密码 -e 电子邮件 -t 电话号码

由于register的设计过程中，没有用到复杂的算法，对create和myread部分做了test，其中create设计思路是检查append userList之后与原输入值是否一致，myread部分是直接判断是否正确读取。结果如下：
![这里写图片描述](http://img.blog.csdn.net/20171031001824124?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


## login 用户登录
## logout 用户登出


login -u 用户名 -p 密码 
logout


 ![1][1]
 
 

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




## del 删除用户 ##

 1. TestReadCur_succs. 用户处于已登录状态, 成功清空 curUser.txt 里的内容. Test通过
 2. TestSeekUsr. 查找 userInfo.json 中用户信息并成功删除. Test通过
 3. TestReadCur_fail. 用户处于未登录状态, 此时删除用户信息失败, 并提示要先登录. Test通过

![这里写图片描述](http://img.blog.csdn.net/20171031213628394?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)