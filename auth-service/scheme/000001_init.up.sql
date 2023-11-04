CREATE TABLE Credentials (
                             id SERIAL NOT NULL,
                             login varchar,
                             password_hash varchar,
                             created_at timestamp,
                             PRIMARY KEY (id)
);
CREATE TABLE user_auth_service (
                                   UUID uuid NOT NULL,
                                   credentials_id int,
                                   PRIMARY KEY (UUID)
);

ALTER TABLE user_auth_service ADD CONSTRAINT user_fk0 FOREIGN KEY (credentials_id) REFERENCES Credentials(id);

