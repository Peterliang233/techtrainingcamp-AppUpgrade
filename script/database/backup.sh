db_user="root"
db_password="mysqlpassword"
db_name="app"
backup_dir="./"
time="$(date +"%Y_%m_%d_%H_%M_%S")"
echo "mysqldump -u $db_user -p$db_password $db_name > $backup_dir/$db_name"_"$time.sql"
mysqldump -u $db_user -p$db_password $db_name > $backup_dir/$db_name"_"$time.sql
find $backup_dir -name $db_name"*.sql" -type f -mmin +1 -exec rm -rf {} \; > /dev/null 2>&1