# register及query相关


##使用方法
如readme.md中所示，register表示注册用户：
![用户名为aaa的注册](http://img.blog.csdn.net/20171030235203152?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
此时打开userInfo文档可以看到在原有记录的末尾新增一条记录：
![新增记录](http://img.blog.csdn.net/20171031000514741?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
且在用户名存在的情况下报错拒绝请求：
![用户名已注册](http://img.blog.csdn.net/20171031000651307?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
其中第一行为log日志。
query的设计较为简单，将文件内容读入一个user数组中，返回该结果，在进行输出即可，实验结果如下：
![query结果](http://img.blog.csdn.net/20171031001122742?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
##test文档的设计
由于register的设计过程中，没有用到复杂的算法，对create和myread部分做了test，其中create设计思路是检查append userList之后与原输入值是否一致，myread部分是直接判断是否正确读取。结果如下：
![这里写图片描述](http://img.blog.csdn.net/20171031001824124?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
