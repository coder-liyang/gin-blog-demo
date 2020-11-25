# gin-blog-demo
https://www.bookstack.cn/read/gin-EDDYCJY-blog/golang-gin-2018-02-16-Gin%E5%AE%9E%E8%B7%B5-%E8%BF%9E%E8%BD%BD%E4%B8%80-Golang%E4%BB%8B%E7%BB%8D%E4%B8%8E%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85.md

标签表:
```sql
    CREATE TABLE `blog_tag` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(100) DEFAULT '' COMMENT '标签名称',
      `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
      `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
      `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
      `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
      `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```
文章表:
```sql
    CREATE TABLE `blog_article` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
      `title` varchar(100) DEFAULT '' COMMENT '文章标题',
      `desc` varchar(255) DEFAULT '' COMMENT '简述',
      `content` text,
      `created_on` int(11) DEFAULT NULL,
      `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
      `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
      `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
      `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
      `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
```
认证表:
```sql
    CREATE TABLE `blog_auth` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `username` varchar(50) DEFAULT '' COMMENT '账号',
      `password` varchar(50) DEFAULT '' COMMENT '密码',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
初始化数据:
```sql
    INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```