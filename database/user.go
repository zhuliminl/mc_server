package database

const createUserTable = `
  CREATE TABLE IF NOT EXISTS user (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(32) COMMENT '随机默认用户名 | 用户自己设定名',
    wechatNickname VARCHAR(32) COMMENT '微信昵称',
    wechatNumber VARCHAR(20) COMMENT '微信号',
  );
`
