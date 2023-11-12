CREATE TABLE Technical_data (
    Id SERIAL NOT NULL,
    Created_at timestamp,
    Update_at timestamp,
    PRIMARY KEY (Id)
);
CREATE TABLE Account (
    UUID uuid NOT NULL,
    Technical_data_id int,
    Login varchar,
    PasswordHash varchar,
    Role int default 1,
    PRIMARY KEY (UUID)
);
ALTER TABLE Account ADD CONSTRAINT account_fk0 FOREIGN KEY (Technical_data_id) REFERENCES technical_data(id);