因为weibo客户端的推送能力在手机上很差，会被杀，特别关心用户后无法即使收到动态推送，所以有了本软件。
本软件会每隔90秒(可设置)去获取weibo用户的开放数据，当有新动态时候将会通过gomail(SMTP)将动态内容，动态URL，动态图片发到到指定的邮箱。

时间仓促，代码质量较差，请包含。

完成的功能：
task        定时器获取信息
mysql       持久化
logging     日志记录
conf        runtime的conf
mail        SMTP发送邮件
