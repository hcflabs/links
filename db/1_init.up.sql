CREATE TABLE links (
    -- user_id serial PRIMARY KEY,
    short_url VARCHAR (250) UNIQUE NOT NULL,
    target_url VARCHAR (2048) NOT NULL,
    permanent BOOLEAN NOT NULL,
    protected BOOLEAN NOT NULL,
    owned_by VARCHAR (255) NOT NULL,
    created TIMESTAMP NOT NULL default current_timestamp,
    modified TIMESTAMP NOT NULL default current_timestamp,
    description VARCHAR (500),
);
-- Create Indexes
CREATE UNIQUE INDEX short_url_index ON links(short_url);
-- Update Functions 
CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$ BEGIN NEW.modified = now();
RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TRIGGER update_links_modtime BEFORE
UPDATE ON links FOR EACH ROW EXECUTE PROCEDURE update_modified_column();