#SET GLOBAL validate_password.length = 6;
#SET GLOBAL validate_password.number_count = 0;
#SET GLOBAL validate_password.mixed_case_count = 0;
#SET GLOBAL validate_password.special_char_count = 0;
FLUSH PRIVILEGES;
CREATE USER 'soulogic'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT, INSERT, UPDATE, DELETE ON `soulogic`.* TO 'soulogic'@'localhost';
