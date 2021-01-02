use gin_gorm_logrus_basic_db;

-- Read user
create user if not exists 'user_r'@'mysql' identified by 'local_read_pwd';
grant select on gin_gorm_logrus_basic_db.* to 'user_r'@'mysql';

-- Read/Write user
create user if not exists 'user_w'@'mysql' identified by 'local_write_pwd';
grant
    select, insert, update, delete, references,
    create, alter, drop, index, create view, show view, trigger,
    create routine, alter routine, execute, create temporary tables
    on *.* to 'user_w'@'mysql';

-- Management user
create user if not exists 'user_m'@'mysql' identified by 'local_manage_pwd';
grant
    create tablespace, create user, alter, drop, usage, create role, drop role,
    process, reload, shutdown, event, file, lock tables, super,
    replication client, replication slave, show databases
    on *.* to 'user_m'@'mysql';