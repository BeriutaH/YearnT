package models

type CoreAccount struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Username   string `gorm:"type:varchar(50);not null;index:user_idx;comment:用户名" json:"username"`
	Password   string `gorm:"type:varchar(150);not null;comment:密码（加密存储）" json:"password"`
	Department string `gorm:"type:varchar(50);comment:部门名称" json:"department"`
	RealName   string `gorm:"type:varchar(50);comment:真实姓名" json:"real_name"`
	Email      string `gorm:"type:varchar(50);comment:邮箱地址" json:"email"`
	IsRecorder uint   `gorm:"type:tinyint(2) not null default 2;comment:是否记录员，1是，2否" json:"is_recorder"`
}

// CoreGlobalConfiguration 核心全局配置
type CoreGlobalConfiguration struct {
	ID            uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Authorization string `gorm:"type:varchar(50);not null;comment:授权信息" json:"authorization"`
	Ldap          DBJSON `gorm:"type:json;comment:LDAP配置" json:"ldap"`
	Message       DBJSON `gorm:"type:json;comment:消息配置" json:"message"`
	Other         DBJSON `gorm:"type:json;comment:其他配置信息" json:"other"`
	Stmt          uint   `gorm:"type:tinyint(2);not null;default:0;comment:状态码（状态标记）" json:"stmt"`
	AuditRole     DBJSON `gorm:"type:json;comment:审核角色配置" json:"audit_role"`
	Board         string `gorm:"type:longtext;comment:公告板内容" json:"board"`
	AI            DBJSON `gorm:"type:json;comment:AI相关配置" json:"ai"`
}

// CoreSqlRecord 核心sql记录
type CoreSqlRecord struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	WorkId    string `gorm:"type:varchar(50);not null;index:workId_idx;comment:工作ID" json:"work_id"`
	SQL       string `gorm:"type:longtext;not null;comment:执行的SQL语句" json:"sql"`
	State     string `gorm:"type:varchar(50);not null;comment:SQL执行状态" json:"state"`
	AffectRow uint   `gorm:"type:int(50);not null;comment:影响的行数" json:"affect_row"`
	Time      string `gorm:"type:varchar(50);not null;comment:执行时间" json:"time"`
	Error     string `gorm:"type:longtext;comment:错误信息" json:"error"`
}

// CoreSqlOrder 核心sql命令
type CoreSqlOrder struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	WorkId      string `gorm:"type:varchar(50);not null;index:workId_idx;comment:工作ID" json:"work_id"`
	Username    string `gorm:"type:varchar(50);not null;index:query_idx;comment:用户名" json:"username"`
	Status      uint   `gorm:"type:tinyint(2);not null;comment:状态" json:"status"`
	Type        int    `gorm:"type:tinyint(2);not null;comment:类型，1-DML，0-DDL" json:"type"` // 1 dml  0 ddl
	Backup      uint   `gorm:"type:tinyint(2);not null;comment:是否备份" json:"backup"`
	IDC         string `gorm:"type:varchar(50);not null;comment:机房" json:"idc"`
	Source      string `gorm:"type:varchar(50);not null;comment:来源" json:"source"`
	SourceId    string `gorm:"type:varchar(200);not null;index:source_idx;comment:来源ID" json:"source_id"`
	DataBase    string `gorm:"type:varchar(50);not null;comment:数据库名称" json:"data_base"`
	Table       string `gorm:"type:varchar(50);not null;comment:表名" json:"table"`
	Date        string `gorm:"type:varchar(50);not null;comment:日期" json:"date"`
	SQL         string `gorm:"type:longtext;not null;comment:SQL语句" json:"sql"`
	Text        string `gorm:"type:longtext;not null;comment:文本描述" json:"text"`
	Assigned    string `gorm:"type:varchar(550);not null;comment:指派人" json:"assigned"`
	Delay       string `gorm:"type:varchar(50);not null;default:'none';comment:延迟信息" json:"delay"`
	RealName    string `gorm:"type:varchar(50);not null;comment:真实姓名" json:"real_name"`
	ExecuteTime string `gorm:"type:varchar(50);comment:执行时间" json:"execute_time"`
	CurrentStep int    `gorm:"type:int(50);not null;default:1;comment:当前步骤" json:"current_step"`
	Relevant    DBJSON `gorm:"type:json;comment:关联信息JSON" json:"relevant"`
	OSCInfo     string `gorm:"type:longtext;default ''" json:"osc_info"`
	File        string `gorm:"type:varchar(200);not null;default:'';comment:文件路径" json:"file"`
}

type CoreRollback struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	WorkId string `gorm:"type:varchar(50);not null;index:workId_idx;comment:工作ID" json:"work_id"`
	SQL    string `gorm:"type:longtext;not null;comment:回滚的SQL语句" json:"sql"`
}

type CoreDataSource struct {
	ID               uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	IDC              string `gorm:"type:varchar(50);not null;comment:机房标识" json:"idc"`
	Source           string `gorm:"type:varchar(50);not null;comment:数据源名称" json:"source"`
	IP               string `gorm:"type:varchar(200);not null;comment:IP地址" json:"ip"`
	Port             int    `gorm:"type:int(10);not null;comment:端口号" json:"port"`
	Username         string `gorm:"type:varchar(50);not null;comment:用户名" json:"username"`
	Password         string `gorm:"type:varchar(150);not null;comment:密码" json:"password"`
	IsQuery          int    `gorm:"type:tinyint(2);not null;comment:读写权限，0写，1读，2读写" json:"is_query"`
	FlowID           int    `gorm:"type:int(100);not null;comment:流程ID" json:"flow_id"`
	SourceId         string `gorm:"type:varchar(200);not null;index:source_idx;comment:数据源唯一标识" json:"source_id"`
	ExcludeDbList    string `gorm:"type:varchar(200);not null;comment:排除的数据库列表" json:"exclude_db_list"`
	InsulateWordList string `gorm:"type:varchar(200);not null;comment:隔离词列表" json:"insulate_word_list"`
	Principal        string `gorm:"type:varchar(150);not null;comment:负责人" json:"principal"`
	CAFile           string `gorm:"type:longtext;default ''" json:"ca_file"`
	Cert             string `gorm:"type:longtext;default '';comment:客户端证书内容" json:"cert"`
	KeyFile          string `gorm:"type:longtext;comment:客户端密钥内容" json:"key_file"`
	DBType           int    `gorm:"type:int(5);not null;default:0;comment:数据库类型，0 MySQL，1 PostgreSQL" json:"db_type"`
	RuleId           int    `gorm:"type:int(100);not null;default:0;comment:规则ID" json:"rule_id"`
}

type CoreGrained struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Username string `gorm:"type:varchar(50);not null;index:user_idx;comment:用户名" json:"username"`
	Group    DBJSON `gorm:"type:json;comment:所属分组（JSON数组）" json:"group"`
}

type CoreRoleGroup struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Name        string `gorm:"type:varchar(50);not null;comment:角色名称" json:"name"`
	Permissions DBJSON `gorm:"type:json;comment:权限列表（JSON数组）" json:"permissions"`
	GroupId     string `gorm:"type:varchar(200);not null;index:group_idx;comment:分组唯一标识" json:"group_id"`
}

type CoreQueryOrder struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	WorkId       string `gorm:"type:varchar(50);not null;index:workId_idx;comment:工单ID" json:"work_id"`
	Username     string `gorm:"type:varchar(50);not null;comment:提交人用户名" json:"username"`
	Date         string `gorm:"type:varchar(50);not null;comment:提交日期" json:"date"`
	ApprovalTime string `gorm:"type:varchar(50);not null;comment:审批时间" json:"approval_time"`
	Text         string `gorm:"type:longtext;not null;comment:SQL内容或说明文本" json:"text"`
	Assigned     string `gorm:"type:varchar(50);not null;comment:分配执行人用户名" json:"assigned"`
	RealName     string `gorm:"type:varchar(50);not null;comment:提交人真实姓名" json:"real_name"`
	Export       uint   `gorm:"type:tinyint(2);not null;comment:是否导出，0否 1是" json:"export"`
	SourceId     string `gorm:"type:varchar(200);not null;index:source_idx;comment:数据源唯一标识" json:"source_id"`
	Status       int    `gorm:"type:tinyint(2);not null;index:status_idx;comment:工单状态" json:"status"`
}

// CoreQueryRecord 记录 SQL 查询执行的相关信息
type CoreQueryRecord struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	WorkId string `gorm:"type:varchar(50);not null;index:workId_idx;comment:工单ID" json:"work_id"`
	SQL    string `gorm:"type:longtext;not null;comment:执行的SQL语句" json:"sql"`
	ExTime int    `gorm:"type:int(10);not null;comment:SQL执行耗时（毫秒）" json:"ex_time"`
	Time   string `gorm:"type:varchar(50);not null;comment:执行时间" json:"time"`
	Source string `gorm:"type:varchar(50);not null;comment:数据源名称" json:"source"`
	Schema string `gorm:"type:varchar(50);not null;comment:数据库名称（schema）" json:"schema"`
}

// CoreAutoTask 自动化任务记录结构体
type CoreAutoTask struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	Name      string `gorm:"type:varchar(50);not null;comment:'任务名称'" json:"name"`
	Source    string `gorm:"type:varchar(50);not null;comment:'数据源标识'" json:"source"`
	SourceId  string `gorm:"type:varchar(200);not null;index:source_idx;comment:'数据源唯一 ID'"  json:"source_id"`
	DataBase  string `gorm:"type:varchar(50);not null;comment:'数据库名称'" json:"data_base"`
	Table     string `gorm:"type:varchar(50);not null;comment:'表名称'" json:"table"`
	Tp        int    `gorm:"type:tinyint(2);not null;comment:'操作类型：0 插入，1 更新，2 删除'" json:"tp"`
	Affectrow uint   `gorm:"type:int(50);not null default 0;comment:'影响行数'" json:"affect_rows"`
	Status    int    `gorm:"type:tinyint(2);not null default 0;comment:'任务状态：0 关闭，1 启用'" json:"status"`
	TaskId    string `gorm:"type:varchar(200);not null;index:task_idx;comment:'任务唯一标识'"  json:"task_id"`
	IDC       string `gorm:"type:varchar(50);not null;comment:'机房标识'" json:"idc"`
}

// CoreWorkflowTpl 工作流模板结构体
type CoreWorkflowTpl struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	Source string `gorm:"type:varchar(50);not null;index:source_idx;comment:'数据源标识'" json:"source"`
	Steps  DBJSON `gorm:"type:json;comment:'流程步骤（DBJSON 格式）'" json:"steps"`
}

// CoreWorkflowDetail 工作流执行记录
type CoreWorkflowDetail struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	WorkId   string `gorm:"type:varchar(50);not null;index:workId_idx;comment:'工单 ID'" json:"work_id"`
	Username string `gorm:"type:varchar(50);not null;index:query_idx;comment:'操作用户'" json:"username"`
	Time     string `gorm:"type:varchar(50);not null;comment:'操作时间'" json:"time"`
	Action   string `gorm:"type:varchar(550);not null;comment:'操作内容'" json:"action"`
}

type CoreOrderComment struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	WorkId   string `gorm:"type:varchar(50);not null;index:workId_idx;comment:'工单 ID'" json:"work_id"`
	Username string `gorm:"type:varchar(50);not null;comment:'评论用户'" json:"username"`
	Content  string `gorm:"type:longtext;comment:'评论内容'" json:"content"`
	Time     string `gorm:"type:varchar(50);not null;comment:'评论时间'" json:"time"`
}

type CoreRules struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	Desc      string `gorm:"type:longtext;not null;comment:'规则描述'" json:"desc"`
	AuditRole DBJSON `gorm:"type:json;comment:'审核角色配置（DBJSON）'" json:"audit_role"`
}

type CoreTotalTickets struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT;comment:'主键 ID'" json:"id"`
	Date       string `gorm:"type:varchar(50);not null;index:date_idx;comment:'统计日期'" json:"date"`
	TotalOrder int64  `gorm:"type:int(50);not null;comment:'总工单数量'" json:"total_order"`
	TotalQuery int64  `gorm:"type:int(50);not null;comment:'总查询数量'" json:"total_query"`
}
