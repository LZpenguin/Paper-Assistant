# Paper-Assistant

文献推送助手 项目组 会议记录

# 第一次组会 2022/9/20



## 会议目标

- 确保每一位同学都能理解产品逻辑（多问）
- 确保每一位同学都有明确的任务分工（精确可量化）



### 1.讨论原型和功能逻辑

**原型的变动以及原因：**

1.发现页优先级降低，原有底部导航入口改为收藏

2.去除首页的分享功能，改为更多

3.将订阅入口从我的页面调整到首页侧边栏

4.期刊新增影响因子和期刊号、头像等信息要素

5.新增订阅侧边栏页面

6.新增我的收藏页面

7.新增历史分享页面



**需要讨论的问题：**

1.收藏的UI设计比较（原型内任何UI设计都可讨论）

2.两个版本收藏页面的逻辑比较和选择

3.如何使用户能够尽可能低成本的下载文献（现有方案：复制DOI and 发送下载链接到邮件 and 提供全文链接）

4.订阅侧边栏的删除方式？研究专题的实现可能性？

5.绑定邮箱的技术实现可能性？

6.收藏页面如何新增文件夹？

7.首页的更多功能需要加什么？

8.要不要显示影响因子？期刊位置放哪里？经济前%？

影响因子放在文献页

9.关键词匹配的逻辑关系？

10.推荐系统的逻辑

推荐些别的期刊？

看过了的文献要不要再推？

11.订阅界面入口改回我的页面

12.取消批量下载，下载入口改为复制链接形式，只存在于文献详情页中，并且通过电脑小程序同步打开方式，复制链接或DOI



**首页需要改动的点**

去掉影响因子

左上角去掉订阅入口

更多里面只保留不喜欢这篇文献



**订阅页需要改动的点**

去掉显式icon，改为长按编辑订阅刊物

去掉“我的研究专题”订阅块

MVP只做期刊



**文献详情页需要改动的点**

期刊要显示关注按钮，类似下图知乎用户，显示期刊期号和影响因子

![img](https://bingyan.feishu.cn/space/api/box/stream/download/asynccode/?code=NjM0YmU1NGI3OGJhMDI3Y2RjOGRiYzI4MjBkZWM5YzZfaGoxZFRoZlN1aVl6NEEyVG51Q0FQQ2I5M3ZmZzFjTDVfVG9rZW46Ym94Y242cm42b2N1UkZERGd4MTIxOW9wZGtnXzE2NjQyODAxNDE6MTY2NDI4Mzc0MV9WNA)

关键词 只显示一行，多的折叠，可以设置展开按钮

新增下载按钮

只给原文链接，不提供下载链接



**收藏页面需要改动的点**

- 长按文献则进行操作，那么显示分享、删除、移动
- 长按文件夹对文件夹进行删除、重命名操作
- 取消批量下载
- 新增文件夹和文件夹排列逻辑如下图

![img](https://bingyan.feishu.cn/space/api/box/stream/download/asynccode/?code=MDk2NWU1MDIwODk0YzRhMjNhNDRiZTZjOGQ4MjA3MzRfWmlpcmFjV1hDeWkyNzh5OXNrQlVyZ2hNdEpjUUR0NXJfVG9rZW46Ym94Y25jWjN0c1FvSEZXaUJQUzZzc3RobmRnXzE2NjQyODAxNDE6MTY2NDI4Mzc0MV9WNA)



我的页进行的改动

删除绑定邮箱

订阅入口移到我的页

联系我们改为关于我们