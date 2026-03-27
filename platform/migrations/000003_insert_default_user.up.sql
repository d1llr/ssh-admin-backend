-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

-- pass: admin

INSERT INTO users (
  id,
  created_at,
  updated_at,
  email,
  password_hash,
  user_status,
  user_role
) VALUES (
           uuid_generate_v4(),
           NOW(),
           NOW(),
           'admin@gmail.com',
           '$2a$12$wJ3RuMrRqyupHiPhyDBcHeYsNmGtr5LWDVj9cb9p.OZaanZ/Y0.l6',
           1,
           'admin'
         );
