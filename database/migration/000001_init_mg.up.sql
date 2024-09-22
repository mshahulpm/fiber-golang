CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; 

create table users (
 user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 email varchar(255) not null unique,
 password varchar(255) not null ,
 created_at timestamp DEFAULT CURRENT_TIMESTAMP, 
 updated_at timestamp DEFAULT CURRENT_TIMESTAMP  
 );

 CREATE TABLE todo (
 todo_id serial PRIMARY KEY,
 todo text not null,
 user_id uuid not null,
 created_at timestamp DEFAULT CURRENT_TIMESTAMP,
 updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
 CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id) on DELETE CASCADE
 ) ;

 -- Create the trigger for the users table
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Create the trigger for the todo table
CREATE TRIGGER update_todo_updated_at
BEFORE UPDATE ON todo
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();