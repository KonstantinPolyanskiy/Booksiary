CREATE TABLE Technical_data (
    Id SERIAL NOT NULL,
    Created_at timestamp,
    Updated_at timestamp,
    PRIMARY KEY (Id)
);
CREATE TABLE Person (
    UUID uuid NOT NULL,
    Technical_data_id int,
    Name varchar,
    Surname varchar,
    Email varchar,
    PRIMARY KEY (UUID)
);
ALTER TABLE Person ADD CONSTRAINT person_fk0 FOREIGN KEY (Technical_data_id) REFERENCES Technical_data(Id);


