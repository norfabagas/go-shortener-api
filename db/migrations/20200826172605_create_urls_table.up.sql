CREATE TABLE IF NOT EXISTS urls(
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  url VARCHAR(255),
  code VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id)
);