CREATE DATABASE IF NOT EXISTS keycloak;
USE keycloak;
CREATE USER 'keycloak'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON keycloak.* TO 'keycloak'@'%' IDENTIFIED BY 'password';
FLUSH PRIVILEGES;