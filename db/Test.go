package db

//CREATE TABLE `tests` (
//  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
//  `content` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
//  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0 正常，-1 已删除，1 已置顶',
//  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//  PRIMARY KEY (`id`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='测试表';

type Test struct {
	Id      int    `gorm:"column:id"`
	Content string `gorm:"column:content"`
	Status  *int   `gorm:"column:status"`
	//CreatedAt *time.Time `gorm:"column:created_at"`
	//UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Test) TableName() string {
	return "tests"
}

var (
	TestStatusDelete  = -1
	TestStatusDefault = 0
	TestStatusTop     = 1
)
